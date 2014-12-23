// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package accsum_test

import (
	"fmt"

	"github.com/soniakeys/accsum"
)

func ExampleCond() {
	p := []float64{1e100, 1e-100, -1e100}
	fmt.Println(accsum.Cond(accsum.Sum2, p))
	// Output: 2e+200
}

func ExampleSum() {
	p := []float64{1, 2, 3, 4}
	fmt.Println(accsum.Sum(p))
	// Output: 10
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
