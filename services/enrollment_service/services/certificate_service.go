package services

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/models"
)

type CertificateService interface {
	GenerateCertificate(ctx context.Context, enrollmentID uint) (*models.Certificate, error)
	GetCertificateByEnrollment(ctx context.Context, enrollmentID uint) (*models.Certificate, error)
	VerifyCertificate(ctx context.Context, certificateURL string) (*models.Certificate, bool, error)
}
