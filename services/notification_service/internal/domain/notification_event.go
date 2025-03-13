package domain

import (
	"encoding/json"
	"time"
)

type EventType string

const (
	EventNotificationCreated   EventType = "notification.created"
	EventNotificationScheduled EventType = "notification.scheduled"
	EventNotificationSent      EventType = "notification.sent"
	EventNotificationFailed    EventType = "notification.failed"
	EventNotificationCancelled EventType = "notification.cancelled"
)

type NotificationEvent struct {
	EventID    string          `json:"event_id"`
	EventType  EventType       `json:"event_type"`
	OccurredAt time.Time       `json:"occurred_at"`
	Data       json.RawMessage `json:"data"`
}

type NotificationCreatedEvent struct {
	NotificationID string           `json:"notification_id"`
	UserID         string           `json:"user_id"`
	Type           NotificationType `json:"type"`
	Title          string           `json:"title"`
	ScheduledAt    *time.Time       `json:"scheduled_at,omitempty"`
}

type NotificationSentEvent struct {
	NotificationID string           `json:"notification_id"`
	UserID         string           `json:"user_id"`
	Type           NotificationType `json:"type"`
	SentAt         time.Time        `json:"sent_at"`
}

type NotificationFailedEvent struct {
	NotificationID string           `json:"notification_id"`
	UserID         string           `json:"user_id"`
	Type           NotificationType `json:"type"`
	FailedAt       time.Time        `json:"failed_at"`
	Reason         string           `json:"reason"`
}

func CreateNotificationCreatedEvent(notification *Notification) (*NotificationEvent, error) {
	data := NotificationCreatedEvent{
		NotificationID: notification.id,
		UserID:         notification.userID,
		Type:           notification.notificationtType,
		Title:          notification.title,
		ScheduledAt:    notification.scheduledAt,
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return &NotificationEvent{
		EventID:    notification.id,
		EventType:  EventNotificationCreated,
		OccurredAt: time.Now(),
		Data:       dataBytes,
	}, nil
}

func CreateNotificationSentEvent(notification *Notification) (*NotificationEvent, error) {
	data := NotificationSentEvent{
		NotificationID: notification.id,
		UserID:         notification.userID,
		Type:           notification.notificationtType,
		SentAt:         *notification.sentAt,
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return &NotificationEvent{
		EventID:    notification.id,
		EventType:  EventNotificationSent,
		OccurredAt: time.Now(),
		Data:       dataBytes,
	}, nil
}

func CreateNotificationFailedEvent(notification *Notification, reason string) (*NotificationEvent, error) {
	data := NotificationFailedEvent{
		NotificationID: notification.id,
		UserID:         notification.userID,
		Type:           notification.notificationtType,
		FailedAt:       time.Now(),
		Reason:         reason,
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return &NotificationEvent{
		EventID:    notification.id,
		EventType:  EventNotificationFailed,
		OccurredAt: time.Now(),
		Data:       dataBytes,
	}, nil
}
