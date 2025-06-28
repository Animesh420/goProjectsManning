package money

import (
	"reflect"
	"testing"
)

type stubRate struct {
	rate ExchangeRate
	err  error
}

func (m stubRate) FetchExchangeRate(_, _ Currency) (ExchangeRate, error) {
	return m.rate, m.err
}

func TestConvert(t *testing.T) {
	tt := map[string]struct {
		amount   Amount
		to       Currency
		validate func(t *testing.T, got Amount, err error)
	}{
		"34.98 USD to EUR": {
			amount: mustParseAmount(t, "34.98", "USD"),
			to:     mustParseCurrency(t, "EUR"),
			validate: func(t *testing.T, got Amount, err error) {
				if err != nil {
					t.Errorf("expected no error, got %s", err.Error())
				}
				expected := Amount{
					quantity: Decimal{
						subunits:  6996,
						precision: 2,
					},
					currency: Currency{code: "EUR", precision: 2},
				}
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("expected %v, got %v", expected, got)
				}

			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := Convert(tc.amount, tc.to, stubRate{rate: ExchangeRate{subunits: 2, precision: 0}})
			tc.validate(t, got, err)
		})
	}
}

func mustParseCurrency(t *testing.T, code string) Currency {
	t.Helper()

	currency, err := ParseCurrency(code)

	if err != nil {
		t.Fatalf("cannot parse currency %s code", code)
	}
	return currency
}

func mustParseAmount(t *testing.T, value string, code string) Amount {
	t.Helper()

	n, err := ParseDecimal(value)
	if err != nil {
		t.Fatalf("invalid number: %s", value)
	}

	currency, err := ParseCurrency(code)
	if err != nil {
		t.Fatalf("invalid currency code: %s", code)
	}

	amount, err := NewAmount(n, currency)
	if err != nil {
		t.Fatalf("cannot create amount with value %v and currency code %s", n, code)
	}

	return amount
}
