// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package accsum_test

import (
	"fmt"

	"github.com/soniakeys/accsum"
)

func ExampleSum() {
	p := []float64{1, 2, 3, 4}
	fmt.Println(accsum.Sum(p))
	// Output: 10
}

func ExampleCond() {
	p := []float64{1e100, 1e-100, -1e100}
	fmt.Println(accsum.Cond(accsum.Sum2, p))
	// Output: 2e+200
}
