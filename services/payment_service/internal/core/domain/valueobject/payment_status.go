package valueobject

import "fmt"

type PaymentStatus string

const (
	Pending   PaymentStatus = "pending"
	Completed PaymentStatus = "completed"
	Failed    PaymentStatus = "failed"
)

func NewPaymentStatus(value string) (PaymentStatus, error) {
	switch value {
	case string(Pending), string(Completed), string(Failed):
		return PaymentStatus(value), nil
	default:
		return "", fmt.Errorf("invalid payment status: %s", value)
	}
}
