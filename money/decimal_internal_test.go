package money

import (
	"errors"
	"reflect"
	"testing"
)

func TestParseDecimal(t *testing.T) {
	tt := map[string]struct {
		decimal  string
		expected Decimal
		err      error
	}{
		"2 decimal digits": {
			decimal:  "1.52",
			expected: Decimal{152, 2},
			err:      nil,
		},
		"no decimal digits": {
			decimal:  "1",
			expected: Decimal{1, 0},
			err:      nil,
		},
		"suffix 0 as decimal digits": {
			decimal:  "1.50",
			expected: Decimal{15, 1},
			err:      nil,
		},
		"prefix 0 as decimal digits": {
			decimal:  "1.02",
			expected: Decimal{102, 2},
			err:      nil,
		},
		"multiple of 10": {
			decimal:  "150",
			expected: Decimal{150, 0},
			err:      nil,
		},
		"invalid decimal part": {
			decimal: "65.pocket",
			err:     ErrInvalidDecimal,
		},
		"Not a number": {
			decimal: "NaN",
			err:     ErrInvalidDecimal,
		},
		"empty string": {
			decimal: "",
			err:     ErrInvalidDecimal,
		},
		"too large": {
			decimal: "1234567890123",
			err:     ErrTooLarge,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ParseDecimal(tc.decimal)

			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}

			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestPow10(t *testing.T) {
	tt := map[string]struct {
		in       byte
		expected int64
	}{
		"10^0": {
			in:       0,
			expected: 1,
		},
		"10^1": {
			in:       1,
			expected: 10,
		},
		"10^2": {
			in:       2,
			expected: 100,
		},
		"10^3": {
			in:       3,
			expected: 1000,
		},
		"10^4": {
			in:       4,
			expected: 10_000,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := pow10(tc.in)

			if got != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, got)
			}
		})
	}
}

func TestSimplify(t *testing.T) {
	tt := map[string]struct {
		in  Decimal
		exp Decimal
	}{
		"1, precision 2 -> 1,2": {
			in:  Decimal{subunits: 1, precision: 2},
			exp: Decimal{subunits: 1, precision: 2},
		},
		"1, precision 5 -> 1,5": {
			in:  Decimal{subunits: 1, precision: 5},
			exp: Decimal{subunits: 1, precision: 5},
		},
		"10000, precision 10 -> 1, 6": {
			in:  Decimal{subunits: 10000, precision: 10},
			exp: Decimal{subunits: 1, precision: 6},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tc.in.simplify()

			if !reflect.DeepEqual(tc.in, tc.exp) {
				t.Errorf("Expected %v, but got %v", tc.exp, tc.in)
			}
		})
	}
}
