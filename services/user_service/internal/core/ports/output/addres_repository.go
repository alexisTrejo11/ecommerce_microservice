package output

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/google/uuid"
)

type AddressRepository interface {
	Create(ctx context.Context, address *entities.Address) error
	FindByID(ctx context.Context, id uint) (*entities.Address, error)
	FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]*entities.Address, error)
	FindDefaultByUserID(ctx context.Context, userID uuid.UUID) (*entities.Address, error)
	Update(ctx context.Context, address *entities.Address) error
	SetDefault(ctx context.Context, id uint, userID uuid.UUID) error
	Delete(ctx context.Context, id uint) error
}
