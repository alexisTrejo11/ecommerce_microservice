package output

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/google/uuid"
)

type CourseRepository interface {
	GetById(id string) (*domain.Course, error)
	GetByCategory(category string) (*[]domain.Course, error)
	GetByInstructorId(instructorId string) (*[]domain.Course, error)
	Create(newCourse domain.Course) (*domain.Course, error)
	Update(id uuid.UUID, updatedCourse domain.Course) (*domain.Course, error)
	Delete(id uuid.UUID) error
}
