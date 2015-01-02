// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package accsum

import (
	"fmt"
	"math"
	"math/rand"
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

// Section:  "Algorithms of "Accurate Sum and Dot Product",
// http://www.ti3.tu-harburg.de/paper/rump/OgRuOi05.pdf
//
// FastTwoSum (1.1)
// TwoSum (3.1)
// split (3.2)
// TwoProduct (3.3)
// Sum2 (4.1, 4.4)
// vecSum (4.3)
// SumK (4.8)
// SumKVert (4.12)
// Dot2 (5.3)
// Dot2Err (5.8)
// DotK (5.10)
// GenDot (6.1)

// FastTwoSum computes an error-free sum of two float64s, with conditions on
// the relative magnitudes.
//
// Error-free means the result x is floating-point sum a+b, and y is the
// floating-point error such that x+y exactly equals a+b.
//
// Results are accurate when |b| <= |a|, but are also still accurate as
// long as no trailing nonzero bit of a is smaller than the least significant
// bit of b.
//
// Dekker algorithm, 3 floating point operations.
func FastTwoSum(a, b float64) (x, y float64) {
	x = a + b
	y = a - x + b
	return
}

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

var splitFactor = math.Ldexp(1, 27) + 1

// split splits a into x, y such that x + y = a and both x and y need at most
// 26 bits in the significand.
//
// Requires 4 floating-point operations (multiplication and subtraction.)
func split(a float64) (x, y float64) {
	c := splitFactor * a
	x = c - (c - a)
	y = a - x
	return
}

// TwoSum computes an error-free product of two float64s.
//
// Result x is a*b, y is the error such that x+y exactly equals a times b.
//
// 17 floating point operations (multiplication and subtraction.)
func TwoProduct(a, b float64) (x, y float64) {
	x = a * b
	a1, a2 := split(a)
	b1, b2 := split(b)
	y = a2*b2 - (x - a1*b1 - a2*b1 - a1*b2)
	return
}

// Sum2 returns a sum of values in p as if computed in twice the precision
// of a float64.
func Sum2(p []float64) float64 {
	if len(p) == 0 {
		return 0.
	}
	s := p[0]
	var e, y float64
	for _, x := range p[1:] {
		s, y = TwoSum(s, x)
		e += y
	}
	return s + e
}

func vecSum(p []float64) {
	if len(p) < 2 {
		return
	}
	s := p[0]
	for i, x := range p[1:] {
		s, p[i] = TwoSum(s, x)
	}
	p[len(p)-1] = s
}

// SumK returns a sum of values in p, as if computed in k-fold precision of
// a float64.
//
// SumK is destructive on values in p.
func SumK(p []float64, K int) float64 {
	for K--; K > 0; K-- {
		vecSum(p)
	}
	return Sum(p)
}

// SumK returns a sum of values in p, as if computed in k-fold precision of
// a float64.
//
// SumKVert computes the same result as SumK but leaves values in p unmodified.
func SumKVert(p []float64, K int) float64 {
	if len(p) < K {
		K = len(p)
	}
	q := make([]float64, K-1)
	for i, s := range p[:len(q)] {
		for k, qk := range q[:i] {
			q[k], s = TwoSum(qk, s)
		}
		q[i] = s
	}
	s := 0. // Unclear from the paper, but this seems right.
	for _, α := range p[len(q):] {
		for k, qk := range q {
			q[k], α = TwoSum(qk, α)
		}
		s += α
	}
	for j, α := range q[:K-2] {
		for k := j + 1; k < len(q); k++ {
			q[k], α = TwoSum(q[k], α)
		}
		s += α
	}
	return s + q[K-2]
}

// Dot2 returns a dot product of x and y as if computed in twice the precision
// of a float64.
func Dot2(x, y []float64) float64 {
	if len(x) == 0 {
		return 0
	}
	q := 0.
	p, s := TwoProduct(x[0], y[0])
	for i := 1; i < len(x); i++ {
		h, r := TwoProduct(x[i], y[i])
		p, q = TwoSum(p, h)
		s += q + r
	}
	return p + s
}

// Dot2 returs a dot product and an error bound.
//
// The result dot is the same 2-fold precision result returned by Dot2,
// the result eb is a rigorous error bound.
func Dot2Err(x, y []float64) (dot, eb float64) {
	p, s := TwoProduct(x[0], y[0])
	e := math.Abs(s)
	q := 0.
	for i := 1; i < len(x); i++ {
		h, r := TwoProduct(x[i], y[i])
		p, q = TwoSum(p, h)
		t := q + r
		s += t
		e += math.Abs(t)
	}
	dot = p + s
	n := float64(len(x))
	δ := n * eps / (1 - 2*n*eps)
	α := eps*math.Abs(dot) + (δ*e + 3*eta/eps)
	eb = α / (1 - 2*eps)
	return
}

// DotK returns a dot product of x and y as if computed in K times the
// precision of a float64.
func DotK(x, y []float64, K int) float64 {
	r := make([]float64, 2*len(x))
	var p, h float64
	p, r[0] = TwoProduct(x[0], y[0])
	for i := 1; i < len(x); i++ {
		h, r[i] = TwoProduct(x[i], y[i])
		p, r[len(x)+i-1] = TwoSum(p, h)
	}
	r[2*len(x)-1] = p
	return SumK(r, K-1)
}

// GenDot generates vectors x and y ill-conditioned for dot product.
//
// Argument n specifies length of result vectors x and y, argument c
// specifies the approximate condition number for a dot product of x and y.
//
// Result d is a computed dot product that is exact or nearly exact,
// result C is the computed number.
//
// GenDot uses the rand package default generator, use rand.Seed as needed
// before calling GenDot.
func GenDot(n int, c float64) (x, y []float64, d, C float64) {
	n2 := (n + 1) / 2
	x = make([]float64, n)
	y = make([]float64, n)

	b := math.Log2(c)
	b2 := b / 2
	e := make([]int, n2)
	last := len(e) - 1
	for i := 1; i < last; i++ {
		e[i] = int(rand.Float64()*b2 + .5)
	}
	e[0] = int(b2+.5) + 1
	e[last] = 0
	for i := 0; i < n2; i++ {
		x[i] = math.Ldexp(rand.Float64()*2-1, e[i])
		y[i] = math.Ldexp(rand.Float64()*2-1, e[i])
	}

	// DotExact.  Is this K reasonable for exact result?
	// fmt.Println("using K =", int(b/20), "for DotExact")
	dx := func(x, y []float64) float64 { return DotK(x, y, int(b/20)) }

	f := b2 / float64(n-1-n2)
	for i := n2; i < n; i++ {
		e2 := int(float64(n-1-i)*f + .5)
		x[i] = math.Ldexp(rand.Float64()*2-1, e2)
		y[i] = (math.Ldexp(rand.Float64()*2-1, e2) - dx(x, y)) / x[i]
	}

	for i := n - 1; i >= 1; i-- {
		j := rand.Intn(i + 1)
		x[i], x[j] = x[j], x[i]
		y[i], y[j] = y[j], y[i]
	}

	d = dx(x, y)
	C = CondDot(dx, x, y)
	return
}

// Section:  Algorithms of "Accurate Floating-Point Summation, Part I:
// Faithful Rounding", http://www.ti3.tu-harburg.de/paper/rump/RuOgOi07I.pdf
//
// extractScalar (3.2)
// extractSlice (3.4)
// transform (4.1, 4.4)
// AccSum (4.5)

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

// transform just as needed for AccSum, without bells and whistles.
func transform(p []float64) (τ1, τ2 float64) {
	return transform3(p, 0, _ΦSum)
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

// Section:  Algorithms of "Accurate Floating-Point Summation, Part II:
// Faithful Rounding", http://www.ti3.tu-harburg.de/paper/rump/RuOgOi07II.pdf
//
// transform3 (3.3)
// AccSign (4.1)
// transformK (6.2)

// suitable values for argument Φ in transform3
func _ΦSum(Ms float64) float64  { return u * Ms * Ms }
func _ΦSign(Ms float64) float64 { return u * Ms }

func transform3(p []float64, ρ float64, Φ func(Ms float64) float64) (τ1, τ2 float64) {
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
	ϕ := Ms * u // "factor to decrease σ"
	_Φ := Φ(Ms) // "stopping criterion"
	for t := ρ; ; {
		τ := extractSlice(σ, p)
		τ1 = t + τ
		if math.Abs(τ1) >= _Φ*σ || σ <= minPos {
			τ2 = t - τ1 + τ
			return
		}
		t = τ1
		if t == 0 {
			return transform3(p, 0, Φ)
		}
		σ *= ϕ
	}
}

// AccSignBit returns the sign bit of the sum of values in p, somewhat faster
// than an accurate sum can be computed.
func AccSignBit(p []float64) bool {
	τ1, _ := transform3(p, 0, _ΦSign)
	return math.Signbit(τ1)
}

func transformK(p []float64, ρ float64) (res, R float64) {
	// code similar to AccSum
	τ1, τ2 := transform3(p, ρ, _ΦSum)
	sum := 0.
	for _, pi := range p {
		sum += pi
	}
	res = sum + τ2 + τ1 // same as AccSum result
	R = τ2 - (res - τ1)
	return
}

func AccSumK(p []float64, K int) []float64 {
	res := make([]float64, K)
	r := 0.
	for k := range res {
		res[k], r = transformK(p, r)
		if res[k] <= minPos {
			break
		}
	}
	return res
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

// nextPowerTwo returns the smallest power of 2 not less than abs(p).
//
// Result is computed in 4 floating point operations.
func nextPowerTwo(p float64) float64 {
	q := invEps * p
	return math.Abs(q - p - q)
}
