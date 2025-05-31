package money

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ErrInvalidDecimal = Error("unable to convert the decimal")
	ErrTooLarge       = Error("quantity over 10^12 is too large")
)

type Decimal struct {
	// subunits is the amount  of subunits
	// Multiply it with precision to get the real value
	subunits int64

	// Number of "subunits" in a unit, expreseesed as a power of 10,
	precision byte
}

// maxDecimal value is a thousand billion, using the short scale
const maxDecimal = 1e12

func ParseDecimal(value string) (Decimal, error) {
	intPart, fracPart, _ := strings.Cut(value, ".")
	subunits, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err.Error())
	}

	if subunits > maxDecimal {
		return Decimal{}, ErrTooLarge
	}

	precision := byte(len(fracPart))
	return Decimal{subunits: subunits, precision: precision}, nil
}

func (d *Decimal) simplify() {
	for d.subunits%10 == 0 && d.precision > 0 {
		d.precision--
		d.subunits /= 10
	}
}
