// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package accsum_test

import (
	"fmt"

	"github.com/soniakeys/accsum"
)

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
