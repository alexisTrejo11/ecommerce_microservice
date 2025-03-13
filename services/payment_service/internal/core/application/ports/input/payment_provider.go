package input

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/entity"
)

type PaymentProvider interface {
	CreatePaymentIntent(ctx context.Context, amount int64, currency string, customerID string, paymentMethodID string) (string, error)
	ConfirmPaymentIntent(ctx context.Context, paymentIntentID string) error
	GetPaymentIntent(ctx context.Context, paymentIntentID string) (*entity.PaymentIntent, error)
}
