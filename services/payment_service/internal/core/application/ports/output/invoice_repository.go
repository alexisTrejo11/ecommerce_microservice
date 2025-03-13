package output

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/entity"
)

type InvoiceRepository interface {
	SaveInvoice(ctx context.Context, invoice *entity.Invoice) error
	GetInvoiceByID(ctx context.Context, id string) (*entity.Invoice, error)
	UpdateInvoiceStatus(ctx context.Context, id string, status entity.InvoiceStatus) error
	ListInvoicesByCustomer(ctx context.Context, customerID string) ([]*entity.Invoice, error)
	ListInvoicesByStatus(ctx context.Context, status entity.InvoiceStatus) ([]*entity.Invoice, error)
}
