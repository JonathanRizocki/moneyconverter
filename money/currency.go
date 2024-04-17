package money

// Currency defines the code of a money and its decimal position.
type Currency struct {
	code      string
	precision byte
}

const (
	// ErrInvalidCurrencyCode is returned when the currency to parse is not a standard 3-letter code.
	ErrInvalidCurrencyCode = Error("invalid currency code")
)

func ParseCurrency(code string) (Currency, error) {
	if len(code) != 3 {
		return Currency{}, ErrInvalidCurrencyCode
	}

	switch code {
	case "IRR":
		return Currency{code: code, precision: 0}, nil
	case "CNY", "VND":
		return Currency{code: code, precision: 1}, nil
	case "BHD", "IQD", "KWD", "LYD", "OMR", "TND":
		return Currency{code: code, precision: 3}, nil
	default:
		return Currency{code: code, precision: 2}, nil
	}
}

// String implements Stringer.
func (c Currency) String() string {
	return c.code + " ||| \n"
}

// Code returns the ISO code for the currency.
func (c Currency) Code() string {
	return c.code
}
