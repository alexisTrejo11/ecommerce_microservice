package valueobject

import "fmt"

type Currency string

var supportedCurrencies = map[string]bool{
	"USD": true,
	"EUR": true,
	"MXN": true,
}

func NewCurrency(value string) (Currency, error) {
	if _, ok := supportedCurrencies[value]; !ok {
		return "", fmt.Errorf("unsupported currency: %s", value)
	}
	return Currency(value), nil
}
