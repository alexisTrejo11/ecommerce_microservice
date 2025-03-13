package domain

import (
	"time"

	"github.com/google/uuid"
)

type SubscriptionStatus string

const (
	Active   SubscriptionStatus = "ACTIVE"
	Expired  SubscriptionStatus = "EXPIRED"
	Canceled SubscriptionStatus = "CANCELED"
)

type Subscription struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	CourseID  uuid.UUID
	Status    SubscriptionStatus
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSubscription(userID uuid.UUID, courseID uuid.UUID, startDate, endDate time.Time) *Subscription {
	return &Subscription{
		ID:        uuid.New(),
		UserID:    userID,
		CourseID:  courseID,
		Status:    Active,
		StartDate: startDate,
		EndDate:   endDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (s *Subscription) Cancel() {
	s.Status = Canceled
	s.UpdatedAt = time.Now()
}
