package input

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/google/uuid"
)

type CourseUseCase interface {
	GetCourseById(id uuid.UUID)
	GetCoursesByCategory(category domain.CourseCategory)
	GetCoursesByInstructorId(instructorId uuid.UUID)
	CourseSearch()
	CreateCourse()
	UpdateCourse()
	DeleteCourse(id uuid.UUID)
}
