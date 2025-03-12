package dtos

import (
	"time"

	"github.com/google/uuid"
)

// CourseInsertDTO represents the data required to create a new course.
// @Description DTO used to insert a new course with necessary fields including title, description, level, and more.
// @SchemaExample { "title": "Go Programming", "description": "Learn Go programming from scratch.", "thumbnail_url": "https://example.com/thumbnail.jpg", "level": "BEGINNER", "category": "PROGRAMMING", "language": "English", "instructor_id": "8c1d73a3-4a33-4c60-914f-76b91b3510ad", "tags": ["Go", "Programming"], "price": 50.0, "is_free": false }
type CourseInsertDTO struct {
	// Title is the name of the course.
	// @example Go Programming
	Title string `json:"title" validate:"required,min=3"`

	// Description is a detailed explanation of the course.
	// @example Learn Go programming from scratch.
	Description string `json:"description" validate:"required"`

	// ThumbnailURL is the URL to the course's thumbnail image.
	// @example https://example.com/thumbnail.jpg
	ThumbnailURL string `json:"thumbnail_url" validate:"omitempty,url"`

	// Level indicates the difficulty level of the course.
	// @example BEGINNER
	Level string `json:"level" validate:"required,oneof=BEGINNER INTERMEDIATE ADVANCED"`

	// Category specifies the course's category.
	// @example PROGRAMMING
	Category string `json:"category" validate:"required,oneof=PROGRAMMING DESIGN_SOFTWARE ENGINEER_SOFTWARE ARCHITECTURE_SOFTWARE AI ART MARKETING SOCIAL_NETWORK LANGUAGE"`

	// Language is the language in which the course is taught.
	// @example English
	Language string `json:"language" validate:"required"`

	// InstructorID is the unique identifier of the course instructor.
	// @example 8c1d73a3-4a33-4c60-914f-76b91b3510ad
	InstructorID uuid.UUID `json:"instructor_id" validate:"required,uuid"`

	// Tags are a list of relevant keywords related to the course.
	// @example ["Go", "Programming"]
	Tags []string `json:"tags"`

	// Price is the cost of the course.
	// @example 50.0
	Price float64 `json:"price" validate:"gte=0"`

	// IsFree indicates whether the course is free or paid.
	// @example false
	IsFree bool `json:"is_free"`
}

// CourseDTO represents a full course with details including instructor, pricing, and enrollment statistics.
// @Description DTO that contains all the details of a course, including its modules, instructor, price, and more.
// @SchemaExample { "id": "abc123", "title": "Go Programming", "slug": "go-programming", "description": "Learn Go programming from scratch.", "thumbnail_url": "https://example.com/thumbnail.jpg", "category": "PROGRAMMING", "level": "BEGINNER", "language": "English", "instructor_id": "8c1d73a3-4a33-4c60-914f-76b91b3510ad", "tags": ["Go", "Programming"], "price": 50.0, "is_free": false, "is_published": true, "enrollment_count": 100, "rating": 4.5, "review_count": 25, "published_at": "2025-03-12T10:00:00Z", "created_at": "2025-03-12T10:00:00Z", "updated_at": "2025-03-12T10:00:00Z", "modules": [...] }
type CourseDTO struct {
	// ID is the unique identifier for the course.
	// @example abc123
	ID string `json:"id"`

	// Title is the name of the course.
	// @example Go Programming
	Title string `json:"title"`

	// Slug is a URL-friendly version of the title.
	// @example go-programming
	Slug string `json:"slug"`

	// Description is a detailed explanation of the course.
	// @example Learn Go programming from scratch.
	Description string `json:"description"`

	// ThumbnailURL is the URL to the course's thumbnail image.
	// @example https://example.com/thumbnail.jpg
	ThumbnailURL string `json:"thumbnail_url"`

	// Category specifies the course's category.
	// @example PROGRAMMING
	Category string `json:"category"`

	// Level indicates the difficulty level of the course.
	// @example BEGINNER
	Level string `json:"level"`

	// Language is the language in which the course is taught.
	// @example English
	Language string `json:"language"`

	// InstructorID is the unique identifier of the course instructor.
	// @example 8c1d73a3-4a33-4c60-914f-76b91b3510ad
	InstructorID string `json:"instructor_id"`

	// Tags are a list of relevant keywords related to the course.
	// @example ["Go", "Programming"]
	Tags []string `json:"tags"`

	// Price is the cost of the course.
	// @example 50.0
	Price float64 `json:"price"`

	// IsFree indicates whether the course is free or paid.
	// @example false
	IsFree bool `json:"is_free"`

	// IsPublished indicates whether the course is published or not.
	// @example true
	IsPublished bool `json:"is_published"`

	// EnrollmentCount is the number of students enrolled in the course.
	// @example 100
	EnrollmentCount int `json:"enrollment_count"`

	// Rating is the average rating of the course.
	// @example 4.5
	Rating float64 `json:"rating"`

	// ReviewCount is the number of reviews the course has received.
	// @example 25
	ReviewCount int `json:"review_count"`

	// PublishedAt is the timestamp when the course was published.
	// @example 2025-03-12T10:00:00Z
	PublishedAt *time.Time `json:"published_at,omitempty"`

	// CreatedAt is the timestamp when the course was created.
	// @example 2025-03-12T10:00:00Z
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt is the timestamp when the course was last updated.
	// @example 2025-03-12T10:00:00Z
	UpdatedAt time.Time `json:"updated_at"`

	// Modules are the modules associated with the course.
	// @example [...]
	Modules []ModuleDTO `json:"modules"`
}
