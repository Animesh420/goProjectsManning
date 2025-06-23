package money

const (
	// ErrTooPrecise is returned if the number is too precise for the currency
	ErrTooPrecise = Error("quantity is too precise")
)

// Amount defines a quantity of money in a given Currency.
type Amount struct {
	quantity Decimal
	currency Currency
}

// NewAmount returns an Amount of money.
func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	if quantity.precision > currency.precision {
		// In order to avoud converting 0.00001 cent, lets exit
		return Amount{}, ErrTooPrecise
	}

	quantity.precision = currency.precision
	return Amount{quantity: quantity, currency: currency}, nil
}
