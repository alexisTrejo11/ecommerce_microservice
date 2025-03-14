package utils

import (
	"fmt"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/domain"
)

type NotificationEventFactory struct{}

func (f *NotificationEventFactory) CreateEvent(eventType domain.EventType, notification domain.Notification) (*domain.NotificationEvent, error) {
	switch eventType {
	case domain.EventNotificationCreated:
		return domain.CreateNotificationCreatedEvent(&notification)
	case domain.EventNotificationScheduled:
		return domain.CreateNotificationScheduledEvent(&notification)
	case domain.EventNotificationSent:
		return domain.CreateNotificationSentEvent(&notification)
	case domain.EventNotificationCancelled:
		return domain.CreateNotificationCancelledEvent(&notification)
	default:
		return nil, fmt.Errorf("not supported notification type")
	}
}
