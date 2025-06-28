package ecbank

// ecbankError defines a sentinel error
type ecbankError string

// ecbankError implements the error interface
func (e ecbankError) Error() string {
	return string(e)
}

const (
	ErrCallingServer      = ecbankError("error calling server")
	ErrClientSide         = ecbankError("error client side")
	ErrServerSide         = ecbankError("error server side")
	ErrUnknownStatusCode  = ecbankError("unknown status code")
	ErrUnexpectedFormat   = ecbankError("unexpected format")
	ErrChangeRateNotFound = ecbankError("change rate not found")
)
