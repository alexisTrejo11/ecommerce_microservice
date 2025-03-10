package usecase

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/google/uuid"
)

type CourseUseCaseImpl struct {
	courseRepository output.CourseRepository
}

func NewCourseUseCase(courseRepository output.CourseRepository) input.CourseUseCase {
	return &CourseUseCaseImpl{
		courseRepository: courseRepository,
	}
}

func (us *CourseUseCaseImpl) GetCourseById(id string) (*dtos.CourseDTO, error) {
	return nil, nil
}

func (us *CourseUseCaseImpl) GetCoursesByCategory(category string) (*[]dtos.CourseDTO, error) {
	return nil, nil
}

func (us *CourseUseCaseImpl) GetCoursesByInstructorId(instructorId string) (*[]dtos.CourseDTO, error) {
	return nil, nil
}

func (us *CourseUseCaseImpl) CourseSearch() (*[]dtos.CourseDTO, error) {
	return nil, nil
}

func (us *CourseUseCaseImpl) CreateCourse(dto dtos.CourseInsertDTO) (*dtos.CourseDTO, error) {
	return nil, nil
}

func (us *CourseUseCaseImpl) UpdateCourse(id uuid.UUID, dto dtos.CourseInsertDTO) (*dtos.CourseDTO, error) {
	return nil, nil
}

func (us *CourseUseCaseImpl) DeleteCourse(id uuid.UUID) error {
	return nil
}
