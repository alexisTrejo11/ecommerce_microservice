package models

import (
	progress "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/model"
	"gorm.io/gorm"
)

// Remove?
type Module struct {
	gorm.Model
	CourseID    uint     `json:"course_id"`
	Title       string   `json:"title"`
	OrderNumber uint     `json:"order_number"`
	Lessons     []Lesson `json:"lessons,omitempty"`
}

type Lesson struct {
	gorm.Model
	ModuleID         uint                       `json:"module_id"`
	Title            string                     `json:"title"`
	Content          string                     `json:"content"`
	Duration         uint                       `json:"duration_minutes"`
	OrderNumber      uint                       `json:"order_number"`
	ContentType      string                     `json:"content_type"` // video, text, quiz, etc.
	CompletedLessons []progress.CompletedLesson `json:"completed_lessons,omitempty"`
}
