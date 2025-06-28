package money

type exchangeRates interface {
	FetchExchangeRate(source, target Currency) (ExchangeRate, error)
}