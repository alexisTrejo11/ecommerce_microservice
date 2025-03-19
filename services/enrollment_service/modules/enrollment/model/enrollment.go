package enrollment

import (
	"time"

	progress "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/model"
	"github.com/google/uuid"
)

type CompletionStatus string

var (
	STARTING CompletionStatus = "STARTING"
	ON_GOING CompletionStatus = "ON_GOING"
	ON_HOLD  CompletionStatus = "ON_HOLD"
	FINISHED CompletionStatus = "FINISHED"
)

type Enrollment struct {
	ID               uuid.UUID        `gorm:"type:char(36);primaryKey"`
	UserID           uuid.UUID        `json:"user_id"`
	CourseID         uuid.UUID        `json:"course_id"`
	EnrollmentDate   time.Time        `json:"enrollment_date"`
	CompletionDate   time.Time        `json:"completion_date,omitempty"`
	CompletionStatus CompletionStatus `json:"completion_status" gorm:"default:in_progress"`
	Progress         float64          `json:"progress" gorm:"default:0"`
	CertificateID    uuid.UUID        `json:"certificate_id"`
	LastAccessedAt   time.Time        `json:"last_accessed_at"`

	CompletedLessons []progress.CompletedLesson `json:"completed_lessons,omitempty"`
}

func NewEnrollment(userID uuid.UUID, courseID uuid.UUID) *Enrollment {
	now := time.Now()
	return &Enrollment{
		UserID:         userID,
		CourseID:       courseID,
		Progress:       0,
		EnrollmentDate: now,
		LastAccessedAt: now,
	}
}

func (e *Enrollment) MarkAsCompleted() {
	e.CompletionStatus = FINISHED
}
