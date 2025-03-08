package output

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/core/domain"
	"github.com/google/uuid"
)

type CartRepository interface {
	CreateCart(ctx context.Context, cart domain.Cart) (*domain.Cart, error)
	UpdateCart(ctx context.Context, cart domain.Cart) (*domain.Cart, error)
	DeleteCart(ctx context.Context, id uuid.UUID) error
	GetById(ctx context.Context, id uuid.UUID) (*domain.Cart, error)
	GetByUserID(ctx context.Context, id uuid.UUID) (*domain.Cart, error)
}
