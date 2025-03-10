package dtos

import "time"

type CourseInsertDTO struct {
	Title        string            `json:"title" binding:"required,min=3"`
	Slug         string            `json:"slug" binding:"required,alphanumdash"`
	Description  string            `json:"description" binding:"required"`
	ThumbnailURL string            `json:"thumbnail_url" binding:"omitempty,url"`
	Category     string            `json:"category" binding:"required"`
	Level        string            `json:"level" binding:"required,oneof=BEGINNER INTERMEDIATE ADVANCED"`
	Language     string            `json:"language" binding:"required"`
	InstructorID string            `json:"instructor_id" binding:"required,uuid"`
	Tags         []string          `json:"tags"`
	Price        float64           `json:"price" binding:"gte=0"`
	IsFree       bool              `json:"is_free"`
	Modules      []ModuleInsertDTO `json:"modules" binding:"dive"`
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
