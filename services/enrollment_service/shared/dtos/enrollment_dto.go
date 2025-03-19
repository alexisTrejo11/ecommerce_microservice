package dtos

import (
	"time"

	enrollment "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/model"
	"github.com/google/uuid"
)

type EnrollmentInsertDTO struct {
	UserID           uuid.UUID                   `json:"user_id" validate:"required"`
	CourseID         uuid.UUID                   `json:"course_id" validate:"required"`
	EnrollmentDate   string                      `json:"enrollment_date" validate:"required,datetime=2006-01-02"`
	CompletionDate   string                      `json:"completion_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
	CompletionStatus enrollment.CompletionStatus `json:"completion_status" validate:"required,oneof=in_progress completed"`
	Progress         float64                     `json:"progress" validate:"gte=0,lte=100"`
	CertificateURL   string                      `json:"certificate_url,omitempty" validate:"omitempty,url"`
	LastAccessedAt   string                      `json:"last_accessed_at,omitempty" validate:"omitempty,datetime=2006-01-02T15:04:05"`
}

type EnrollmentDTO struct {
	ID               uuid.UUID                   `json:"id"`
	UserID           uuid.UUID                   `json:"user_id"`
	CourseID         uuid.UUID                   `json:"course_id"`
	EnrollmentDate   time.Time                   `json:"enrollment_date"`
	CompletionDate   *time.Time                  `json:"completion_date,omitempty"`
	CompletionStatus enrollment.CompletionStatus `json:"completion_status"`
	Progress         float64                     `json:"progress"`
	CertificateURL   *string                     `json:"certificate_url,omitempty"`
	LastAccessedAt   *time.Time                  `json:"last_accessed_at,omitempty"`
	CompletedLessons []CompletedLessonDTO        `json:"completed_lessons,omitempty"`
}
