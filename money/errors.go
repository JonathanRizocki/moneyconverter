package money

// Error defines an error
type Error string

// Error implements the error interface.
func (e Error) Error() string {
	return string(e)
}

const (
	// ErrInvalidDecimal is return if the decimal is malformed
	ErrInvalidDecimal = Error("unable to conver the decimal")

	// ErrTooLarge is returned if the quantity is too large.
	// This would cause floating point precision errors.
	ErrTooLarge = Error("quantity over 10^12 is too large")

	// ErrInvalidCurrencyCode is returned when the currency to parse is not a standard 3-letter code.
	ErrInvalidCurrencyCode = Error("invalid currency code")
)
