package dtos

import "time"

type LessonInsertDTO struct {
	Title     string              `json:"title" binding:"required"`
	VideoURL  string              `json:"video_url" binding:"omitempty,url"`
	Content   string              `json:"content"`
	Duration  int                 `json:"duration" binding:"required,min=1"`
	Order     int                 `json:"order" binding:"required,min=0"`
	IsPreview bool                `json:"is_preview"`
	Resources []ResourceInsertDTO `json:"resources" binding:"dive"`
}

type LessonDTO struct {
	ID        string        `json:"id"`
	Title     string        `json:"title"`
	VideoURL  string        `json:"video_url"`
	Content   string        `json:"content"`
	Duration  int           `json:"duration"`
	Order     int           `json:"order"`
	IsPreview bool          `json:"is_preview"`
	CreatedAt time.Time     `json:"created_at"`
	Resources []ResourceDTO `json:"resources"`
}
