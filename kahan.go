// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package accsum

import (
	"math"
)

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
