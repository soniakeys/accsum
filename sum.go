// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package accsum

// Sum.go:  Algorithms not developed by S. Rump and colleagues but mentioned
// in one or more papers, implemented in the collection of Matlab reference
// code, or otherwise of interest.

import (
	"math"
	"sort"
)

// Sum returns a sum of the values in p.
//
// The algoithm is the simple sequential sum.
//
// It performs len(p) floating point additions.
func Sum(p []float64) (sum float64) {
	for _, x := range p {
		sum += x
	}
	return
}

// XSum returns a sum of the values in p.
//
// The algorithm is "XBLAS quadruple precision summantion."
func XSum(p []float64) float64 {
	// 4.6 of "Algorithms of "Accurate Sum and Dot Product",
	// http://www.ti3.tu-harburg.de/paper/rump/OgRuOi05.pdf
	var s, t float64
	for _, π := range p {
		t1, t2 := TwoSum(s, π)
		s, t = FastTwoSum(t1, t2+t)
	}
	return s + t
}

// Dot returns a dot product of the values in x and y.
//
// The algoithm is simple sequential sum of products.
//
// X and y must be of the same length, panic or nonsense results otherwise.
//
// Dot performs 2*len(x) floating point operations.
func Dot(x, y []float64) float64 {
	s := 0.
	for i, xi := range x {
		s += xi * y[i]
	}
	return s
}

// XDot returns a dot product of the values in x and y.
//
// The algorithm is "XBLAS quadruple precision dot product."
//
// X and y must be of the same length, panic or nonsense results otherwise.
func XDot(x, y []float64) float64 {
	var s, t float64
	for i, xi := range x {
		h, r := TwoProduct(xi, y[i])
		s1, s2 := TwoSum(s, h)
		t1, t2 := TwoSum(t, r)
		s2 += t
		t1, s2 = FastTwoSum(s1, s2)
		t2 += s2
		s, t = FastTwoSum(t1, t2)
	}
	return s
}

// KahanSum returns a sum of the values in p.
//
// The algoithm is Kahan (1965), often termed "compensated" summation.
//
// It performs 4 * len(p) floating point operations (additions or
// subtractions.)
func KahanSum(p []float64) float64 {
	var s, c float64
	for _, x := range p {
		y := x - c
		t := s + y
		c = t - s - y
		s = t
	}
	return s
}

// KahanB computes a sum of the values in p.
//
// The algorithm is Kahan-Babuška-Neumaier, sometimes termed a "balancing
// summation.)
//
// It performs 7 * len(p) + 1 floating point operations (addition, subtraction,
// Abs, and comparison.)
func KahanB(p []float64) float64 {
	s := p[0]
	c := 0.
	for _, x := range p[1:] {
		t := s + x
		if math.Abs(s) >= math.Abs(x) {
			c += s - t + x
		} else {
			c += x - t + s
		}
		s = t
	}
	return s + c
}

// PriestSum computes a sum of the values in p.
//
// Algorithm following Matlab code PriestSum.m by S.M. Rump.  This is Priest's
// "Doubly compensated summation" on p. 64 of the 1992 paper "On properties"
// of floating point arithmetics: Numerical stability."
//
// Time complexity is O(n log n) in len(p).
func PriestSum(p []float64) float64 {
	if len(p) == 0 {
		return 0.
	}
	sort.Sort(priest(p))
	s := p[0]
	c := 0.
	for _, π := range p[1:] {
		y, u := FastTwoSum(c, π)
		t, v := FastTwoSum(s, y)
		z := u + v
		s, c = FastTwoSum(t, z)
	}
	return s
}

// a type for sorting by decreasing magnitude
type priest []float64

func (p priest) Len() int           { return len(p) }
func (p priest) Less(i, j int) bool { return math.Abs(p[i]) > math.Abs(p[j]) }
func (p priest) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// CondSum computes the condition number of summation function f over slice s.
//
// CondSum is not destructive on s even if f is destructive on its argument.
func CondSum(f func([]float64) float64, s []float64) float64 {
	c := append([]float64{}, s...)
	absSum := math.Abs(f(c))
	for i, x := range s {
		c[i] = math.Abs(x)
	}
	return f(c) / absSum
}

// CondDot computes the condition number of dot product function f over slices
// x and y.
func CondDot(f func(x, y []float64) float64, x, y []float64) float64 {
	cx := append([]float64{}, x...)
	cy := append([]float64{}, y...)
	absDot := math.Abs(f(cx, cy))
	for i, xi := range x {
		cx[i] = math.Abs(xi)
		cy[i] = math.Abs(y[i])
	}
	return 2 * f(cx, cy) / absDot
}
