package domain

import "time"

type NotificationType string
type NotificationStatus string

const (
	Email NotificationType = "EMAIL"
	Sms   NotificationType = "SMS"
	PUSH  NotificationType = "PUSH"
)

const (
	Pending NotificationStatus = "PENDING"
	Sent    NotificationStatus = "SENT"
	Failed  NotificationStatus = "FAILED"
)

type Notification struct {
	ID        uint
	UserID    string
	Type      NotificationType
	Content   string
	Status    NotificationStatus
	CreatedAt time.Time
	UpdatedAt time.Time
	SentAt    *time.Time
}
