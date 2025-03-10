package output

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/google/uuid"
)

type LessonRepository interface {
	GetById(ctx context.Context, id string) (*domain.Lesson, error)
	GetByModuleId(ctx context.Context, id string) (*[]domain.Lesson, error)
	Create(ctx context.Context, newLesson domain.Lesson) (*domain.Lesson, error)
	Update(ctx context.Context, id uuid.UUID, updatedLesson domain.Lesson) (*domain.Lesson, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
