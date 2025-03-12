package input

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/google/uuid"
)

type ResourceUseCase interface {
	GetResourceById(ctx context.Context, id uuid.UUID) (*dtos.ResourceDTO, error)
	GetResourcesByLessonId(ctx context.Context, lessonId uuid.UUID) (*[]dtos.ResourceDTO, error)
	CreateResource(ctx context.Context, dto dtos.ResourceInsertDTO) (*dtos.ResourceDTO, error)
	UpdateResource(ctx context.Context, id uuid.UUID, dto dtos.ResourceInsertDTO) (*dtos.ResourceDTO, error)
	DeleteResource(ctx context.Context, id uuid.UUID) error
}
