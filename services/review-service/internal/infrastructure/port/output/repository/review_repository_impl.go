package repository

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/port/output/models"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/shared/mapper"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewRepositoryImpl struct {
	db     *gorm.DB
	mapper mapper.ReviewMapper
}

func NewReviewRepositoryImpl(db *gorm.DB) *ReviewRepositoryImpl {
	return &ReviewRepositoryImpl{db: db}
}

func (r *ReviewRepositoryImpl) Save(ctx context.Context, review *domain.Review) error {
	model := r.mapper.DomainToModel(review)
	return r.db.WithContext(ctx).Save(model).Error
}

func (r *ReviewRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*domain.Review, error) {
	var model models.ReviewModel
	err := r.db.WithContext(ctx).First(&model, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return r.mapper.ModelToDomain(&model), nil
}

func (r *ReviewRepositoryImpl) GetByCourseID(ctx context.Context, courseID uuid.UUID) ([]domain.Review, error) {
	var models []models.ReviewModel
	err := r.db.WithContext(ctx).Where("course_id = ?", courseID).Find(&models).Error
	if err != nil {
		return nil, err
	}

	domainReviews := r.mapper.ModelsToDomainList(models)
	return domainReviews, nil
}

func (r *ReviewRepositoryImpl) GetByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Review, error) {
	var models []models.ReviewModel
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&models).Error
	if err != nil {
		return nil, err
	}

	domainReviews := r.mapper.ModelsToDomainList(models)
	return domainReviews, nil
}

func (r *ReviewRepositoryImpl) DeleteByID(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&models.ReviewModel{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
