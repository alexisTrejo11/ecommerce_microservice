package domain

import (
	"time"

	"github.com/google/uuid"
)

type NotificationType string

const (
	TypeEmail NotificationType = "EMAIL"
	TypePush  NotificationType = "PUSH"
	TypeSMS   NotificationType = "SMS"
	TypeInApp NotificationType = "IN_APP"
)

type NotificationStatus string

const (
	StatusPending   NotificationStatus = "PENDING"
	StatusSent      NotificationStatus = "SENT"
	StatusFailed    NotificationStatus = "FAILED"
	StatusCancelled NotificationStatus = "CANCELLED"
)

type Notification struct {
	ID          string
	UserID      string
	Type        NotificationType
	Title       string
	Content     string
	Metadata    map[string]string
	Status      NotificationStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
	SentAt      *time.Time
	ScheduledAt *time.Time
}

func NewNotification(userID string, notificationType NotificationType, title, content string, metadata map[string]string) (*Notification, error) {
	if userID == "" {
		return nil, ErrUserIDRequired
	}

	if title == "" {
		return nil, ErrTitleRequired
	}

	if content == "" {
		return nil, ErrContentRequired
	}

	now := time.Now()

	return &Notification{
		ID:        uuid.New().String(),
		UserID:    userID,
		Type:      notificationType,
		Title:     title,
		Content:   content,
		Metadata:  metadata,
		Status:    StatusPending,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (n *Notification) ScheduleFor(scheduledTime time.Time) error {
	if scheduledTime.Before(time.Now()) {
		return ErrCannotSchedulePast
	}

	n.ScheduledAt = &scheduledTime
	n.UpdatedAt = time.Now()
	return nil
}

func (n *Notification) MarkAsSent() {
	now := time.Now()
	n.Status = StatusSent
	n.SentAt = &now
	n.UpdatedAt = now
}

func (n *Notification) MarkAsFailed() {
	n.Status = StatusFailed
	n.UpdatedAt = time.Now()
}

func (n *Notification) Cancel() error {
	if n.Status != StatusPending {
		return ErrCannotCancelNonPending
	}

	n.Status = StatusCancelled
	n.UpdatedAt = time.Now()
	return nil
}

func (n *Notification) IsScheduled() bool {
	return n.ScheduledAt != nil && n.ScheduledAt.After(time.Now())
}

func (n *Notification) ShouldSendNow() bool {
	return n.Status == StatusPending &&
		(n.ScheduledAt == nil || !n.ScheduledAt.After(time.Now()))
}
