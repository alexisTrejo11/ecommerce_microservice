package repository

import (
	"context"

	certificate "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/model"
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
		return err
	}

	return nil
}

func (r *CertificateRepositoryImpl) GetByEnrollment(ctx context.Context, enrollmentID uuid.UUID) (*certificate.Certificate, error) {
	var certificate certificate.Certificate
	if err := r.db.WithContext(ctx).Where("enrollment_id = ?", enrollmentID).First(&certificate).Error; err != nil {
		return nil, err
	}

	return &certificate, nil
}

func (r *CertificateRepositoryImpl) GetByUser(ctx context.Context, userID uuid.UUID) (*[]certificate.Certificate, error) {
	var certificates []certificate.Certificate
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&certificates).Error; err != nil {
		return nil, err
	}

	return &certificates, nil
}

func (r *CertificateRepositoryImpl) Update(ctx context.Context, certificate *certificate.Certificate) error {
	if err := r.db.WithContext(ctx).Save(certificate).Error; err != nil {
		return err
	}
	return nil
}
