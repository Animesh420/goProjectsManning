package money

import "fmt"

// ExchangeRate represents a rate to convert from one currency
type ExchangeRate Decimal

func GetExchangeRate(value string) ExchangeRate {
	unitExchangeRate, err := ParseDecimal(value)
	if err != nil {
		panic("Unable to create unit exchange rate")
	}
	return ExchangeRate(unitExchangeRate)
}

// Convert applies the change rate to convert an amount to a target currency
func Convert(amount Amount, to Currency, rates exchangeRates) (Amount, error) {

	r, err := rates.FetchExchangeRate(amount.currency, to)
	if err != nil {
		return Amount{}, fmt.Errorf("cannot get change rate: %w", err)
	}
	// Convert to the target currency applying to fetched change rate
	convertedValue := applyExchangeRate(amount, to, r)

	// Validate the converted amount is in the handled bounded range
	if err := convertedValue.validate(); err != nil {
		return Amount{}, err
	}

	return convertedValue, nil
}

// multiply a decimal with exchange rate and return the product

func multiply(d Decimal, rate ExchangeRate) (Decimal, error) {
	dec := Decimal{
		subunits:  d.subunits * rate.subunits,
		precision: d.precision + rate.precision,
	}

	// Clean the representation by removing trailing zeros
	dec.simplify()
	return dec, nil
}

func applyExchangeRate(a Amount, target Currency, rate ExchangeRate) Amount {
	converted, err := multiply(a.quantity, rate)
	if err != nil {
		return Amount{}
	}

	switch {
	case converted.precision > target.precision:
		converted.subunits = converted.subunits / pow10(converted.precision-target.precision)
	case converted.precision < target.precision:
		converted.subunits = converted.subunits * pow10(target.precision-converted.precision)
	}
	converted.precision = target.precision
	return Amount{
		currency: target,
		quantity: converted,
	}
}
