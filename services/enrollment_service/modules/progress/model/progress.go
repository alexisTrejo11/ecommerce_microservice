package progress

import (
	"time"

	"github.com/google/uuid"
)

type CompletedLesson struct {
	ID           uuid.UUID  `gorm:"type:char(36);primaryKey"`
	EnrollmentID uuid.UUID  `json:"enrollment_id"`
	LessonID     uuid.UUID  `json:"lesson_id"`
	IsCompleted  bool       `json:"is_completed"`
	CompletedAt  *time.Time `json:"completed_at"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `gorm:"index"`
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
