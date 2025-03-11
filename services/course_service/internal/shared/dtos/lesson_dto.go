package dtos

import (
	"time"

	"github.com/google/uuid"
)

type LessonInsertDTO struct {
	Title     string              `json:"title" validate:"required"`
	VideoURL  string              `json:"video_url" validate:"omitempty,url"`
	Content   string              `json:"content"`
	Duration  int                 `json:"duration" validate:"required,min=1"`
	Order     int                 `json:"order" validate:"min=0"`
	ModuleId  uuid.UUID           `json:"module_id" validate:"required"`
	IsPreview bool                `json:"is_preview"`
	Resources []ResourceInsertDTO `json:"resources" validate:"dive"`
}

type LessonDTO struct {
	ID        uuid.UUID     `json:"id"`
	Title     string        `json:"title"`
	VideoURL  string        `json:"video_url"`
	Content   string        `json:"content"`
	Duration  int           `json:"duration"`
	Order     int           `json:"order"`
	IsPreview bool          `json:"is_preview"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Resources []ResourceDTO `json:"resources"`
}
