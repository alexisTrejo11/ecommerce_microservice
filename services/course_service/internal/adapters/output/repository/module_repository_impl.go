package repository

import (
	"context"
	"errors"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/models"
	customErrors "github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/errors"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModuleRepositoryImpl struct {
	db               gorm.DB
	lessonRepository output.LessonRepository
	mappers          mappers.ModuleMapper
}

func NewModuleRepository(db gorm.DB, lessonRepository output.LessonRepository) output.ModuleRepository {
	return &ModuleRepositoryImpl{
		db:               db,
		lessonRepository: lessonRepository,
	}
}

func (r *ModuleRepositoryImpl) GetById(ctx context.Context, id string) (*domain.Module, error) {
	var moduleModel models.ModuleModel
	if err := r.db.WithContext(ctx).First(&moduleModel, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErrors.ErrModuleNotFound
		}
		return nil, customErrors.NewDomainError("DATABASE_ERROR", "Error retrieving module from database", err)
	}

	module, err := r.mappers.ModelToDomain(moduleModel)
	if err != nil {
		return nil, customErrors.NewDomainError("MAPPING_ERROR", "Error mapping module model to domain", err)
	}

	if err := r.fetchLessons(ctx, module); err != nil {
		return nil, customErrors.ErrLessonFetchErrorDB
	}

	return module, nil
}

func (r *ModuleRepositoryImpl) GetByCourseId(ctx context.Context, id string) (*[]domain.Module, error) {
	var moduleModels []models.ModuleModel
	if err := r.db.WithContext(ctx).
		Where("course_id = ?", id).
		Find(&moduleModels).Error; err != nil {
		return nil, customErrors.NewDomainError("DATABASE_ERROR", "Error retrieving modules by course ID", err)
	}

	modules := r.mappers.ModelsToDomains(moduleModels)
	for i := range modules {
		if err := r.fetchLessons(ctx, &modules[i]); err != nil {
			return nil, customErrors.ErrLessonFetchErrorDB
		}
	}

	return &modules, nil
}

func (r *ModuleRepositoryImpl) Create(ctx context.Context, newModule domain.Module) (*domain.Module, error) {
	moduleModel := r.mappers.DomainToModel(newModule)

	if err := r.db.WithContext(ctx).Create(&moduleModel).Error; err != nil {
		return nil, customErrors.NewDomainError("DATABASE_ERROR", "Error creating module", err)
	}

	module, err := r.mappers.ModelToDomain(*moduleModel)
	if err != nil {
		return nil, customErrors.NewDomainError("MAPPING_ERROR", "Error mapping module model to domain", err)
	}

	if err := r.fetchLessons(ctx, module); err != nil {
		return nil, customErrors.ErrLessonFetchErrorDB
	}

	return module, nil
}

func (r *ModuleRepositoryImpl) Update(ctx context.Context, id uuid.UUID, updatedModule domain.Module) (*domain.Module, error) {
	var existingModel models.ModuleModel
	if err := r.db.WithContext(ctx).First(&existingModel, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErrors.ErrModuleNotFound
		}
		return nil, customErrors.NewDomainError("DATABASE_ERROR", "Error finding module for update", err)
	}

	modelUpdated := r.mappers.DomainToModel(updatedModule)
	modelUpdated.ID = existingModel.ID
	modelUpdated.CreatedAt = existingModel.CreatedAt
	modelUpdated.UpdatedAt = time.Now()

	if err := r.db.WithContext(ctx).Save(&modelUpdated).Error; err != nil {
		return nil, customErrors.NewDomainError("DATABASE_ERROR", "Error updating module", err)
	}

	module, err := r.mappers.ModelToDomain(*modelUpdated)
	if err != nil {
		return nil, customErrors.NewDomainError("MAPPING_ERROR", "Error mapping module model to domain", err)
	}

	if err := r.fetchLessons(ctx, module); err != nil {
		return nil, customErrors.ErrLessonFetchErrorDB
	}

	return module, nil
}

func (r *ModuleRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	var existingModel models.ModuleModel
	if err := r.db.WithContext(ctx).First(&existingModel, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErrors.ErrModuleNotFoundDB
		}
		return customErrors.NewDomainError("DATABASE_ERROR", "Error finding module for delete", err)
	}

	if err := r.db.WithContext(ctx).Delete(&models.ModuleModel{}, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErrors.ErrModuleNotFound
		}
		return customErrors.NewDomainError("DATABASE_ERROR", "Error deleting module", err)
	}
	return nil
}

func (r *ModuleRepositoryImpl) fetchLessons(ctx context.Context, module *domain.Module) error {
	lessons, err := r.lessonRepository.GetByModuleId(ctx, module.ID().String())
	if err != nil {
		return customErrors.ErrLessonFetchErrorDB
	}

	err = module.SetLessons(*lessons)
	if err != nil {
		return customErrors.NewDomainError("INVALID_OPERATION", "Error setting lessons for module", err)
	}

	return nil
}
