package mocks

import (
	"context"

	certificate "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockCertificateRepository struct {
	mock.Mock
}

func (m *MockCertificateRepository) Create(ctx context.Context, certificate *certificate.Certificate) error {
	args := m.Called(ctx, certificate)
	return args.Error(0)
}

func (m *MockCertificateRepository) GetByEnrollment(ctx context.Context, enrollmentID uuid.UUID) (*certificate.Certificate, error) {
	args := m.Called(ctx, enrollmentID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*certificate.Certificate), args.Error(1)
}

func (m *MockCertificateRepository) GetByUser(ctx context.Context, userID uuid.UUID) (*[]certificate.Certificate, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(*[]certificate.Certificate), args.Error(1)
}

func (m *MockCertificateRepository) Update(ctx context.Context, cert *certificate.Certificate) error {
	args := m.Called(ctx, cert)
	return args.Error(0)
}
