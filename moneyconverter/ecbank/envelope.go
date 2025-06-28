package ecbank

import (
	"encoding/xml"
	"fmt"
	"io"
	"learngo-pockets/moneyconverter/money"
)

// structures used for XML decoding

const baseCurrencyCode = "EUR"

type envelope struct {
	Rates []currencyRate `xml:"Cube>Cube>Cube"`
}

type currencyRate struct {
	Currency string  `xml:"currency,attr"`
	Rate     float64 `xml:"rate,attr"`
}

// exchangeRates builds a map of all the supported exchange rates.
func (e envelope) mappedChangeRates() map[string]float64 {

	rates := make(map[string]float64, len(e.Rates)+1)

	for _, c := range e.Rates {
		rates[c.Currency] = c.Rate
	}

	rates[baseCurrencyCode] = 1.0
	return rates

}

// exchangeRate reads the change rate from the Envelope's contents.
func (e envelope) exchangeRate(source, target string) (money.ExchangeRate, error) {
	if source == target {
		return money.GetExchangeRate("1.0"), nil
	}

	rates := e.mappedChangeRates()

	sourceFactor, sourceFound := rates[source]
	if !sourceFound {
		return money.GetExchangeRate("0.0"), fmt.Errorf("failed to find the source currency %s", source)
	}

	targetFactor, targetFound := rates[target]
	if !targetFound {
		return money.GetExchangeRate("0.0"), fmt.Errorf("failed to find the target currency %s", target)
	}

	rate, err := money.ParseDecimal(fmt.Sprintf("%.6f", targetFactor/sourceFactor))
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("unable to parse exchange rate from %s to %s: %w", source, target, err)
	}

	return money.ExchangeRate(rate), nil
}

func readRateFromResponse(source, target string, respBody io.Reader) (money.ExchangeRate, error) {
	// read the response
	decoder := xml.NewDecoder(respBody)

	var ecbMessage envelope
	err := decoder.Decode(&ecbMessage)
	if err != nil {
		return money.GetExchangeRate("0.0"), fmt.Errorf("%w: %s", ErrUnexpectedFormat, err)
	}

	rate, err := ecbMessage.exchangeRate(source, target)
	if err != nil {
		return money.GetExchangeRate("0.0"), fmt.Errorf("%w: %s", ErrChangeRateNotFound, err)
	}

	return rate, nil
}
