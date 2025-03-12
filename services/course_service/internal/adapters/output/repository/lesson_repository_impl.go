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

type LessonRepositoryImpl struct {
	db      gorm.DB
	mappers mappers.LessonMappers
}

func NewLessonRepository(db gorm.DB) output.LessonRepository {
	return &LessonRepositoryImpl{
		db: db,
	}
}

func (r *LessonRepositoryImpl) GetById(ctx context.Context, id string) (*domain.Lesson, error) {
	var LessonModel models.LessonModel
	if err := r.db.WithContext(ctx).First(&LessonModel, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return r.mappers.ModelToDomain(LessonModel), nil
}

func (r *LessonRepositoryImpl) GetByModuleId(ctx context.Context, id string) (*[]domain.Lesson, error) {
	var LessonModels []models.LessonModel
	if err := r.db.WithContext(ctx).Where("module_id = ?", id).Find(&LessonModels).Error; err != nil {
		return nil, err
	}

	lessons := r.mappers.ModelsToDomains(LessonModels)
	return lessons, nil
}

func (r *LessonRepositoryImpl) Create(ctx context.Context, newLesson domain.Lesson) (*domain.Lesson, error) {
	LessonModel := r.mappers.DomainToModel(newLesson)

	if err := r.db.WithContext(ctx).Create(&LessonModel).Error; err != nil {
		return nil, err
	}

	return r.mappers.ModelToDomain(*LessonModel), nil
}

func (r *LessonRepositoryImpl) Update(ctx context.Context, id uuid.UUID, updatedLesson domain.Lesson) (*domain.Lesson, error) {
	var existingModel models.LessonModel
	if err := r.db.WithContext(ctx).First(&existingModel, "id = ?", id).Error; err != nil {
		return nil, err
	}

	newModel := r.mappers.DomainToModel(updatedLesson)
	newModel.ID = existingModel.ID

	if err := r.db.WithContext(ctx).Save(&newModel).Error; err != nil {
		return nil, err
	}

	return r.mappers.ModelToDomain(*newModel), nil
}

func (r *LessonRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	var courseModel models.CourseModel
	if err := r.db.WithContext(ctx).First(&courseModel, "id = ?", id).Error; err != nil {
		return err
	}

	if err := r.db.WithContext(ctx).Delete(&models.LessonModel{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
