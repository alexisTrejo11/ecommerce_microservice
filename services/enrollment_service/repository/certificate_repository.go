package repository

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/models"
)

type CertificateRepository interface {
	Create(ctx context.Context, certificate *models.Certificate) error
	GetByEnrollment(ctx context.Context, enrollmentID uint) (*models.Certificate, error)
	Update(ctx context.Context, certificate *models.Certificate) error
}
