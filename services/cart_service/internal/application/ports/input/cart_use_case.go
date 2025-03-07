package input

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/application/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/domain"
	"github.com/google/uuid"
)

type CartUseCase interface {
	CreateCart(ctx context.Context, userID uuid.UUID) error
	Buy(ctx context.Context, userID uuid.UUID, excludeItemsIDs []*uuid.UUID) error
	AddItems(ctx context.Context, userID uuid.UUID, items []dtos.CartItemFetchedDTO) error
	RemoveItems(ctx context.Context, userID uuid.UUID, itemIDs []uuid.UUID) error
	GetCart(ctx context.Context, userID uuid.UUID) (*domain.Cart, error)
	DeleteCart(ctx context.Context, userID uuid.UUID) error
}
