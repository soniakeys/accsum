// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package accsum_test

import (
	"fmt"

	"github.com/soniakeys/accsum"
)

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

func ExampleAccSum() {
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
	fmt.Printf("AccSum:   %.16e\n", accsum.AccSum(p))
	fmt.Println("Triangle:            ", n*(n+1)/2)
	// Output:
	// Simple:   1.0000000000146203e+20
	// AccSum:   1.0000000000147541e+20
	// Triangle:             1475412681
}

func ExampleCond() {
	p := []float64{1e100, 1e-100, -1e100}
	fmt.Println(accsum.Cond(p))
	// Output: 2e+200
}

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
	fmt.Printf("PrecSum:  %.16e\n", accsum.PrecSum(p, 22))
	fmt.Println("Triangle:            ", n*(n+1)/2)
	// Output:
	// Simple:   1.0000000000146203e+20
	// PrecSum:  1.0000000000147541e+20
	// Triangle:             1475412681
}
