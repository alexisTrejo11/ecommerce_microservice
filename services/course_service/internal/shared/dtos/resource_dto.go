package dtos

import (
	"time"

	"github.com/google/uuid"
)

// ResourceDTO represents a resource associated with a lesson in the system.
// @Description DTO that contains information about a specific resource.
// @SchemaExample { "id": "d4fc16ff-4b2a-4a47-9cd6-8579d1f2cf0c", "lesson_id": "b7a92b1d-6a84-43f9-bfa0-c3703d13bc3f", "title": "Introduction to Go", "type": "PDF", "url": "https://example.com/intro-to-go.pdf", "created_at": "2025-03-01T12:34:56Z", "updated_at": "2025-03-01T12:34:56Z" }
type ResourceDTO struct {
	// ID is the unique identifier for the resource.
	// @example d4fc16ff-4b2a-4a47-9cd6-8579d1f2cf0c
	ID uuid.UUID `json:"id"`

	// LessonID is the unique identifier for the lesson associated with the resource.
	// @example b7a92b1d-6a84-43f9-bfa0-c3703d13bc3f
	LessonID uuid.UUID `json:"lesson_id"`

	// Title is the name or title of the resource.
	// @example Introduction to Go
	Title string `json:"title"`

	// Type refers to the type of resource (e.g., PDF, SLIDES, LINK, CODE, OTHER).
	// @example PDF
	Type string `json:"type"`

	// URL is the link or address where the resource can be accessed.
	// @example https://example.com/intro-to-go.pdf
	URL string `json:"url"`

	// CreatedAt is the timestamp when the resource was created.
	// @example 2025-03-01T12:34:56Z
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt is the timestamp when the resource was last updated.
	// @example 2025-03-01T12:34:56Z
	UpdatedAt time.Time `json:"updated_at"`
}

// ResourceInsertDTO represents the data needed to create a new resource.
// @Description DTO used to insert a new resource with the required fields.
// @SchemaExample { "title": "Introduction to Go", "lesson_id": "b7a92b1d-6a84-43f9-bfa0-c3703d13bc3f", "type": "PDF", "url": "https://example.com/intro-to-go.pdf" }
type ResourceInsertDTO struct {
	// Title is the name or title of the resource.
	// @example Introduction to Go
	Title string `json:"title" validate:"required"`

	// LessonID is the unique identifier for the lesson associated with the resource.
	// @example b7a92b1d-6a84-43f9-bfa0-c3703d13bc3f
	LessonID uuid.UUID `json:"lesson_id" validate:"required"`

	// Type refers to the type of resource (e.g., PDF, SLIDES, LINK, CODE, OTHER).
	// @example PDF
	Type string `json:"type" validate:"required,oneof=PDF SLIDES LINK CODE OTHER"`

	// URL is the link or address where the resource can be accessed.
	// @example https://example.com/intro-to-go.pdf
	URL string `json:"url" validate:"required,url"`
}
