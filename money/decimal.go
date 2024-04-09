package money

import (
	"fmt"
	"strconv"
	"strings"
)

// Decimal can represent a floating-point number with a fixed precision.
// example: 1.52 = 152 * 10^(-2) will be stored as {152, 2}
type Decimal struct {

	// subunits is the amount of subunits.
	// Multiply it by the precion to get the real value.
	subunits int64

	// number of subunits in a unit, expressed as a power of 10.
	precision byte
}

// ParseDecimal converts a string into its Decimal representation.
// It assumes there is up to one decimal separator, and that
// the separator is '.' (full stop character).
func ParseDecimal(value string) (Decimal, error) {
	intPart, fracPart, _ := strings.Cut(value, ".")

	// maxDecimal is the number of digits in a thousand billion.
	const maxDecimal = 12

	if len(intPart) > maxDecimal {
		return Decimal{}, ErrTooLarge
	}

	subunits, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err.Error())
	}

	precision := byte(len(fracPart))

	return Decimal{subunits: subunits, precision: precision}, nil
}
