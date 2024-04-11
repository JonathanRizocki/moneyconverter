package money

type Amount struct {
	quantity Decimal
	currency Currency
}

const (
	// ErrTooPrecise is returned if the number is too precise for
	// the currency
	ErrTooPrecise = Error("quantity is too precise")
)

func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	if quantity.precision > currency.precision {
		// In order to avoid converting 0.00001 cent, let's exit now
		return Amount{}, ErrTooPrecise
	}

	quantity.precision = currency.precision

	return Amount{quantity: quantity, currency: currency}, nil
}
