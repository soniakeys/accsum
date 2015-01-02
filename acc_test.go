// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package accsum_test

import (
	"fmt"
	//	"math"
	"math/rand"

	"github.com/soniakeys/accsum"
)

func ExampleFastTwoSum() {
	a, b := .2, .1
	x, y := accsum.FastTwoSum(a, b)
	fmt.Printf("a: % .20f\n", a)
	fmt.Printf("b: % .20f\n", b)
	fmt.Printf("x: % .20f\n", x)
	fmt.Printf("y: % .20f\n", y)
	// Output:
	// a:  0.20000000000000001110
	// b:  0.10000000000000000555
	// x:  0.30000000000000004441
	// y: -0.00000000000000002776
}

func ExampleTwoSum() {
	a, b := .1, .2
	x, y := accsum.TwoSum(a, b)
	fmt.Printf("a: % .20f\n", a)
	fmt.Printf("b: % .20f\n", b)
	fmt.Printf("x: % .20f\n", x)
	fmt.Printf("y: % .20f\n", y)
	// Output:
	// a:  0.10000000000000000555
	// b:  0.20000000000000001110
	// x:  0.30000000000000004441
	// y: -0.00000000000000002776
}

func ExampleTwoProduct() {
	a := 1e10 + 1
	b := 1e6 + 1
	fmt.Println(accsum.TwoProduct(a, b))
	// Output: 1.0000010001e+16 1
}

func ExampleSum2() {
	n := 54321
	p := make([]float64, n+1)
	for i := range p {
		p[i] = float64(i)
	}
	p[0] = 1e20
	fmt.Printf("Simple:   %.16e\n", accsum.Sum(p))
	fmt.Printf("Sum2:     %.16e\n", accsum.Sum2(p))
	fmt.Println("Triangle:            ", n*(n+1)/2)
	// Output:
	// Simple:   1.0000000000146203e+20
	// Sum2:     1.0000000000147541e+20
	// Triangle:             1475412681
}

func ExampleSumK() {
	n := 54321
	p := make([]float64, n+1)
	for i := range p {
		p[i] = float64(i)
	}
	p[0] = 1e20
	fmt.Printf("Simple:   %.16e\n", accsum.Sum(p))
	fmt.Printf("SumK:     %.16e\n", accsum.SumK(p, 2))
	fmt.Println("Triangle:            ", n*(n+1)/2)
	// Output:
	// Simple:   1.0000000000146203e+20
	// SumK:     1.0000000000147541e+20
	// Triangle:             1475412681
}

func ExampleSumKVert() {
	n := 54321
	p := make([]float64, n+1)
	for i := range p {
		p[i] = float64(i)
	}
	p[0] = 1e20
	fmt.Printf("Simple:   %.16e\n", accsum.Sum(p))
	fmt.Printf("SumKVert: %.16e\n", accsum.SumKVert(p, 2))
	fmt.Println("Triangle:            ", n*(n+1)/2)
	// Output:
	// Simple:   1.0000000000146203e+20
	// SumKVert: 1.0000000000147541e+20
	// Triangle:             1475412681
}

func ExampleDot2() {
	n := 4321
	x := make([]float64, n+1)
	for i := range x {
		x[i] = float64(i)
	}
	x[0] = 1e11
	fmt.Printf("Simple:   %.16e\n", accsum.Dot(x, x))
	fmt.Printf("Dot2:     %.16e\n", accsum.Dot2(x, x))
	fmt.Println("Square triangle:      ", n*(n+1)*(2*n+1)/6)
	// Output:
	// Simple:   1.0000000000026734e+22
	// Dot2:     1.0000000000026902e+22
	// Square triangle:       26901858961
}

func ExampleDot2Err() {
	n := 4321
	x := make([]float64, n+1)
	for i := range x {
		x[i] = float64(i)
	}
	x[0] = 1e11
	fmt.Printf("Simple:   %.16e\n", accsum.Dot(x, x))
	dot, err := accsum.Dot2Err(x, x)
	fmt.Printf("Dot2:     %.16e\n", dot)
	fmt.Printf("Square triangle:   %15d\n", n*(n+1)*(2*n+1)/6)
	fmt.Printf("Dot2Err:           %15.0f\n", err)
	// Output:
	// Simple:   1.0000000000026734e+22
	// Dot2:     1.0000000000026902e+22
	// Square triangle:       26901858961
	// Dot2Err:                   1110223
}

func ExampleDotK() {
	n := 4321
	x := make([]float64, n+1)
	for i := range x {
		x[i] = float64(i)
	}
	x[0] = 1e11
	fmt.Printf("Simple:    %.16e\n", accsum.Dot(x, x))
	fmt.Printf("DotK(K=3): %.16e\n", accsum.DotK(x, x, 2))
	fmt.Println("Square triangle:       ", n*(n+1)*(2*n+1)/6)
	// Output:
	// Simple:    1.0000000000026734e+22
	// DotK(K=3): 1.0000000000026902e+22
	// Square triangle:        26901858961
}

func ExampleGenDot() {
	rand.Seed(42)
	x, y, d, c := accsum.GenDot(6, 1e31)
	fmt.Println("+x          y")
	for i, xi := range x {
		fmt.Printf("%+.2e  %+.2e\n", xi, y[i])
	}
	fmt.Println()
	fmt.Printf("condition:  %.2e\n", c)
	fmt.Printf("dot exact: % .6f\n", d)
	fmt.Printf("Dot2(x,y): % .6f\n", accsum.Dot2(x, y))
	fmt.Printf("Dot(x,y):  % .2e\n", accsum.Dot(x, y))
	// Output:
	// +x          y
	// -3.91e+15  +9.38e+14
	// -5.20e+14  -7.04e+15
	// -5.64e-01  +5.61e+07
	// -3.05e+05  -4.78e+05
	// -2.34e-01  +6.26e-01
	// +1.96e+07  +4.90e+07
	//
	// condition:  5.30e+31
	// dot exact: -0.276634
	// Dot2(x,y): -0.250000
	// Dot(x,y):   3.99e+14
}

func ExampleAccSum() {
	n := 54321
	p := make([]float64, n+1)
	for i := range p {
		p[i] = float64(i)
	}
	p[0] = 1e20
	fmt.Printf("Simple:   %.16e\n", accsum.Sum(p))
	fmt.Printf("AccSum:   %.16e\n", accsum.AccSum(p))
	fmt.Println("Triangle:            ", n*(n+1)/2)
	// Output:
	// Simple:   1.0000000000146203e+20
	// AccSum:   1.0000000000147541e+20
	// Triangle:             1475412681
}

/*
func ExampleAccSumK() {
	n := 54321
	p := make([]float64, n+1)
	for i := range p {
		p[i] = float64(i)
	}
	p[0] = 1e20
	fmt.Printf("Simple:    %.16e\n", accsum.Sum(p))
	r := accsum.AccSumK(p, 2)
	fmt.Printf("AccSumK:   %.16e + %g\n", r[0], r[1])
	fmt.Println("Triangle:             ", n*(n+1)/2)
	// Output:
	// Simple:   1.0000000000146203e+20
	// AccSum:   1.0000000000147541e+20 + 2681
	// Triangle:             1475412681
}
*/

/* fails.  don't know why
func ExampleAccSignBit() {
	fmt.Println("0  negative:    ", math.Signbit(0))
	fmt.Println("-1 negative:    ", math.Signbit(-1))
	fmt.Println("AccSignBit 0:   ", accsum.AccSignBit([]float64{0}))
	fmt.Println("AccSignBit -1:  ", accsum.AccSignBit([]float64{-1}))
	p := []float64{1e20, -1, -1e20}
	fmt.Println("Signbit(Sum(p)):", math.Signbit(accsum.Sum(p)))
	p = []float64{1e20, -1, -1e20}
	fmt.Println("AccSignBit(p):  ", accsum.AccSignBit(p))
	fmt.Println("p after AccSignBit:", p)
	// Output:
	// 0  negative:     false
	// -1 negative:     true
	// AccSignBit 0:    false
	// AccSignBit -1:   true
	// Signbit(Sum(p)): false
	// AccSignBit(p):   true
}
*/

func ExamplePrecSum() {
	n := 54321
	p := make([]float64, n+1)
	for i := range p {
		p[i] = float64(i)
	}
	p[0] = 1e20
	s := 0.
	for _, x := range p {
		s += x
	}
	fmt.Printf("Simple:   %.16e\n", s)
	fmt.Printf("PrecSum:  %.16e\n", accsum.PrecSum(p, 2))
	fmt.Println("Triangle:            ", n*(n+1)/2)
	// Output:
	// Simple:   1.0000000000146203e+20
	// PrecSum:  1.0000000000147541e+20
	// Triangle:             1475412681
}
