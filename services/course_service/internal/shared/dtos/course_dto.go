package dtos

import (
	"time"

	"github.com/google/uuid"
)

type CourseInsertDTO struct {
	Title        string    `json:"title" validate:"required,min=3"`
	Description  string    `json:"description" validate:"required"`
	ThumbnailURL string    `json:"thumbnail_url" validate:"omitempty,url"`
	Level        string    `json:"level" validate:"required,oneof=BEGINNER INTERMEDIATE ADVANCED"`
	Category     string    `json:"category" validate:"required,oneof=PROGRAMMING DESIGN_SOFTWARE ENGINEER_SOFTWARE ARCHITECTURE_SOFTWARE AI ART MARKETING SOCIAL_NETWORK LANGUAGE"`
	Language     string    `json:"language" validate:"required"`
	InstructorID uuid.UUID `json:"instructor_id" validate:"required,uuid"`
	Tags         []string  `json:"tags"`
	Price        float64   `json:"price" validate:"gte=0"`
	IsFree       bool      `json:"is_free"`
}

type CourseDTO struct {
	ID              string      `json:"id"`
	Title           string      `json:"title"`
	Slug            string      `json:"slug"`
	Description     string      `json:"description"`
	ThumbnailURL    string      `json:"thumbnail_url"`
	Category        string      `json:"category"`
	Level           string      `json:"level"`
	Language        string      `json:"language"`
	InstructorID    string      `json:"instructor_id"`
	Tags            []string    `json:"tags"`
	Price           float64     `json:"price"`
	IsFree          bool        `json:"is_free"`
	IsPublished     bool        `json:"is_published"`
	EnrollmentCount int         `json:"enrollment_count"`
	Rating          float64     `json:"rating"`
	ReviewCount     int         `json:"review_count"`
	PublishedAt     *time.Time  `json:"published_at,omitempty"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	Modules         []ModuleDTO `json:"modules"`
}
