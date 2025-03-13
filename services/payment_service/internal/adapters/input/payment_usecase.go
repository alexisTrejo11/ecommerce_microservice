package input

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/entity"
	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/valueobject"
)

type PaymentUseCase interface {
	ProcessPayment(ctx context.Context, amount float64, currency string, method string) (*entity.PaymentIntent, error)
	GetPaymentByID(ctx context.Context, id valueobject.ID) (*entity.PaymentIntent, error)
}
