package repository

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/models"
	customErrors "github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/errors"
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
	var lessonModel models.LessonModel
	if err := r.db.WithContext(ctx).First(&lessonModel, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErrors.ErrLessonNotFoundDB
		}
		return nil, customErrors.NewDomainError("DATABASE_ERROR", "Error retrieving lesson from database", err)
	}

	return r.mappers.ModelToDomain(lessonModel), nil
}

func (r *LessonRepositoryImpl) GetByModuleId(ctx context.Context, id string) (*[]domain.Lesson, error) {
	var lessonModels []models.LessonModel
	if err := r.db.WithContext(ctx).Where("module_id = ?", id).Find(&lessonModels).Error; err != nil {
		return nil, customErrors.NewDomainError("DATABASE_ERROR", "Error retrieving lessons by module ID", err)
	}

	lessons := r.mappers.ModelsToDomains(lessonModels)
	return lessons, nil
}

func (r *LessonRepositoryImpl) Create(ctx context.Context, newLesson domain.Lesson) (*domain.Lesson, error) {
	lessonModel := r.mappers.DomainToModel(newLesson)

	if err := r.db.WithContext(ctx).Create(&lessonModel).Error; err != nil {
		return nil, customErrors.NewDomainError("DATABASE_ERROR", "Error creating lesson", err)
	}

	return r.mappers.ModelToDomain(*lessonModel), nil
}

func (r *LessonRepositoryImpl) Update(ctx context.Context, id uuid.UUID, updatedLesson domain.Lesson) (*domain.Lesson, error) {
	var existingModel models.LessonModel
	if err := r.db.WithContext(ctx).First(&existingModel, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErrors.ErrLessonNotFoundDB
		}
		return nil, customErrors.NewDomainError("DATABASE_ERROR", "Error finding lesson for update", err)
	}

	newModel := r.mappers.DomainToModel(updatedLesson)
	newModel.ID = existingModel.ID

	if err := r.db.WithContext(ctx).Save(&newModel).Error; err != nil {
		return nil, customErrors.NewDomainError("DATABASE_ERROR", "Error updating lesson", err)
	}

	return r.mappers.ModelToDomain(*newModel), nil
}

func (r *LessonRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	var lessonModel models.LessonModel
	if err := r.db.WithContext(ctx).First(&lessonModel, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErrors.ErrLessonNotFoundDB
		}
		return customErrors.NewDomainError("DATABASE_ERROR", "Error finding lesson for deletion", err)
	}

	if err := r.db.WithContext(ctx).Delete(&models.LessonModel{}, "id = ?", id).Error; err != nil {
		return customErrors.NewDomainError("DATABASE_ERROR", "Error deleting lesson", err)
	}
	return nil
}
