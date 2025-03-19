package services

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
	"github.com/google/uuid"
)

type EnrollmentService interface {
	GetEnrollmentByID(ctx context.Context, id uuid.UUID) (*dtos.EnrollmentDTO, error)
	GetEnrollmentByUserAndCourse(ctx context.Context, userID, courseID uuid.UUID) (*dtos.EnrollmentDTO, error)
	GetUserEnrollments(ctx context.Context, userID uuid.UUID, page, limit int) ([]dtos.EnrollmentDTO, int64, error)
	GetCourseEnrollments(ctx context.Context, courseID uuid.UUID, page, limit int) ([]dtos.EnrollmentDTO, int64, error)

	EnrollUserInCourse(ctx context.Context, userID, courseID uuid.UUID) (*dtos.EnrollmentDTO, error)
	CancelEnrollment(ctx context.Context, enrollmentID uuid.UUID) error
	MarkEnrollmentComplete(ctx context.Context, enrollmentID uuid.UUID) error
	IsUserEnrolledInCourse(ctx context.Context, userID, courseID uuid.UUID) bool
}
