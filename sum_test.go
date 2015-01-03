// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package accsum_test

import (
	"fmt"

	"github.com/soniakeys/accsum"
)

func ExampleCondDot() {
	x := []float64{1e10, 1, 1e10}
	y := []float64{1e10, 1, -1e10}
	fmt.Println(accsum.CondDot(accsum.Dot2, x, y))
	// Output: 4e+20
}

func ExampleCondSum() {
	p := []float64{1e100, 1, -1e100}
	fmt.Println(accsum.CondSum(accsum.Sum2, p))
	// Output: 2e+100
}

func ExampleDot() {
	x := []float64{1, 2, 3}
	y := []float64{3, 1, 4}
	fmt.Println(accsum.Dot(x, y))
	// Output: 17
}

func ExampleKahanSum() {
	n := 54321
	p := make([]float64, n+1)
	for i := range p {
		p[i] = float64(i)
	}
	p[0] = 1e20
	fmt.Printf("Simple:    %.16e\n", accsum.Sum(p))
	fmt.Printf("KahanSum:  %.16e\n", accsum.KahanSum(p))
	fmt.Println("Triangle:             ", n*(n+1)/2)
	// Output:
	// Simple:    1.0000000000146203e+20
	// KahanSum:  1.0000000000147541e+20
	// Triangle:              1475412681
}

func ExampleKahanB() {
	n := 54321
	p := make([]float64, n+1)
	for i := range p {
		p[i] = float64(i)
	}
	p[0] = 1e20
	fmt.Printf("Simple:   %.16e\n", accsum.Sum(p))
	fmt.Printf("KahanB:   %.16e\n", accsum.KahanB(p))
	fmt.Println("Triangle:            ", n*(n+1)/2)
	// Output:
	// Simple:   1.0000000000146203e+20
	// KahanB:   1.0000000000147541e+20
	// Triangle:             1475412681
}

func ExamplePairSum() {
	n := 54321
	p := make([]float64, n+1)
	for i := range p {
		p[i] = float64(i)
	}
	p[0] = 1e20
	fmt.Printf("Simple:   %.16e\n", accsum.Sum(p))
	fmt.Printf("PairSum:  %.16e\n", accsum.PriestSum(p))
	fmt.Println("Triangle:            ", n*(n+1)/2)
	// Output:
	// Simple:   1.0000000000146203e+20
	// PairSum:  1.0000000000147541e+20
	// Triangle:             1475412681
}

func ExamplePriestSum() {
	n := 54321
	p := make([]float64, n+1)
	for i := range p {
		p[i] = float64(i)
	}
	p[0] = 1e20
	fmt.Printf("Simple:     %.16e\n", accsum.Sum(p))
	fmt.Printf("PriestSum:  %.16e\n", accsum.PriestSum(p))
	fmt.Println("Triangle:              ", n*(n+1)/2)
	// Output:
	// Simple:     1.0000000000146203e+20
	// PriestSum:  1.0000000000147541e+20
	// Triangle:               1475412681
}

func ExampleSum() {
	p := []float64{1, 2, 3, 4}
	fmt.Println(accsum.Sum(p))
	// Output: 10
}

func ExampleXDot() {
	n := 4321
	x := make([]float64, n+1)
	for i := range x {
		x[i] = float64(i)
	}
	x[0] = 1e11
	fmt.Printf("Simple:   %.16e\n", accsum.Dot(x, x))
	fmt.Printf("XDot:     %.16e\n", accsum.XDot(x, x))
	fmt.Println("Square triangle:      ", n*(n+1)*(2*n+1)/6)
	// Output:
	// Simple:   1.0000000000026734e+22
	// XDot:     1.0000000000026902e+22
	// Square triangle:       26901858961
}

func ExampleXSum() {
	n := 54321
	p := make([]float64, n+1)
	for i := range p {
		p[i] = float64(i)
	}
	p[0] = 1e20
	fmt.Printf("Simple:   %.16e\n", accsum.Sum(p))
	fmt.Printf("XSum:     %.16e\n", accsum.XSum(p))
	fmt.Println("Triangle:            ", n*(n+1)/2)
	// Output:
	// Simple:   1.0000000000146203e+20
	// XSum:     1.0000000000147541e+20
	// Triangle:             1475412681
}
