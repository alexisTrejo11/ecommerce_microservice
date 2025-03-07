package input

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/application/dtos"
	"github.com/google/uuid"
)

type CartUseCase interface {
	CreateCart(ctx context.Context, userID uuid.UUID) error
	Buy(ctx context.Context, userID uuid.UUID, excludeItemsIDs []*uuid.UUID) error
	AddItems(ctx context.Context, userID uuid.UUID, items []dtos.CartItemFetchedDTO) error
	RemoveItems(ctx context.Context, userID uuid.UUID, itemIDs []uuid.UUID) error
	GetCartByUserId(ctx context.Context, userID uuid.UUID) (*dtos.CartDTO, error)
	GetCartById(ctx context.Context, id uuid.UUID) (*dtos.CartDTO, error)
	DeleteCart(ctx context.Context, userID uuid.UUID) error
}
