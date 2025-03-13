package input

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/entity"
	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/valueobject"
)

type InvoiceUseCase interface {
	CreateInvoice(ctx context.Context, subscriptionID string, amountDue valueobject.Money) (*entity.Invoice, error)
	MarkInvoiceAsPaid(ctx context.Context, invoiceID string) error
	GetUserInvoices(ctx context.Context, customerID string) ([]*entity.Invoice, error)
	GetAllInvoices(ctx context.Context, status entity.InvoiceStatus) ([]*entity.Invoice, error)
}
