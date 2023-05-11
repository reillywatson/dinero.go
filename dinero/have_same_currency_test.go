package dinero_test

import (
	"testing"

	"github.com/DustinJSilk/dinero.go/currency"
	"github.com/DustinJSilk/dinero.go/dinero"
)

func TestHaveSameCurrency(t *testing.T) {
	type test struct {
		description string
		a           dinero.Dinero[int]
		b           dinero.Dinero[int]
		expect      bool
	}

	tests := []test{
		{
			description: "returns true when currencies are equal",
			a:           dinero.NewDinero(2000, currency.USD),
			b:           dinero.NewDinero(1000, currency.USD),
			expect:      true,
		},
		{
			description: "returns false when currencies are not equal",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(500, currency.EUR),
			expect:      false,
		},
		{
			description: "returns true when currencies are structurally equal",
			a:           dinero.NewDinero(500, currency.NewCurrency("USD", 10, 2)),
			b:           dinero.NewDinero(500, currency.NewCurrency("USD", 10, 2)),
			expect:      true,
		},
	}

	for _, tc := range tests {
		got := dinero.HaveSameCurrency(tc.a, tc.b)

		if tc.expect != got {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func BenchmarkHaveSameCurrency(b *testing.B) {
	da := dinero.NewDinero(100, currency.USD)
	db := dinero.NewDinero(100, currency.USD)

	for i := 0; i < b.N; i++ {
		dinero.HaveSameCurrency(da, db)
	}
}
