package repository

import (
	"context"

	certificate "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/model"
	"github.com/google/uuid"
)

type CertificateRepository interface {
	Create(ctx context.Context, certificate *certificate.Certificate) error
	GetByEnrollment(ctx context.Context, enrollmentID uuid.UUID) (*certificate.Certificate, error)
	GetByUser(ctx context.Context, userID uuid.UUID) (*[]certificate.Certificate, error)
	Update(ctx context.Context, certificate *certificate.Certificate) error
}
