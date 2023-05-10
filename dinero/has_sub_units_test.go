package dinero_test

import (
	"testing"

	"dinero.go/currency"
	"dinero.go/dinero"
)

func TestHasSubUnits(t *testing.T) {
	type test struct {
		description string
		value       dinero.Dinero[int]
		expect      bool
	}

	tests := []test{
		// decimal currencies
		{
			description: "returns false when there are no sub-units",
			value:       dinero.NewDinero(1100, currency.USD),
			expect:      false,
		},
		{
			description: "returns true when there are sub-units based on a custom scale",
			value:       dinero.NewDineroWithScale(1100, currency.USD, 3),
			expect:      true,
		},
		{
			description: "returns true when there are sub-units",
			value:       dinero.NewDinero(1150, currency.USD),
			expect:      true,
		},
		{
			description: "returns false when there are no sub-units based on a custom scale",
			value:       dinero.NewDineroWithScale(1150, currency.USD, 1),
			expect:      false,
		},
		// non-decimal currencies'
		{
			description: "returns false when there are no sub-units",
			value:       dinero.NewDinero(10, currency.MGA),

			expect: false,
		},
		{
			description: "returns true when there are sub-units",
			value:       dinero.NewDinero(11, currency.MGA),
			expect:      true,
		},
	}

	for _, tc := range tests {
		got := tc.value.HasSubUnits()

		if tc.expect != got {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func BenchmarkHasSubUnits(b *testing.B) {
	da := dinero.NewDinero(100, currency.USD)

	for i := 0; i < b.N; i++ {
		da.HasSubUnits()
	}
}