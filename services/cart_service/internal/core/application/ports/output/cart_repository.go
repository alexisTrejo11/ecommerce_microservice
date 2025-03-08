package output

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/core/domain"
	"github.com/google/uuid"
)

type ProductRepository interface {
	GetById(ctx context.Context, id uuid.UUID) (*domain.Cart, error)
}
