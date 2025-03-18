package progress

import (
	"time"

	"gorm.io/gorm"
)

type CompletedLesson struct {
	gorm.Model
	EnrollmentID uint      `json:"enrollment_id"`
	LessonID     uint      `json:"lesson_id"`
	CompletedAt  time.Time `json:"completed_at"`
}
