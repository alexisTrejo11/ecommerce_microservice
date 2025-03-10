package input

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/google/uuid"
)

type CourseUseCase interface {
	GetCourseById(id string) (*dtos.CourseDTO, error)
	GetCoursesByCategory(category string) (*[]dtos.CourseDTO, error)
	GetCoursesByInstructorId(instructorId string) (*[]dtos.CourseDTO, error)
	CourseSearch() (*[]dtos.CourseDTO, error)
	CreateCourse(dto dtos.CourseInsertDTO) (*dtos.CourseDTO, error)
	UpdateCourse(id uuid.UUID, dto dtos.CourseInsertDTO) (*dtos.CourseDTO, error)
	DeleteCourse(id uuid.UUID) error
}
