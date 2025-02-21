package repository

import (
	"context"
	"errors"

	models "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MFARepositoryImpl struct {
	db *gorm.DB
}

func NewMFARepository(db *gorm.DB) output.MFARepository {
	return &MFARepositoryImpl{
		db: db,
	}
}

func (r *MFARepositoryImpl) Create(ctx context.Context, mfa *entities.MFA) error {
	return nil
}

func (r *MFARepositoryImpl) FindByUserID(ctx context.Context, userID uuid.UUID) (*entities.MFA, error) {
	var mfaModel models.MFAModel
	if err := r.db.First(&mfaModel, "user_id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return nil, nil
}

func (r *MFARepositoryImpl) Update(ctx context.Context, mfa *entities.MFA) error {
	return nil
}

func (r *MFARepositoryImpl) Delete(ctx context.Context, userID uuid.UUID) error {
	if err := r.db.Where("user_id = ?", userID).Delete(&models.MFAModel{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}
		return err
	}

	return nil
}
