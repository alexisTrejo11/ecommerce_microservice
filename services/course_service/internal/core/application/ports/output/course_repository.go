package output

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/google/uuid"
)

type CourseRepository interface {
	GetById(ctx context.Context, id string) (*domain.Course, error)
	GetByCategory(ctx context.Context, category string) (*[]domain.Course, error)
	GetByInstructorId(ctx context.Context, instructorId string) (*[]domain.Course, error)
	Create(ctx context.Context, newCourse domain.Course) (*domain.Course, error)
	Update(ctx context.Context, id uuid.UUID, updatedCourse domain.Course) (*domain.Course, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
