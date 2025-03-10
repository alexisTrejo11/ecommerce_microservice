package domain

import (
	"time"

	"github.com/google/uuid"
)

type Lesson struct {
	ID        uuid.UUID  `json:"id"`
	Title     string     `json:"title"`
	VideoURL  string     `json:"video_url"`
	Content   string     `json:"content"`
	Resources []Resource `json:"resources"`
	Duration  int        `json:"duration"`
	Order     int        `json:"order"`
	IsPreview bool       `json:"is_preview"`
	CreatedAt time.Time  `json:"created_at"`
}
