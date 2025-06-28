package ecbank

import (
	"errors"
	"fmt"
	"learngo-pockets/moneyconverter/money"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	client http.Client
	url    string
}

const (
	clientErrorClass = 4
	serverErrorClass = 5
)

func NewClient(timeout time.Duration) Client {
	return Client{
		client: http.Client{Timeout: timeout},
	}
}

// checkStatusCode returns a different error
// depending on the returned status code
func checkStatusCode(statusCode int) error {
	switch {
	case statusCode == http.StatusOK:
		return nil
	case httpStatusClass(statusCode) == clientErrorClass:
		return fmt.Errorf("%w:%d", ErrClientSide, statusCode)
	case httpStatusClass(statusCode) == serverErrorClass:
		return fmt.Errorf("%w:%d", ErrServerSide, statusCode)
	default:
		return fmt.Errorf("%w: %d", ErrUnknownStatusCode, statusCode)
	}
}

// httpStatusClass returns the class of a http status code.
func httpStatusClass(statusCode int) int {
	const httpErrorClassSize = 100
	return statusCode / httpErrorClassSize
}

// FetchExchangeRate fetches the ExchangeRate for the day and returns it.
// func (c Client) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {
// 	const path = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"
// 	if c.url == "" {
// 		c.url = path
// 	}
// 	resp, err := http.Get(c.url)
// 	defer resp.Body.Close()
// 	if err != nil {
// 		return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrServerSide, err.Error())
// 	}
// 	return money.ExchangeRate{}, nil
// }

// 1xx (100 to 199) informational response—The request was received, and the server continues to process.
// 2xx (200 to 299) successful—The request was successfully received, understood, and accepted.
// 3xx (300 to 399) redirection—Further action needs to be taken to complete the request.
// 4xx (400 to 499) client error—The request contains bad syntax or can’t be fulfilled.
// 5xx (500 to 599) server error—The server failed to fulfill an apparently valid request.

// FetchExchangeRate fetches the ExchangeRate for the day and returns it
func (c Client) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {
	const path = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"
	if c.url == "" {
		c.url = path
	}

	resp, err := c.client.Get(c.url)

	if err != nil {
		var urlErr *url.Error

		if ok := errors.As(err, &urlErr); ok && urlErr.Timeout() {
			return money.GetExchangeRate("0.0"), fmt.Errorf("timeout calling the code, %s", err.Error())
		}
		return money.GetExchangeRate("0.0"), fmt.Errorf("%w: %s", ErrServerSide, err.Error())
	}

	defer resp.Body.Close()

	if err = checkStatusCode(resp.StatusCode); err != nil {
		return money.GetExchangeRate("0.0"), err
	}

	rate, err := readRateFromResponse(
		source.Code(), target.Code(), resp.Body)

	if err != nil {
		return money.GetExchangeRate("0.0"), err
	}

	return rate, nil

}
