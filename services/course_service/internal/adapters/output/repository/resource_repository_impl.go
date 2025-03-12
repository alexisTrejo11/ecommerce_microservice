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
	var ResourceModel models.ResourceModel
	if err := r.db.WithContext(ctx).First(&ResourceModel, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return r.mappers.ModelToDomain(ResourceModel), nil
}

func (r *ResourceRepositoryImpl) GetByLessonId(ctx context.Context, id string) (*[]domain.Resource, error) {
	var ResourceModels []models.ResourceModel
	if err := r.db.WithContext(ctx).Where("id = ?", id).Find(&ResourceModels).Error; err != nil {
		return nil, err
	}

	Resources := make([]domain.Resource, len(ResourceModels))
	for i, model := range ResourceModels {
		Resources[i] = *r.mappers.ModelToDomain(model)
	}
	return &Resources, nil
}

func (r *ResourceRepositoryImpl) Create(ctx context.Context, newResource domain.Resource) (*domain.Resource, error) {
	ResourceModel := r.mappers.DomainToModel(newResource)

	if err := r.db.WithContext(ctx).Create(&ResourceModel).Error; err != nil {
		return nil, err
	}

	return r.mappers.ModelToDomain(*ResourceModel), nil
}

func (r *ResourceRepositoryImpl) Update(ctx context.Context, id uuid.UUID, updatedResource domain.Resource) (*domain.Resource, error) {
	var existingModel models.ResourceModel
	if err := r.db.WithContext(ctx).First(&existingModel, "id = ?", id).Error; err != nil {
		return nil, err
	}

	newModel := r.mappers.DomainToModel(updatedResource)
	newModel.ID = existingModel.ID

	if err := r.db.WithContext(ctx).Save(&newModel).Error; err != nil {
		return nil, err
	}

	return r.mappers.ModelToDomain(*newModel), nil
}

func (r *ResourceRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	var courseModel models.CourseModel
	if err := r.db.WithContext(ctx).First(&courseModel, "id = ?", id).Error; err != nil {
		return err
	}

	if err := r.db.WithContext(ctx).Delete(&models.ResourceModel{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
