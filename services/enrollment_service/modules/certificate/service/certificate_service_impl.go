package services

import (
	"context"
	"time"

	certificate "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/model"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/repository"
	enrollmentRepo "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/repository"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
	mapper "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/mappers"
	"github.com/google/uuid"
)

type CertificateServiceImpl struct {
	certificateRepository repository.CertificateRepository
	enrollmentRepository  enrollmentRepo.EnrollmentRepository
}

func NewCertificateService(certificateRepository repository.CertificateRepository, enrollmentRepository enrollmentRepo.EnrollmentRepository) CertificateService {
	return &CertificateServiceImpl{
		certificateRepository: certificateRepository,
		enrollmentRepository:  enrollmentRepository,
	}
}

func (s *CertificateServiceImpl) GenerateCertificate(ctx context.Context, enrollmentID uuid.UUID) (*dtos.CertificateDTO, error) {
	if err := s.validateCertificationCreation(ctx, enrollmentID); err != nil {
		return nil, err
	}

	certificate := s.generateCertificate(enrollmentID)
	s.certificateRepository.Create(ctx, certificate)

	certificateDTO := mapper.ToCertificateDTO(*certificate)
	return &certificateDTO, nil
}

func (s *CertificateServiceImpl) GetCertificateByEnrollment(ctx context.Context, enrollmentID uuid.UUID) (*dtos.CertificateDTO, error) {
	certificate, err := s.certificateRepository.GetByEnrollment(ctx, enrollmentID)
	if err != nil {
		return nil, err
	}

	certificateDTO := mapper.ToCertificateDTO(*certificate)
	return &certificateDTO, nil
}

func (s *CertificateServiceImpl) GetCertificateByUserID(ctx context.Context, userID uuid.UUID) (*[]dtos.CertificateDTO, error) {
	certificates, err := s.certificateRepository.GetByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	certificateDTOs := mapper.ToCertificateDTOs(*certificates)
	return &certificateDTOs, nil
}

func (s *CertificateServiceImpl) VerifyCertificate(ctx context.Context, certificateURL string) (*dtos.CertificateDTO, bool, error) {
	return nil, true, nil
}

func (s *CertificateServiceImpl) generateCertificate(enrollmentID uuid.UUID) *certificate.Certificate {
	// Implement URL Certificate
	return &certificate.Certificate{
		ID:             uuid.New(),
		EnrollmentID:   enrollmentID,
		IssuedAt:       time.Now(),
		CertificateURL: "www.placeholder.com/certificate",
	}
}

func (s *CertificateServiceImpl) validateCertificationCreation(ctx context.Context, enrollmentID uuid.UUID) error {
	_, err := s.enrollmentRepository.GetByID(ctx, enrollmentID)
	if err != nil {
		return err
	}

	// Validate All Lesson completed
	return nil
}
