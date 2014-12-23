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

// Cond computes the condition number of summation function over slice s.
//
// Cond is not destructive on s even if f is destructive on its argument.
func Cond(f func([]float64) float64, s []float64) float64 {
	c := append([]float64{}, s...)
	absSum := math.Abs(f(c))
	for i, x := range s {
		c[i] = math.Abs(x)
	}
	return f(c) / absSum
}
