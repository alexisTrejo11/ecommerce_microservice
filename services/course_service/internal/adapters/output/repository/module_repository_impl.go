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

type ModuleRepositoryImpl struct {
	db      gorm.DB
	mappers mappers.ModuleMapper
}

func NewModuleRepository(db gorm.DB) output.ModuleRepository {
	return &ModuleRepositoryImpl{
		db: db,
	}
}

func (r *ModuleRepositoryImpl) GetById(ctx context.Context, id string) (*domain.Module, error) {
	var moduleModel models.ModuleModel
	if err := r.db.WithContext(ctx).First(&moduleModel, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return r.mappers.ModelToDomain(moduleModel), nil
}

func (r *ModuleRepositoryImpl) GetByCourseId(ctx context.Context, id string) (*[]domain.Module, error) {
	var moduleModels []models.ModuleModel
	if err := r.db.WithContext(ctx).Where(&moduleModels, "id = ?", id).Error; err != nil {
		return nil, err
	}

	modules := make([]domain.Module, len(moduleModels))
	for i, model := range moduleModels {
		modules[i] = *r.mappers.ModelToDomain(model)
	}
	return &modules, nil
}

func (r *ModuleRepositoryImpl) Create(ctx context.Context, newModule domain.Module) (*domain.Module, error) {
	moduleModel := r.mappers.DomainToModel(newModule)

	if err := r.db.WithContext(ctx).Create(&moduleModel).Error; err != nil {
		return nil, err
	}

	return r.mappers.ModelToDomain(*moduleModel), nil
}

func (r *ModuleRepositoryImpl) Update(ctx context.Context, id uuid.UUID, updatedModule domain.Module) (*domain.Module, error) {
	var existingModel models.ModuleModel
	if err := r.db.WithContext(ctx).First(&existingModel, "id = ?", id).Error; err != nil {
		return nil, err
	}

	newModel := r.mappers.DomainToModel(updatedModule)
	newModel.ID = existingModel.ID

	if err := r.db.WithContext(ctx).Save(&newModel).Error; err != nil {
		return nil, err
	}

	return r.mappers.ModelToDomain(*newModel), nil
}

func (r *ModuleRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&models.ModuleModel{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
