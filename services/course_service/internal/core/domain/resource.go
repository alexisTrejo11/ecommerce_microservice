package domain

import (
	"github.com/google/uuid"
)

type Resource struct {
	ID    uuid.UUID    `json:"id"`
	Title string       `json:"title"`
	Type  ResourceType `json:"type"`
	URL   string       `json:"url"`
}

type ResourceType string

const (
	PDF    ResourceType = "PDF"
	SLIDES ResourceType = "SLIDES"
	LINK   ResourceType = "LINK"
	CODE   ResourceType = "CODE"
	OTHER  ResourceType = "OTHER"
)
