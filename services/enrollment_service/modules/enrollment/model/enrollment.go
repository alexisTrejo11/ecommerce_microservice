package enrollment

import (
	"errors"
	"fmt"
	"time"

	progress "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/model"
	"github.com/google/uuid"
)

type CompletionStatus string

var (
	STARTING  CompletionStatus = "STARTING"
	ON_GOING  CompletionStatus = "ON_GOING"
	ON_HOLD   CompletionStatus = "ON_HOLD"
	CANCELLED CompletionStatus = "CANCELLED"
	FINISHED  CompletionStatus = "FINISHED"
)

// Properties are public cause migration need public fields
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
		UserID:           userID,
		CourseID:         courseID,
		Progress:         0,
		EnrollmentDate:   now,
		LastAccessedAt:   now,
		CompletionStatus: STARTING,
	}
}

func (e *Enrollment) Complete() {
	e.CompletionStatus = FINISHED
	e.CompletionDate = time.Now()
}

func (e *Enrollment) IsEnrollmentFinished() bool {
	return e.CompletionStatus == FINISHED
}

func (e *Enrollment) UpdateLastAccessedAt() {
	e.LastAccessedAt = time.Now()
}

func (e *Enrollment) SetCompletionStatus(status CompletionStatus) error {
	if !e.CanUpdateStatus() {
		return fmt.Errorf("cannot update status of a finished enrollment")
	}

	switch status {
	case STARTING, ON_GOING, ON_HOLD, CANCELLED:
		e.CompletionStatus = status
	default:
		return fmt.Errorf("invalid completion status")
	}

	return nil
}

func (e *Enrollment) CanUpdateStatus() bool {
	return e.CompletionStatus != FINISHED
}

func (e *Enrollment) CalculateProgress(totalLessons int) error {
	if totalLessons <= 0 {
		return fmt.Errorf("total lessons must be greater than zero")
	}

	completedLessons := len(e.CompletedLessons)
	e.Progress = float64(completedLessons) / float64(totalLessons) * 100

	if e.Progress == 100 {
		e.Complete()
	}

	return nil
}

func (e *Enrollment) Cancel() error {
	limitCancelDate := e.EnrollmentDate.Add(time.Hour * 48)

	if limitCancelDate.After(time.Now()) {
		return errors.New("too late to cancel your enrollment")
	}

	// TODO: Check les than %10

	if err := e.SetCompletionStatus(CANCELLED); err != nil {
		return err
	}

	return nil
}
