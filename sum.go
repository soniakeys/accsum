// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package accsum

import "math"

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
