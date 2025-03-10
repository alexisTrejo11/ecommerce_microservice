package output

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/google/uuid"
)

type ResourceRepository interface {
	GetById(ctx context.Context, id string) (*domain.Resource, error)
	Create(ctx context.Context, newResource domain.Resource) (*domain.Resource, error)
	Update(ctx context.Context, id uuid.UUID, updatedResource domain.Resource) (*domain.Resource, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
