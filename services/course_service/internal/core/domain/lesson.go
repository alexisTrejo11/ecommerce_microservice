package domain

import (
	"time"

	"github.com/google/uuid"
)

type Lesson struct {
	ID        uuid.UUID
	Title     string
	VideoURL  string
	Content   string
	Resources []Resource
	Duration  int
	Order     int
	IsPreview bool
	CreatedAt time.Time
}
