package domain

import (
	"time"

	"github.com/google/uuid"
)

type Resource struct {
	ID        uuid.UUID
	LessonID  uuid.UUID
	Title     string
	Type      ResourceType
	URL       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ResourceType string

const (
	PDF    ResourceType = "PDF"
	SLIDES ResourceType = "SLIDES"
	LINK   ResourceType = "LINK"
	CODE   ResourceType = "CODE"
	OTHER  ResourceType = "OTHER"
)
