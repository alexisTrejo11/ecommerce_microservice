package output

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/entity"
	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/valueobject"
)

type PaymentRepository interface {
	Save(ctx context.Context, payment *entity.Payment) error
	GetByID(ctx context.Context, id valueobject.ID) (*entity.Payment, error)
	GetByUserID(ctx context.Context, userID valueobject.ID) ([]*entity.Payment, error)
	GetByExternalID(ctx context.Context, externalID string) (*entity.Payment, error)
	Update(ctx context.Context, payment *entity.Payment) error
}
