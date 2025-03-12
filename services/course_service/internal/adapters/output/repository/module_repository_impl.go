package repository

import (
	"context"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/models"
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
		return nil, err
	}

	module := r.mappers.ModelToDomain(moduleModel)
	r.fetchLessons(ctx, module)

	return module, nil
}

func (r *ModuleRepositoryImpl) GetByCourseId(ctx context.Context, id string) (*[]domain.Module, error) {
	var moduleModels []models.ModuleModel

	if err := r.db.WithContext(ctx).
		Where("course_id = ?", id).
		Find(&moduleModels).Error; err != nil {
		return nil, err
	}

	modules := r.mappers.ModelsToDomains(moduleModels)
	for i := range modules {
		if err := r.fetchLessons(ctx, &modules[i]); err != nil {
			return nil, err
		}
	}

	return &modules, nil
}

func (r *ModuleRepositoryImpl) Create(ctx context.Context, newModule domain.Module) (*domain.Module, error) {
	moduleModel := r.mappers.DomainToModel(newModule)

	if err := r.db.WithContext(ctx).Create(&moduleModel).Error; err != nil {
		return nil, err
	}

	module := r.mappers.ModelToDomain(*moduleModel)
	r.fetchLessons(ctx, module)

	return module, nil
}

func (r *ModuleRepositoryImpl) Update(ctx context.Context, id uuid.UUID, updatedModule domain.Module) (*domain.Module, error) {
	var existingModel models.ModuleModel
	if err := r.db.WithContext(ctx).First(&existingModel, "id = ?", id).Error; err != nil {
		return nil, err
	}

	modelUpdated := r.mappers.DomainToModel(updatedModule)
	modelUpdated.ID = existingModel.ID
	modelUpdated.CreatedAt = existingModel.CreatedAt
	modelUpdated.UpdatedAt = time.Now()

	if err := r.db.WithContext(ctx).Save(&modelUpdated).Error; err != nil {
		return nil, err
	}

	module := r.mappers.ModelToDomain(*modelUpdated)
	r.fetchLessons(ctx, module)

	return module, nil
}

func (r *ModuleRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&models.ModuleModel{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *ModuleRepositoryImpl) fetchLessons(ctx context.Context, module *domain.Module) error {
	lessons, err := r.lessonRepository.GetByModuleId(ctx, module.ID.String())
	if err != nil {
		return err
	}

	module.Lessons = *lessons

	return nil
}
