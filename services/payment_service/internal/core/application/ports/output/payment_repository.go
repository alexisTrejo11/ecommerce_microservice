package output

import (
	"context"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/entity"
)

type PaymentRepository interface {
	SavePaymentIntent(ctx context.Context, pi *entity.PaymentIntent) error
	GetPaymentIntentByID(ctx context.Context, id string) (*entity.PaymentIntent, error)
	UpdatePaymentIntentStatus(ctx context.Context, id string, status entity.PaymentIntentStatus) error
	ListPaymentIntentsByCustomer(ctx context.Context, customerID string) ([]*entity.PaymentIntent, error)
	ListPaymentIntentsByDateRange(ctx context.Context, fromDate, toDate time.Time) ([]*entity.PaymentIntent, error)
}
