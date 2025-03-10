package output

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/google/uuid"
)

type LessonRepository interface {
	GetById(id string) (*domain.Lesson, error)
	Create(newLesson domain.Lesson) (*domain.Lesson, error)
	Update(id uuid.UUID, updatedLesson domain.Lesson) (*domain.Lesson, error)
	Delete(id uuid.UUID) error
}
