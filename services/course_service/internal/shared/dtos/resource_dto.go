package dtos

import (
	"time"

	"github.com/google/uuid"
)

type ResourceDTO struct {
	ID        uuid.UUID `json:"id"`
	LessonID  uuid.UUID `json:"lesson_id"`
	Title     string    `json:"title"`
	Type      string    `json:"type"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResourceInsertDTO struct {
	Title    string    `json:"title" validate:"required"`
	LessonID uuid.UUID `json:"lesson_id" validate:"required"`
	Type     string    `json:"type" validate:"required,oneof=PDF SLIDES LINK CODE OTHER"`
	URL      string    `json:"url" validate:"required,url"`
}
