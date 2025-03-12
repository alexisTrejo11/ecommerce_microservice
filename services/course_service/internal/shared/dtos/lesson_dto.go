package dtos

import (
	"time"

	"github.com/google/uuid"
)

// LessonInsertDTO represents the data required to create a new lesson within a module.
// @Description DTO used to insert a new lesson, including resources and other required fields.
// @SchemaExample { "title": "Lesson 1", "video_url": "https://example.com/video", "content": "This is a lesson on Go", "duration": 120, "order": 1, "module_id": "1c1cdb5c-d6e4-4fb0-9755-f30b9d4fbb8c", "is_preview": true, "resources": [{"title": "Resource 1", "lesson_id": "f2b02b99-4789-4c30-a9b9-b574fbcbd7cd", "type": "PDF", "url": "https://example.com/resource"}] }
type LessonInsertDTO struct {
	// Title is the title of the lesson.
	// @example Lesson 1
	Title string `json:"title" validate:"required"`

	// VideoURL is the URL for the lesson's video.
	// @example https://example.com/video
	VideoURL string `json:"video_url" validate:"omitempty,url"`

	// Content is the textual content of the lesson.
	// @example This is a lesson on Go
	Content string `json:"content"`

	// Duration is the duration of the lesson in minutes.
	// @example 120
	Duration int `json:"duration" validate:"required,min=1"`

	// Order is the sequence number of the lesson within the module.
	// @example 1
	Order int `json:"order" validate:"min=0"`

	// ModuleId is the unique identifier for the module the lesson belongs to.
	// @example 1c1cdb5c-d6e4-4fb0-9755-f30b9d4fbb8c
	ModuleId uuid.UUID `json:"module_id" validate:"required"`

	// IsPreview indicates whether the lesson is available as a preview.
	// @example true
	IsPreview bool `json:"is_preview"`

	// Resources is a list of resources associated with the lesson.
	// @example [{"title": "Resource 1", "lesson_id": "f2b02b99-4789-4c30-a9b9-b574fbcbd7cd", "type": "PDF", "url": "https://example.com/resource"}]
	Resources []ResourceInsertDTO `json:"resources" validate:"dive"`
}

// LessonDTO represents a lesson with all its details, including metadata such as creation and update times.
// @Description DTO that contains the full details of a lesson within a module, including resources and metadata.
// @SchemaExample { "id": "f2b02b99-4789-4c30-a9b9-b574fbcbd7cd", "title": "Lesson 1", "video_url": "https://example.com/video", "content": "This is a lesson on Go", "duration": 120, "order": 1, "is_preview": true, "created_at": "2025-03-12T10:00:00Z", "updated_at": "2025-03-12T10:00:00Z" }
type LessonDTO struct {
	// ID is the unique identifier for the lesson.
	// @example f2b02b99-4789-4c30-a9b9-b574fbcbd7cd
	ID uuid.UUID `json:"id"`

	// Title is the title of the lesson.
	// @example Lesson 1
	Title string `json:"title"`

	// VideoURL is the URL for the lesson's video.
	// @example https://example.com/video
	VideoURL string `json:"video_url"`

	// Content is the textual content of the lesson.
	// @example This is a lesson on Go
	Content string `json:"content"`

	// Duration is the duration of the lesson in minutes.
	// @example 120
	Duration int `json:"duration"`

	// Order is the sequence number of the lesson within the module.
	// @example 1
	Order int `json:"order"`

	// IsPreview indicates whether the lesson is available as a preview.
	// @example true
	IsPreview bool `json:"is_preview"`

	// CreatedAt is the timestamp when the lesson was created.
	// @example 2025-03-12T10:00:00Z
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt is the timestamp when the lesson was last updated.
	// @example 2025-03-12T10:00:00Z
	UpdatedAt time.Time `json:"updated_at"`
}
