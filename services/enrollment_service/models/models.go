package models

import (
	"time"

	"gorm.io/gorm"
)

type Enrollment struct {
	gorm.Model
	UserID           uint      `json:"user_id"`
	CourseID         uint      `json:"course_id"`
	EnrollmentDate   time.Time `json:"enrollment_date"`
	CompletionDate   time.Time `json:"completion_date,omitempty"`
	CompletionStatus string    `json:"completion_status" gorm:"default:in_progress"`
	Progress         float64   `json:"progress" gorm:"default:0"`
	PaymentStatus    string    `json:"payment_status" gorm:"default:pending"`
	PaymentAmount    float64   `json:"payment_amount"`
	CertificateURL   string    `json:"certificate_url,omitempty"`
	LastAccessedAt   time.Time `json:"last_accessed_at"`

	CompletedLessons []CompletedLesson `json:"completed_lessons,omitempty"`
}

type CompletedLesson struct {
	gorm.Model
	EnrollmentID uint      `json:"enrollment_id"`
	LessonID     uint      `json:"lesson_id"`
	CompletedAt  time.Time `json:"completed_at"`
}

type Certificate struct {
	gorm.Model
	EnrollmentID   uint      `json:"enrollment_id"`
	IssuedAt       time.Time `json:"issued_at"`
	CertificateURL string    `json:"certificate_url"`
	ExpiresAt      time.Time `json:"expires_at,omitempty"`
}

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
	ModuleID         uint              `json:"module_id"`
	Title            string            `json:"title"`
	Content          string            `json:"content"`
	Duration         uint              `json:"duration_minutes"`
	OrderNumber      uint              `json:"order_number"`
	ContentType      string            `json:"content_type"` // video, text, quiz, etc.
	CompletedLessons []CompletedLesson `json:"completed_lessons,omitempty"`
}
