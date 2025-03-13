package domain

import (
	"errors"
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
	id                string
	userID            string
	notificationtType NotificationType
	title             string
	content           string
	metadata          map[string]string
	status            NotificationStatus
	createdAt         time.Time
	updatedAt         time.Time
	sentAt            *time.Time
	scheduledAt       *time.Time
}

func NewNotification(userID string, notificationType NotificationType, title, content string, metadata map[string]string) (*Notification, error) {
	if userID == "" {
		return nil, errors.New("user ID no puede estar vacío")
	}

	if title == "" {
		return nil, errors.New("título no puede estar vacío")
	}

	if content == "" {
		return nil, errors.New("contenido no puede estar vacío")
	}

	now := time.Now()

	return &Notification{
		id:                uuid.New().String(),
		userID:            userID,
		notificationtType: notificationType,
		title:             title,
		content:           content,
		metadata:          metadata,
		status:            StatusPending,
		createdAt:         now,
		updatedAt:         now,
	}, nil
}

func (n *Notification) ScheduleFor(scheduledTime time.Time) error {
	if scheduledTime.Before(time.Now()) {
		return errors.New("no se puede programar una notificación para el pasado")
	}

	n.scheduledAt = &scheduledTime
	n.updatedAt = time.Now()
	return nil
}

func (n *Notification) MarkAsSent() {
	now := time.Now()
	n.status = StatusSent
	n.sentAt = &now
	n.updatedAt = now
}

func (n *Notification) MarkAsFailed() {
	n.status = StatusFailed
	n.updatedAt = time.Now()
}

func (n *Notification) Cancel() error {
	if n.status != StatusPending {
		return errors.New("solo se pueden cancelar notificaciones pendientes")
	}

	n.status = StatusCancelled
	n.updatedAt = time.Now()
	return nil
}

func (n *Notification) IsScheduled() bool {
	return n.scheduledAt != nil && n.scheduledAt.After(time.Now())
}

func (n *Notification) ShouldSendNow() bool {
	return n.status == StatusPending &&
		(n.scheduledAt == nil || !n.scheduledAt.After(time.Now()))
}
