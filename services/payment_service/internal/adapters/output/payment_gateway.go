package output

import (
	"context"
)

type PaymentGateway interface {
	Charge(ctx context.Context, amount float64, currency, method string) (transactionID string, err error)
}
