package services

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
	"github.com/google/uuid"
)

type CertificateService interface {
	GenerateCertificate(ctx context.Context, enrollmentID uuid.UUID) (*dtos.CertificateDTO, error)
	GetCertificateByEnrollment(ctx context.Context, enrollmentID uuid.UUID) (*dtos.CertificateDTO, error)
	GetCertificateByUserID(ctx context.Context, userID uuid.UUID) (*[]dtos.CertificateDTO, error)
	VerifyCertificate(ctx context.Context, certificateURL string) (*dtos.CertificateDTO, bool, error)
}
