package input

import (
	"context"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/entity"
	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/valueobject"
)

type PaymentUseCase interface {
	CreatePaymentIntent(ctx context.Context, customerID string, amount valueobject.Money, paymentMethodID string) (*entity.PaymentIntent, error)
	ConfirmPaymentIntent(ctx context.Context, paymentIntentID string) error
	GetPaymentIntentByID(ctx context.Context, paymentIntentID string) (*entity.PaymentIntent, error)
	GetUserPayments(ctx context.Context, customerID string) ([]*entity.PaymentIntent, error)
	GetAllPayments(ctx context.Context, fromDate, toDate time.Time) ([]*entity.PaymentIntent, error)
}
