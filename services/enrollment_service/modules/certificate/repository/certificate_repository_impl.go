package repository

import (
	"context"
	"errors"
	"log"

	certificate "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/model"
	appErr "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/error"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CertificateRepositoryImpl struct {
	db *gorm.DB
}

func NewCertificateRepository(db *gorm.DB) CertificateRepository {
	return &CertificateRepositoryImpl{db: db}
}

func (r *CertificateRepositoryImpl) Create(ctx context.Context, certificate *certificate.Certificate) error {
	if err := r.db.WithContext(ctx).Create(&certificate).Error; err != nil {
		log.Printf("Error creating certificate: %v", err)
		return appErr.ErrDB
	}
	return nil
}

func (r *CertificateRepositoryImpl) GetByEnrollment(ctx context.Context, enrollmentID uuid.UUID) (*certificate.Certificate, error) {
	var certificate certificate.Certificate
	if err := r.db.WithContext(ctx).Where("enrollment_id = ?", enrollmentID).First(&certificate).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErr.ErrCertificateNotFoundDB
		}
		log.Printf("Error fetching certificate by enrollment ID %s: %v", enrollmentID, err)
		return nil, appErr.ErrDB
	}
	return &certificate, nil
}

func (r *CertificateRepositoryImpl) GetByUser(ctx context.Context, userID uuid.UUID) (*[]certificate.Certificate, error) {
	var certificates []certificate.Certificate
	if err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&certificates).Error; err != nil {
		return nil, appErr.ErrDB
	}
	return &certificates, nil
}

func (r *CertificateRepositoryImpl) Update(ctx context.Context, certificate *certificate.Certificate) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Save(certificate).Error; err != nil {
		tx.Rollback()
		return appErr.ErrDB
	}

	return tx.Commit().Error
}
