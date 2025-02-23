package repository

import (
	"context"
	"errors"
	"fmt"

	models "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MFARepositoryImpl struct {
	db         *gorm.DB
	mfaMappers mappers.MfaMappers
}

func NewMFARepository(db *gorm.DB) output.MFARepository {
	return &MFARepositoryImpl{
		db: db,
	}
}

func (r *MFARepositoryImpl) findMFAByUserID(ctx context.Context, userID uuid.UUID) (*models.MFAModel, error) {
	var mfaModel models.MFAModel
	if err := r.db.First(&mfaModel, "user_id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("MFA not found for userID %s: %w", userID, ErrNotFound)
		}
		return nil, fmt.Errorf("error finding MFA for userID %s: %w", userID, err)
	}
	return &mfaModel, nil
}

func (r *MFARepositoryImpl) Create(ctx context.Context, mfa *entities.MFA) error {
	mfaModel := r.mfaMappers.DomainToModel(*mfa)

	if err := r.db.WithContext(ctx).Create(&mfaModel).Error; err != nil {
		return fmt.Errorf("error creating MFA: %w", err)
	}

	return nil
}

func (r *MFARepositoryImpl) FindByUserID(ctx context.Context, userID uuid.UUID) (*entities.MFA, error) {
	mfaModel, err := r.findMFAByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	mfa := r.mfaMappers.ModelToDomain(*mfaModel)
	return mfa, nil
}

func (r *MFARepositoryImpl) Update(ctx context.Context, mfa *entities.MFA) error {
	mfaModel := r.mfaMappers.DomainToModel(*mfa)

	if err := r.db.WithContext(ctx).Save(&mfaModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("MFA not found: %w", ErrNotFound)
		}
		return fmt.Errorf("error updating MFA: %w", err)
	}

	return nil
}

func (r *MFARepositoryImpl) Delete(ctx context.Context, userID uuid.UUID) error {
	_, err := r.findMFAByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if err := r.db.Where("user_id = ?", userID).Delete(&models.MFAModel{}).Error; err != nil {
		return fmt.Errorf("error deleting MFA for userID %s: %w", userID, err)
	}

	return nil
}
