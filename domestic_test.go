// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package accsum

import (
	"math"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	for i := 0; i < 1000; i++ {
		a := rand.ExpFloat64()
		if rand.Intn(2) == 0 {
			a = -a
		}
		x, y := split(a)
		if x+y != a {
			t.Fatalf(`split(%g)
results sum to       %g
want original value, %g`, a, x+y, a)
		}
		bx := math.Float64bits(x)
		by := math.Float64bits(y)
		// test low 26 bits of significand
		if bx&(1<<27-1) != 0 || by&(1<<27-1) != 0 {
			t.Fatalf(`split(%g) =
%064b
%064b
want 26 zero bits:
%s%s`, a, bx, by, strings.Repeat(".", 64-26), strings.Repeat("0", 26))
		}
	}
}

func TestNextPowerTwo(t *testing.T) {
	for _, tc := range []struct{ a, np2 float64 }{
		{15, 16},
		{16, 16}, // a power of 2 returns itself
		{17, 32},
		{1, 1},
		{2, 2},
		{0, 0},  // special case 0
		{-1, 1}, // it's next from absolute value
		{-3, 4},
	} {
		if got := nextPowerTwo(tc.a); got != tc.np2 {
			t.Fatalf("nextPowerTwo(%g) = %g, want %g\n", tc.a, got, tc.np2)
		}
	}
}

func TestExtractScalar(t *testing.T) {
	for _, tc := range []struct {
		exp   int
		q, pʹ float64
	}{
		{51, 3, 0.14159265358979312},
		{50, 3.25, -0.10840734641020688},
		{49, 3.125, 0.016592653589793116},
	} {
		σ := math.Ldexp(1, tc.exp) // 2^exp
		gotQ, gotPʹ := extractScalar(σ, math.Pi)
		if gotQ != tc.q || gotPʹ != tc.pʹ {
			t.Fatalf(`extractScalar(2^%d)
results: %g, %g
want:    %g, %g`, tc.exp, gotQ, gotPʹ, tc.q, tc.pʹ)
		}
	}
}

func TestExtractSlice(t *testing.T) {
	p := []float64{1, 22, 333}
	exp := 60
	τGot := extractSlice(math.Ldexp(1, exp), p)
	τWant := 256.
	pWant := []float64{1, 22, 77}
	if τGot != τWant || !reflect.DeepEqual(p, pWant) {
		t.Log("τ:   ", τGot)
		for i, x := range p {
			t.Logf("p[%d]: %.0f\n", i, x)
		}
		t.Fatal("Huh.")
	}
}
