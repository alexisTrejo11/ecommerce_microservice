package progress

import (
	"time"

	"gorm.io/gorm"
)

type CompletedLesson struct {
	gorm.Model
	EnrollmentID uint       `json:"enrollment_id"`
	LessonID     uint       `json:"lesson_id"`
	IsCompleted  bool       `json:"is_completed"`
	CompletedAt  *time.Time `json:"completed_at"`
}

func (cl *CompletedLesson) MarkAsCompleted() {
	now := time.Now()
	cl.IsCompleted = true
	cl.CompletedAt = &now
}

func (cl *CompletedLesson) MarkAsIncomplete() {
	cl.IsCompleted = false
	cl.CompletedAt = nil
}
