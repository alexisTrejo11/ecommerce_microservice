package input

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/entity"
)

type CustomerUseCase interface {
	CreateCustomer(ctx context.Context, email, name string) (*entity.Customer, error)
	AddPaymentMethodToCustomer(ctx context.Context, customerID, paymentMethodID string) error
}
