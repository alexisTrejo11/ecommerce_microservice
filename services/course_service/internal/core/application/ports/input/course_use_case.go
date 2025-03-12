package input

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/google/uuid"
)

type CourseUseCase interface {
	GetCourseById(ctx context.Context, id uuid.UUID) (*dtos.CourseDTO, error)
	GetCoursesByCategory(ctx context.Context, category domain.CourseCategory) (*[]dtos.CourseDTO, error)
	GetCoursesByInstructorId(ctx context.Context, instructorId string) (*[]dtos.CourseDTO, error)
	CourseSearch(ctx context.Context) (*[]dtos.CourseDTO, error)
	CreateCourse(ctx context.Context, dto dtos.CourseInsertDTO) (*dtos.CourseDTO, error)
	UpdateCourse(ctx context.Context, id uuid.UUID, dto dtos.CourseInsertDTO) (*dtos.CourseDTO, error)
	PublishCourse(ctx context.Context, id uuid.UUID) error
	EnrollStudentInCourse(ctx context.Context, courseId uuid.UUID) error
	UpdateCourseRating(ctx context.Context, courseId uuid.UUID, rating float64) error
	DeleteCourse(ctx context.Context, id uuid.UUID) error
}
