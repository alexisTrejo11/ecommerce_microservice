package dtos

import (
	"time"

	"github.com/google/uuid"
)

// CompletedLessonDTO defines the structure for representing a completed lesson.
// swagger:model CompletedLessonDTO
type CompletedLessonDTO struct {
	// The unique identifier of the completed lesson.
	// Example: 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `json:"id"`

	// The ID of the enrollment associated with the completed lesson.
	// Example: 550e8400-e29b-41d4-a716-446655440001
	EnrollmentID uuid.UUID `json:"enrollment_id"`

	// The ID of the lesson that was completed.
	// Example: 550e8400-e29b-41d4-a716-446655440002
	LessonID uuid.UUID `json:"lesson_id"`

	// The date and time when the lesson was completed.
	// Format: date-time
	// Example: 2023-10-01T12:34:56
	CompletedAt time.Time `json:"completed_at"`
}

// CompletedLessonInsertDTO defines the structure for creating a new completed lesson.
// swagger:model CompletedLessonInsertDTO
type CompletedLessonInsertDTO struct {
	// The ID of the enrollment associated with the completed lesson.
	// Required: true
	// Example: 550e8400-e29b-41d4-a716-446655440001
	EnrollmentID uuid.UUID `json:"enrollment_id" validate:"required"`

	// The ID of the lesson that was completed.
	// Required: true
	// Example: 550e8400-e29b-41d4-a716-446655440002
	LessonID uuid.UUID `json:"lesson_id" validate:"required"`

	// The date and time when the lesson was completed.
	// Required: true
	// Format: date-time
	// Example: 2023-10-01T12:34:56
	CompletedAt string `json:"completed_at" validate:"required,datetime=2006-01-02T15:04:05"`
}
