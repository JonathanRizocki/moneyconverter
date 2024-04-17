package ecb

import (
	"fmt"
	"learngo-pockets/moneyconverter/money"
	"net/http"
)

// EuroCentralBank can call the bank to retrieve exchange rates.
type EuroCentralBank struct {
}

const (
	ErrClientSide        = ecbankError("client side error when contacting ECB")
	ErrSeverSide         = ecbankError("server side error when contacting ECB")
	ErrUnknownStatusCode = ecbankError("unknown status code contacting ECB")
)

// FetchExchangeRate fetches the ExchangeRate for the day and returns it.
func (ecb EuroCentralBank) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {

	const euroxrefURL = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

	resp, err := http.Get(euroxrefURL)
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrSeverSide, err.Error())
	}
	defer resp.Body.Close()

	if err = checkStatusCode(resp.StatusCode); err != nil {
		return money.ExchangeRate{}, err
	}

	fmt.Printf("response %v", resp)
	return money.ExchangeRate{}, nil
}

const (
	clientErrorClass = 4
	serverErrorClass = 5
)

// checkStatusCode returns a different error depending on the returned
// status code.
func checkStatusCode(statusCode int) error {
	switch {
	case statusCode == http.StatusOK:
		return nil
	case httpStatusClass(statusCode) == clientErrorClass:
		return fmt.Errorf("%w: %d", ErrClientSide, statusCode)
	case httpStatusClass(statusCode) == serverErrorClass:
		return fmt.Errorf("%w: %d", ErrSeverSide, statusCode)
	default:
		return fmt.Errorf("%w: %d", ErrUnknownStatusCode, statusCode)
	}

}

// httpStatusClass returns the class of a http status code.
func httpStatusClass(statusCode int) int {
	const httpErrorClassSize = 100
	return statusCode / httpErrorClassSize
}
