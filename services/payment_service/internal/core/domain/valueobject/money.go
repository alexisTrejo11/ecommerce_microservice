package valueobject

import "fmt"

type Money struct {
	Amount   float64
	Currency Currency
}

func NewMoney(amount float64, currency Currency) (Money, error) {
	if amount < 0 {
		return Money{}, fmt.Errorf("amount cannot be negative")
	}

	return Money{Amount: amount, Currency: currency}, nil
}
