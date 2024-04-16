package money

import (
	"errors"
	"testing"
)

func TestNewAmount(t *testing.T) {
	tt := map[string]struct {
		quantity Decimal
		currency Currency
		want     Amount
		err      error
	}{
		"$15.46 USD": {
			quantity: Decimal{subunits: 1546, precision: 2},
			currency: Currency{code: "USD", precision: 2},
			want: Amount{
				quantity: Decimal{subunits: 1546, precision: 2},
				currency: Currency{code: "USD", precision: 2},
			},
		},
		"$15.00 USD": {
			quantity: Decimal{subunits: 1500, precision: 2},
			currency: Currency{code: "USD", precision: 2},
			want: Amount{
				quantity: Decimal{subunits: 1500, precision: 2},
				currency: Currency{code: "USD", precision: 2},
			},
		},
		"$1.50 USD": {
			quantity: Decimal{subunits: 150, precision: 2},
			currency: Currency{code: "USD", precision: 2},
			want: Amount{
				quantity: Decimal{subunits: 150, precision: 2},
				currency: Currency{code: "USD", precision: 2},
			},
		},
		"$0.15 USD": {
			quantity: Decimal{subunits: 15, precision: 2},
			currency: Currency{code: "USD", precision: 2},
			want: Amount{
				quantity: Decimal{subunits: 15, precision: 2},
				currency: Currency{code: "USD", precision: 2},
			},
		},
		"$1.500 USD": {
			quantity: Decimal{subunits: 1500, precision: 3},
			currency: Currency{code: "USD", precision: 2},
			err:      ErrTooPrecise,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := NewAmount(tc.quantity, tc.currency)

			if !errors.Is(err, tc.err) {
				t.Errorf("Expected error %v, but got %v", tc.err, err)
			}

			if got != tc.want {
				t.Errorf("Expected amount %v, but got %v", tc.want, got)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tt := map[string]struct {
		in  Amount
		err error
	}{
		"nominal": {
			in: Amount{
				quantity: Decimal{subunits: 1_000_000_000, precision: 2},
				currency: Currency{code: "USD", precision: 2},
			},
			err: nil,
		},
		"too precise": {
			in: Amount{
				quantity: Decimal{subunits: 1_000, precision: 50},
				currency: Currency{code: "USD", precision: 2},
			},
			err: ErrTooPrecise,
		},
		"too large": {
			in: Amount{
				quantity: Decimal{subunits: 1_000_000_000_000_000, precision: 2},
				currency: Currency{code: "USD", precision: 2},
			},
			err: ErrTooLarge,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {

			err := tc.in.validate()
			if !errors.Is(err, tc.err) {
				t.Errorf("Expected error %v, but got %v", tc.err, err)
			}
		})
	}
}
