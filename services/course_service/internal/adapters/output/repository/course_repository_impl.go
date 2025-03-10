package repository

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/models"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseRepositoryImpl struct {
	db      gorm.DB
	mappers mappers.CourseMappers
}

func NewCourseRepository(db gorm.DB) output.CourseRepository {
	return &CourseRepositoryImpl{
		db: db,
	}
}

func (r *CourseRepositoryImpl) GetById(ctx context.Context, id string) (*domain.Course, error) {
	var courseModel models.CourseModel
	if err := r.db.WithContext(ctx).First(&courseModel, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return r.mappers.ModelToDomain(courseModel), nil
}

func (r *CourseRepositoryImpl) GetByCategory(ctx context.Context, category string) (*[]domain.Course, error) {
	var courseModels []models.CourseModel
	if err := r.db.WithContext(ctx).Where("category = ?", category).Find(&courseModels).Error; err != nil {
		return nil, err
	}

	return r.mappers.ModelsToDomains(courseModels), nil
}

func (r *CourseRepositoryImpl) GetByInstructorId(ctx context.Context, instructorId string) (*[]domain.Course, error) {
	var courseModels []models.CourseModel
	if err := r.db.WithContext(ctx).Where("instructor_id = ?", instructorId).Find(&courseModels).Error; err != nil {
		return nil, err
	}

	return r.mappers.ModelsToDomains(courseModels), nil
}

func (r *CourseRepositoryImpl) Create(ctx context.Context, newCourse domain.Course) (*domain.Course, error) {
	courseModel := r.mappers.DomainToModel(newCourse)

	if err := r.db.WithContext(ctx).Create(&courseModel).Error; err != nil {
		return nil, err
	}

	return r.mappers.ModelToDomain(courseModel), nil
}

func (r *CourseRepositoryImpl) Update(ctx context.Context, id uuid.UUID, updatedCourse domain.Course) (*domain.Course, error) {
	var existingModel models.CourseModel
	if err := r.db.WithContext(ctx).First(&existingModel, "id = ?", id).Error; err != nil {
		return nil, err
	}

	newModel := r.mappers.DomainToModel(updatedCourse)
	newModel.ID = existingModel.ID

	if err := r.db.WithContext(ctx).Save(&newModel).Error; err != nil {
		return nil, err
	}

	return r.mappers.ModelToDomain(newModel), nil
}

func (r *CourseRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&models.CourseModel{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
