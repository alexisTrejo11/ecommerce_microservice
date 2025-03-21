package dtos

import (
	"time"

	enrollment "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/model"
	"github.com/google/uuid"
)

// EnrollmentInsertDTO defines the structure for creating a new enrollment.
// swagger:model EnrollmentInsertDTO
type EnrollmentInsertDTO struct {
	// The ID of the user enrolling in the course.
	// Required: true
	// Example: 550e8400-e29b-41d4-a716-446655440000
	UserID uuid.UUID `json:"user_id" validate:"required"`

	// The ID of the course being enrolled in.
	// Required: true
	// Example: 550e8400-e29b-41d4-a716-446655440001
	CourseID uuid.UUID `json:"course_id" validate:"required"`

	// The date when the user enrolled in the course.
	// Required: true
	// Format: date
	// Example: 2023-10-01
	EnrollmentDate string `json:"enrollment_date" validate:"required,datetime=2006-01-02"`

	// The date when the user completed the course (optional).
	// Format: date
	// Example: 2023-12-01
	CompletionDate string `json:"completion_date,omitempty" validate:"omitempty,datetime=2006-01-02"`

	// The completion status of the enrollment.
	// Required: true
	// Enum: in_progress, completed
	// Example: in_progress
	CompletionStatus enrollment.CompletionStatus `json:"completion_status" validate:"required,oneof=in_progress completed"`

	// The progress of the user in the course, represented as a percentage.
	// Required: true
	// Minimum: 0
	// Maximum: 100
	// Example: 75.5
	Progress float64 `json:"progress" validate:"gte=0,lte=100"`

	// The URL where the certificate can be accessed (optional).
	// Format: uri
	// Example: https://example.com/certificates/12345
	CertificateURL string `json:"certificate_url,omitempty" validate:"omitempty,url"`

	// The last time the user accessed the course (optional).
	// Format: date-time
	// Example: 2023-11-15T14:23:45
	LastAccessedAt string `json:"last_accessed_at,omitempty" validate:"omitempty,datetime=2006-01-02T15:04:05"`
}

// EnrollmentDTO defines the structure for representing an enrollment.
// swagger:model EnrollmentDTO
type EnrollmentDTO struct {
	// The unique identifier of the enrollment.
	// Example: 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `json:"id"`

	// The ID of the user enrolled in the course.
	// Example: 550e8400-e29b-41d4-a716-446655440001
	UserID uuid.UUID `json:"user_id"`

	// The ID of the course being enrolled in.
	// Example: 550e8400-e29b-41d4-a716-446655440002
	CourseID uuid.UUID `json:"course_id"`

	// The date when the user enrolled in the course.
	// Format: date
	// Example: 2023-10-01
	EnrollmentDate time.Time `json:"enrollment_date"`

	// The date when the user completed the course (optional).
	// Format: date
	// Example: 2023-12-01
	CompletionDate *time.Time `json:"completion_date,omitempty"`

	// The completion status of the enrollment.
	// Enum: in_progress, completed
	// Example: in_progress
	CompletionStatus enrollment.CompletionStatus `json:"completion_status"`

	// The progress of the user in the course, represented as a percentage.
	// Minimum: 0
	// Maximum: 100
	// Example: 75.5
	Progress float64 `json:"progress"`

	// The URL where the certificate can be accessed (optional).
	// Format: uri
	// Example: https://example.com/certificates/12345
	CertificateURL *string `json:"certificate_url,omitempty"`

	// The last time the user accessed the course (optional).
	// Format: date-time
	// Example: 2023-11-15T14:23:45
	LastAccessedAt *time.Time `json:"last_accessed_at,omitempty"`

	// The list of lessons completed by the user (optional).
	CompletedLessons []CompletedLessonDTO `json:"completed_lessons,omitempty"`
}
