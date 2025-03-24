package certificate

import (
	"context"
	"testing"
	"time"

	certificate "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/model"
	services "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/service"
	enrollment "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/model"
	mocks "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/tests/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGenerateCertificate_Success(t *testing.T) {
	mockCertRepo := new(mocks.MockCertificateRepository)
	mockEnrollRepo := new(mocks.MockEnrollmentRepository)
	service := services.NewCertificateService(mockCertRepo, mockEnrollRepo)

	enrollmentID := uuid.New()
	existingEnrollment := &enrollment.Enrollment{
		ID: enrollmentID,
	}

	// Mock EnrollmentRepository.GetByID
	mockEnrollRepo.On("GetByID", mock.Anything, enrollmentID).Return(existingEnrollment, nil)

	// Mock CertificateRepository.Create
	mockCertRepo.On("Create", mock.Anything, mock.MatchedBy(func(cert *certificate.Certificate) bool {
		return cert.EnrollmentID == enrollmentID && cert.CertificateURL == "www.placeholder.com/certificate"
	})).Return(nil)

	result, err := service.GenerateCertificate(context.Background(), enrollmentID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, enrollmentID, result.EnrollmentID)
	mockCertRepo.AssertExpectations(t)
	mockEnrollRepo.AssertExpectations(t)
}

func TestGetCertificateByEnrollment_Success(t *testing.T) {
	mockCertRepo := new(mocks.MockCertificateRepository)
	mockEnrollRepo := new(mocks.MockEnrollmentRepository)
	service := services.NewCertificateService(mockCertRepo, mockEnrollRepo)

	enrollmentID := uuid.New()
	certificate := &certificate.Certificate{
		ID:             uuid.New(),
		EnrollmentID:   enrollmentID,
		IssuedAt:       time.Now(),
		CertificateURL: "www.placeholder.com/certificate",
	}

	// Mock CertificateRepository.GetByEnrollment
	mockCertRepo.On("GetByEnrollment", mock.Anything, enrollmentID).Return(certificate, nil)

	result, err := service.GetCertificateByEnrollment(context.Background(), enrollmentID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, enrollmentID, result.EnrollmentID)
	mockCertRepo.AssertExpectations(t)
}

func TestGetCertificateByUserID_Success(t *testing.T) {
	mockCertRepo := new(mocks.MockCertificateRepository)
	mockEnrollRepo := new(mocks.MockEnrollmentRepository)
	service := services.NewCertificateService(mockCertRepo, mockEnrollRepo)

	userID := uuid.New()
	certificates := []certificate.Certificate{
		{
			ID:             uuid.New(),
			EnrollmentID:   uuid.New(),
			IssuedAt:       time.Now(),
			CertificateURL: "www.placeholder.com/certificate1",
		},
		{
			ID:             uuid.New(),
			EnrollmentID:   uuid.New(),
			IssuedAt:       time.Now(),
			CertificateURL: "www.placeholder.com/certificate2",
		},
	}

	// Mock CertificateRepository.GetByUser
	mockCertRepo.On("GetByUser", mock.Anything, userID).Return(&certificates, nil)

	result, err := service.GetCertificateByUserID(context.Background(), userID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, *result, 2)
	mockCertRepo.AssertExpectations(t)
}

func TestVerifyCertificate_Success(t *testing.T) {
	mockCertRepo := new(mocks.MockCertificateRepository)
	mockEnrollRepo := new(mocks.MockEnrollmentRepository)
	service := services.NewCertificateService(mockCertRepo, mockEnrollRepo)

	certificateURL := "www.placeholder.com/certificate"

	result, isValid, err := service.VerifyCertificate(context.Background(), certificateURL)

	assert.NoError(t, err)
	assert.True(t, isValid)
	assert.Nil(t, result)
}
