package services

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/models"
)

type EnrollmentService interface {
	EnrollUserInCourse(ctx context.Context, userID, courseID uint, paymentAmount float64) (*models.Enrollment, error)
	GetEnrollmentByID(ctx context.Context, id uint) (*models.Enrollment, error)
	GetEnrollmentByUserAndCourse(ctx context.Context, userID, courseID uint) (*models.Enrollment, error)
	UpdateEnrollment(ctx context.Context, enrollment *models.Enrollment) error
	CancelEnrollment(ctx context.Context, enrollmentID uint) error
	ListUserEnrollments(ctx context.Context, userID uint, page, limit int) ([]models.Enrollment, int64, error)
	ListCourseEnrollments(ctx context.Context, courseID uint, page, limit int) ([]models.Enrollment, int64, error)
	UpdateEnrollmentProgress(ctx context.Context, enrollmentID uint, progress float64) error
	MarkEnrollmentComplete(ctx context.Context, enrollmentID uint) error
	IsUserEnrolledInCourse(ctx context.Context, userID, courseID uint) (bool, error)
}
