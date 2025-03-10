package output

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/google/uuid"
)

type ModuleRepository interface {
	GetById(ctx context.Context, id string) (*domain.Module, error)
	GetByCourseId(ctx context.Context, id string) (*[]domain.Module, error)
	Create(ctx context.Context, newModule domain.Module) (*domain.Module, error)
	Update(ctx context.Context, id uuid.UUID, updatedModule domain.Module) (*domain.Module, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
