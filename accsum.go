// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package accsum

import (
	"fmt"
	"math"
)

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

	// PrecSum maximum length of argument p.
	nMax = 1<<26 - 2
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
//
// AccSum is destructive on p.
//
// Result is a faithful rounding of the sum of values in p.
func AccSum(p []float64) float64 {
	τ1, τ2 := transform(p)
	sum := 0.
	for _, pi := range p { // order not important
		sum += pi
	}
	return sum + τ2 + τ1 // order important
}

// Cond computes the condition number of the summation of p.
func Cond(p []float64) float64 {
	c := append([]float64{}, p...)
	absSum := math.Abs(AccSum(c))
	for i, x := range p {
		c[i] = math.Abs(x)
	}
	return AccSum(c) / absSum
}

// PrecSum returns an accurate sum of values in p.
//
// Result is a faithful rounding of the sum or else has relative error <=
// 2^(-53*k) * Cond(p).
func PrecSum(p []float64, K int) float64 {
	switch {
	case len(p) == 0:
		return 0.
	case len(p) > nMax:
		panic(fmt.Sprintf("len(p) = %d exceeds limit, %d", len(p), nMax))
	}
	μ := math.Abs(p[0])
	for _, x := range p[1:] {
		if a := math.Abs(x); x > μ {
			μ = a
		}
	}
	μ /= 1 - float64(len(p))*2*eps
	if μ == 0 {
		return 0.
	}
	σ0 := nextPowerTwo(μ)
	if math.IsInf(σ0, 0) {
		return σ0
	}
	Ms := nextPowerTwo(float64(len(p) + 2))
	M := math.Log2(Ms)
	ϕ := Ms * u
	// len(σ) is L in paper and reference code.  also, paper and reference code
	// seem to allocate and then compute an extra σ element that is never used.
	σ := make([]float64,
		int(math.Ceil((float64(K)*math.Log2(u)-2)/(math.Log2(u)+M))-1))
	for k := 0; ; {
		if σ0 <= minPos {
			σ = σ[:k]
			break
		}
		σ[k] = σ0
		k++
		if k == len(σ) {
			break
		}
		σ0 *= ϕ
	}
	if len(σ) == 0 {
		sum := 0.
		for _, x := range p {
			sum += x
		}
		return sum
	}
	var q, sum float64
	τ := make([]float64, len(σ))
	for _, π := range p {
		for k, σk := range σ {
			q, π = extractScalar(σk, π)
			τ[k] += q
		}
		sum += π
	}
	π := τ[0]
	e := 0.
	for _, τk := range τ[1:] {
		π, q = FastTwoSum(π, τk)
		e += q
	}
	return sum + e + π
}
