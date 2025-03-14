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

type ResourceRepositoryImpl struct {
	db      gorm.DB
	mappers mappers.ResourceMapper
}

func NewResourceRepository(db gorm.DB) output.ResourceRepository {
	return &ResourceRepositoryImpl{
		db: db,
	}
}

func (r *ResourceRepositoryImpl) GetById(ctx context.Context, id string) (*domain.Resource, error) {
	var resourceModel models.ResourceModel
	if err := r.db.WithContext(ctx).First(&resourceModel, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErrors.ErrResourceNotFoundDB
		}
		return nil, customErrors.NewDomainError("DATABASE_ERROR", "Error retrieving resource from database", err)
	}

	return r.mappers.ModelToDomain(resourceModel), nil
}

func (r *ResourceRepositoryImpl) GetByLessonId(ctx context.Context, id string) (*[]domain.Resource, error) {
	var resourceModels []models.ResourceModel
	if err := r.db.WithContext(ctx).Where("lesson_id = ?", id).Find(&resourceModels).Error; err != nil {
		return nil, customErrors.NewDomainError("DATABASE_ERROR", "Error retrieving resources by lesson ID", err)
	}

	resources := make([]domain.Resource, len(resourceModels))
	for i, model := range resourceModels {
		resources[i] = *r.mappers.ModelToDomain(model)
	}
	return &resources, nil
}

func (r *ResourceRepositoryImpl) Create(ctx context.Context, newResource domain.Resource) (*domain.Resource, error) {
	resourceModel := r.mappers.DomainToModel(newResource)

	if err := r.db.WithContext(ctx).Create(&resourceModel).Error; err != nil {
		return nil, customErrors.NewDomainError("DATABASE_ERROR", "Error creating resource", err)
	}

	return r.mappers.ModelToDomain(*resourceModel), nil
}

func (r *ResourceRepositoryImpl) Update(ctx context.Context, id uuid.UUID, updatedResource domain.Resource) (*domain.Resource, error) {
	var existingModel models.ResourceModel
	if err := r.db.WithContext(ctx).First(&existingModel, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErrors.ErrResourceNotFoundDB
		}
		return nil, customErrors.NewDomainError("DATABASE_ERROR", "Error finding resource for update", err)
	}

	newModel := r.mappers.DomainToModel(updatedResource)
	newModel.ID = existingModel.ID

	if err := r.db.WithContext(ctx).Save(&newModel).Error; err != nil {
		return nil, customErrors.NewDomainError("DATABASE_ERROR", "Error updating resource", err)
	}

	return r.mappers.ModelToDomain(*newModel), nil
}

func (r *ResourceRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	var resourceModel models.ResourceModel
	if err := r.db.WithContext(ctx).First(&resourceModel, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErrors.ErrResourceNotFoundDB
		}
		return customErrors.NewDomainError("DATABASE_ERROR", "Error finding resource for deletion", err)
	}

	if err := r.db.WithContext(ctx).Delete(&models.ResourceModel{}, "id = ?", id).Error; err != nil {
		return customErrors.NewDomainError("DATABASE_ERROR", "Error deleting resource", err)
	}
	return nil
}
