package dtos

import (
	"time"

	"github.com/google/uuid"
)

type CompletedLessonDTO struct {
	ID           uuid.UUID `json:"id"`
	EnrollmentID uuid.UUID `json:"enrollment_id"`
	LessonID     uuid.UUID `json:"lesson_id"`
	CompletedAt  time.Time `json:"completed_at"`
}

type CompletedLessonInsertDTO struct {
	EnrollmentID uuid.UUID `json:"enrollment_id" validate:"required"`
	LessonID     uuid.UUID `json:"lesson_id" validate:"required"`
	CompletedAt  string    `json:"completed_at" validate:"required,datetime=2006-01-02T15:04:05"`
}
