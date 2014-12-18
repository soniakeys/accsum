// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package accsum

import "math"

// Defining constants for IEEE 754 binary64, the Go float64 type.
const (
	P    = 53   // float64 precision
	EMax = 1023 // float64 max exponent
)

var (
	// eps is the "relative rounding error unit," the distance from 1.0 to
	// the next smaller number.
	eps    = math.Ldexp(1, -P)
	invEps = math.Ldexp(1, P)

	// 1/2 error unit
	u = math.Ldexp(1, -P-1)

	// eta is the "underflow unit," the smallest positive subnormal number.
	eta = math.Ldexp(1, 2-EMax-P)

	// minPos is the smallest positive normalized number.
	minPos = math.Ldexp(1, 1-EMax)
)

// TwoSum computes an error-free sum of two float64s.
//
// Knuth algorithm, 6 floating point operations.
//
// Result x is a+b, y is the error such that x+y exactly equals a+b.
func TwoSum(a, b float64) (x, y float64) {
	x = a + b
	z := x - a
	y = a - (x - z) + (b - z)
	return
}

// FastTwoSum computes an error-free sum of two float64s, with conditions on
// the relative magnitudes.
//
// Dekker algorithm, 3 floating point operations.
//
// Results are accurate when |b| <= |a|, but are also still accurate as
// long as no trailing nonzero bit of a is smaller than the least significant
// bit of b.
//
// Result x is a+b, y is the error such that x+y exactly equals a+b.
func FastTwoSum(a, b float64) (x, y float64) {
	x = a + b
	y = a - x + b
	return
}

// nextPowerTwo returns the smallest power of 2 not less than abs(p).
//
// Result is computed in 4 floating point operations.
func nextPowerTwo(p float64) float64 {
	q := invEps * p
	return math.Abs(q - p - q)
}

// extractScalar splits p relative to σ, which must be an integral power of 2.
//
// Return value q is the high order part of p, return value pʹ is the remainder
// such that q+pʹ exactly equals p.
// 3 floating point operations.
func extractScalar(σ, p float64) (q, pʹ float64) {
	q = σ + p - σ
	pʹ = p - q
	return
}

// extractSlice splits elements of p relative to σ.
//
// As with extractScalar, σ must be an integral power of 2.  extractSlice
// calls extractScalar on each element of p.  It replaces each element of p
// with the high order part q, and sums all remainders pʹ to the return
// value τ.
//
// Return value τ plus the sum of the new elements of p will exactly equal the
// sum of the original elements of p.
//
// 4 * len(p) floating point operations.
func extractSlice(σ float64, p []float64) (τ float64) {
	var q float64
	for i, pi := range p {
		q, p[i] = extractScalar(σ, pi)
		τ += q
	}
	return
}

func transform(p []float64) (τ1, τ2 float64) {
	if len(p) == 0 {
		return
	}
	μ := math.Abs(p[0])
	for _, x := range p[1:] {
		if a := math.Abs(x); x > μ {
			μ = a
		}
	}
	if μ == 0 {
		return
	}
	Ms := nextPowerTwo(float64(len(p) + 2))
	σ := Ms * nextPowerTwo(μ) // "extraction unit"
	if math.IsInf(σ, 0) {
		return σ, σ
	}
	ϕ := Ms * u  // "factor to decrease σ"
	_Φ := Ms * ϕ // "stopping criterion"

	for t := 0.; ; {
		τ := extractSlice(σ, p)
		τ1 := t + τ
		if math.Abs(τ1) >= _Φ*σ || σ <= minPos {
			return FastTwoSum(t, τ)
		}
		t = τ1
		if t == 0 {
			return transform(p)
		}
		σ *= ϕ
	}
}

// AccSum returns an accurate sum of values in p.
func AccSum(p []float64) float64 {
	τ1, τ2 := transform(p)
	sum := 0.
	for _, pi := range p { // order not important
		sum += pi
	}
	return sum + τ2 + τ1 // order important
}
