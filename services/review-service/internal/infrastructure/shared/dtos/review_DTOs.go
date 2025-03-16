package dtos

import (
	"time"

	"github.com/google/uuid"
)

type ReviewDTO struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	CourseID   uuid.UUID `json:"course_id"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	IsApproved bool      `json:"is_approved"`
}

type ReviewInsertDTO struct {
	UserID   uuid.UUID `json:"user_id" validate:"required"`
	CourseID uuid.UUID `json:"course_id" validate:"required"`
	Rating   int       `json:"rating" validate:"gte=1,lte=5"`
	Comment  string    `json:"comment" validate:"omitempty,max=500"`
}
