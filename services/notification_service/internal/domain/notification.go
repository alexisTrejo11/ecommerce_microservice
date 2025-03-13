// internal/domain/notification.go
package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// NotificationType representa los tipos de notificaciones soportados
type NotificationType string

const (
	TypeEmail NotificationType = "EMAIL"
	TypePush  NotificationType = "PUSH"
	TypeSMS   NotificationType = "SMS"
	TypeInApp NotificationType = "IN_APP"
)

// NotificationStatus representa los estados posibles de una notificación
type NotificationStatus string

const (
	StatusPending   NotificationStatus = "PENDING"
	StatusSent      NotificationStatus = "SENT"
	StatusFailed    NotificationStatus = "FAILED"
	StatusCancelled NotificationStatus = "CANCELLED"
)

// Notification es la entidad principal del dominio
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

// NewNotification crea una nueva notificación con validaciones
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

// ScheduleFor programa la notificación para enviarse en un momento específico
func (n *Notification) ScheduleFor(scheduledTime time.Time) error {
	if scheduledTime.Before(time.Now()) {
		return errors.New("no se puede programar una notificación para el pasado")
	}

	n.ScheduledAt = &scheduledTime
	n.UpdatedAt = time.Now()
	return nil
}

// MarkAsSent marca la notificación como enviada
func (n *Notification) MarkAsSent() {
	now := time.Now()
	n.Status = StatusSent
	n.SentAt = &now
	n.UpdatedAt = now
}

// MarkAsFailed marca la notificación como fallida
func (n *Notification) MarkAsFailed() {
	n.Status = StatusFailed
	n.UpdatedAt = time.Now()
}

// Cancel cancela una notificación pendiente
func (n *Notification) Cancel() error {
	if n.Status != StatusPending {
		return errors.New("solo se pueden cancelar notificaciones pendientes")
	}

	n.Status = StatusCancelled
	n.UpdatedAt = time.Now()
	return nil
}

// IsScheduled verifica si la notificación está programada para el futuro
func (n *Notification) IsScheduled() bool {
	return n.ScheduledAt != nil && n.ScheduledAt.After(time.Now())
}

// ShouldSendNow verifica si la notificación debe enviarse ahora
func (n *Notification) ShouldSendNow() bool {
	return n.Status == StatusPending &&
		(n.ScheduledAt == nil || !n.ScheduledAt.After(time.Now()))
}
