package domain

import (
	"github.com/google/uuid"
)

type Module struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Order   int       `json:"order"`
	Lessons []Lesson  `json:"lessons"`
}
