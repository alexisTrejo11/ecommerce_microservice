package output

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/entity"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, customer *entity.Customer) error
	AddPaymentMethodToCustomer(ctx context.Context, customerID, paymentMethodID string) error
	GetCustomerByID(ctx context.Context, id string) (*entity.Customer, error)
}
