package valueobject

import "fmt"

type PaymentMethod string

const (
	Card   PaymentMethod = "card"
	Stripe PaymentMethod = "stripe"
	Paypal PaymentMethod = "paypal"
)

func NewPaymentMethod(value string) (PaymentMethod, error) {
	switch value {
	case string(Card), string(Stripe), string(Paypal):
		return PaymentMethod(value), nil
	default:
		return "", fmt.Errorf("invalid payment method: %s", value)
	}
}
