package errors

import "errors"

var (
	ErrAddressValidation   = errors.New("address validation error")
	ErrInvalidPostalCode   = errors.New("invalid postal code format")
	ErrInvalidCountry      = errors.New("invalid country code")
	ErrRequiredField       = errors.New("field is required")
	ErrMultipleDefaults    = errors.New("only one default address allowed")
	ErrForbbiden           = errors.New("not allowed to get this data")
	ErrAccountBanned       = errors.New("account has been banned")
	ErrAccountNotActivated = errors.New("account has not been activated")
)
