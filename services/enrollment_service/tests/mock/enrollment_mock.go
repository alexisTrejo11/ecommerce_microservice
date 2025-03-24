package mocks

import (
	"context"

	enrollment "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockEnrollmentRepository struct {
	mock.Mock
}

func (m *MockEnrollmentRepository) GetByID(ctx context.Context, id uuid.UUID) (*enrollment.Enrollment, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*enrollment.Enrollment), args.Error(1)
}

func (m *MockEnrollmentRepository) GetByUserAndCourse(ctx context.Context, userID, courseID uuid.UUID) (*enrollment.Enrollment, error) {
	args := m.Called(ctx, userID, courseID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*enrollment.Enrollment), args.Error(1)
}

func (m *MockEnrollmentRepository) GetByUser(ctx context.Context, userID uuid.UUID, page, limit int) ([]enrollment.Enrollment, int64, error) {
	args := m.Called(ctx, userID, page, limit)
	if args.Get(0) == nil {
		return nil, 0, args.Error(2)
	}
	return args.Get(0).([]enrollment.Enrollment), args.Get(1).(int64), args.Error(2)
}

func (m *MockEnrollmentRepository) GetByCourse(ctx context.Context, courseID uuid.UUID, page, limit int) ([]enrollment.Enrollment, int64, error) {
	args := m.Called(ctx, courseID, page, limit)
	if args.Get(0) == nil {
		return nil, 0, args.Error(2)
	}
	return args.Get(0).([]enrollment.Enrollment), args.Get(1).(int64), args.Error(2)
}

func (m *MockEnrollmentRepository) Create(ctx context.Context, enrollment *enrollment.Enrollment) error {
	args := m.Called(ctx, enrollment)
	return args.Error(0)
}

func (m *MockEnrollmentRepository) Update(ctx context.Context, enrollment *enrollment.Enrollment) error {
	args := m.Called(ctx, enrollment)
	return args.Error(0)
}

func (m *MockEnrollmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
