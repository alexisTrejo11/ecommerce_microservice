package repository

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseRepositoryImpl struct {
	db gorm.DB
}

func NewCourseRepository(db gorm.DB) input.CourseUseCase {
	return &CourseRepositoryImpl{
		db: db,
	}
}

func (r *CourseRepositoryImpl) GetCourseById(id string) (*dtos.CourseDTO, error) {
	return nil, nil
}

func (r *CourseRepositoryImpl) GetCoursesByCategory(category string) (*[]dtos.CourseDTO, error) {
	return nil, nil
}

func (r *CourseRepositoryImpl) GetCoursesByInstructorId(instructorId string) (*[]dtos.CourseDTO, error) {
	return nil, nil
}

func (r *CourseRepositoryImpl) CourseSearch() (*[]dtos.CourseDTO, error) {
	return nil, nil
}

func (r *CourseRepositoryImpl) CreateCourse(dto dtos.CourseInsertDTO) (*dtos.CourseDTO, error) {
	return nil, nil
}

func (r *CourseRepositoryImpl) UpdateCourse(id uuid.UUID, dto dtos.CourseInsertDTO) (*dtos.CourseDTO, error) {
	return nil, nil
}

func (r *CourseRepositoryImpl) DeleteCourse(id uuid.UUID) error {
	return nil
}
