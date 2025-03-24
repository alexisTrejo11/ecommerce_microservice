package mocks

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockProgressService struct {
	mock.Mock
}

func (m *MockProgressService) CreateCourseTrackRecord(ctx context.Context, enrollmentID uuid.UUID) error {
	args := m.Called(ctx, enrollmentID)
	return args.Error(0)
}

func (m *MockProgressService) MarkLessonComplete(ctx context.Context, enrollmentID, lessonID uuid.UUID) error {
	args := m.Called(ctx, enrollmentID, lessonID)
	return args.Error(0)
}

func (m *MockProgressService) MarkLessonIncomplete(ctx context.Context, enrollmentID, lessonID uuid.UUID) error {
	args := m.Called(ctx, enrollmentID, lessonID)
	return args.Error(0)
}

func (m *MockProgressService) GetCourseProgress(ctx context.Context, enrollmentID uuid.UUID) ([]dtos.CompletedLessonDTO, error) {
	args := m.Called(ctx, enrollmentID)
	return args.Get(0).([]dtos.CompletedLessonDTO), args.Error(1)
}

func (m *MockProgressService) CalculateProgress(ctx context.Context, enrollmentID uuid.UUID) (float64, error) {
	args := m.Called(ctx, enrollmentID)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockProgressService) IsLessonCompleted(ctx context.Context, enrollmentID, lessonID uuid.UUID) (error, bool) {
	args := m.Called(ctx, enrollmentID, lessonID)
	return args.Error(0), args.Bool(1)
}
