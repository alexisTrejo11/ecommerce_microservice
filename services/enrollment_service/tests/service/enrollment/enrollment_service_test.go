package mocks

import (
	"context"
	"testing"
	"time"

	enrollment "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/model"
	services "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/service"
	appErr "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/error"
	mocks "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/tests/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetEnrollmentByID_Success(t *testing.T) {
	mockRepo := new(mocks.MockEnrollmentRepository)
	service := services.NewEnrollmentService(mockRepo)

	id := uuid.New()
	enrollment := &enrollment.Enrollment{
		ID:             id,
		UserID:         uuid.New(),
		CourseID:       uuid.New(),
		EnrollmentDate: time.Now(),
	}

	mockRepo.On("GetByID", mock.Anything, id).Return(enrollment, nil)

	result, err := service.GetEnrollmentByID(context.Background(), id)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, enrollment.ID, result.ID)
	mockRepo.AssertExpectations(t)
}

func TestGetEnrollmentByID_NotFound(t *testing.T) {
	mockRepo := new(mocks.MockEnrollmentRepository)
	service := services.NewEnrollmentService(mockRepo)

	id := uuid.New()

	mockRepo.On("GetByID", mock.Anything, id).Return(nil, appErr.ErrNotFoundDB)

	result, err := service.GetEnrollmentByID(context.Background(), id)

	assert.Error(t, err)
	assert.Equal(t, appErr.ErrNotFoundDB, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestEnrollUserInCourse_Success(t *testing.T) {
	mockRepo := new(mocks.MockEnrollmentRepository)
	service := services.NewEnrollmentService(mockRepo)

	userID := uuid.New()
	courseID := uuid.New()

	expectedEnrollment := enrollment.NewEnrollment(userID, courseID)

	mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(enrollment *enrollment.Enrollment) bool {
		return enrollment.UserID == expectedEnrollment.UserID &&
			enrollment.CourseID == expectedEnrollment.CourseID &&
			enrollment.CompletionStatus == expectedEnrollment.CompletionStatus &&
			enrollment.Progress == expectedEnrollment.Progress
	})).Return(nil)

	result, err := service.EnrollUserInCourse(context.Background(), userID, courseID)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, userID, result.UserID)
	assert.Equal(t, courseID, result.CourseID)

	mockRepo.AssertExpectations(t)
}

func TestEnrollUserInCourse_Failure(t *testing.T) {
	mockRepo := new(mocks.MockEnrollmentRepository)
	service := services.NewEnrollmentService(mockRepo)

	userID := uuid.New()
	courseID := uuid.New()

	expectedEnrollment := enrollment.NewEnrollment(userID, courseID)

	mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(enrollment *enrollment.Enrollment) bool {
		return enrollment.UserID == expectedEnrollment.UserID &&
			enrollment.CourseID == expectedEnrollment.CourseID &&
			enrollment.CompletionStatus == expectedEnrollment.CompletionStatus &&
			enrollment.Progress == expectedEnrollment.Progress
	})).Return(appErr.ErrDB)

	result, err := service.EnrollUserInCourse(context.Background(), userID, courseID)

	assert.Error(t, err)
	assert.Equal(t, appErr.ErrDB, err)
	assert.Nil(t, result)

	mockRepo.AssertExpectations(t)
}

func TestCancelEnrollment_Success(t *testing.T) {
	mockRepo := new(mocks.MockEnrollmentRepository)
	service := services.NewEnrollmentService(mockRepo)

	enrollmentID := uuid.New()
	existingEnrollment := &enrollment.Enrollment{
		ID:               enrollmentID,
		CompletionStatus: enrollment.STARTING,
	}

	mockRepo.On("GetByID", mock.Anything, enrollmentID).Return(existingEnrollment, nil)
	mockRepo.On("Update", mock.Anything, existingEnrollment).Return(nil)

	err := service.CancelEnrollment(context.Background(), uuid.New(), enrollmentID)

	assert.NoError(t, err)         // Ensure no error is returned
	mockRepo.AssertExpectations(t) // Ensure all expectations were met
}

func TestCancelEnrollment_NotFound(t *testing.T) {
	mockRepo := new(mocks.MockEnrollmentRepository)
	service := services.NewEnrollmentService(mockRepo)

	enrollmentID := uuid.New()

	mockRepo.On("GetByID", mock.Anything, enrollmentID).Return(nil, appErr.ErrNotFoundDB)

	err := service.CancelEnrollment(context.Background(), uuid.New(), enrollmentID)

	assert.Error(t, err)
	assert.Equal(t, appErr.ErrNotFoundDB, err)
	mockRepo.AssertExpectations(t)
}
