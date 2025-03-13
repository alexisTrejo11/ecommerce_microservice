package output

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/domain"
)

type NotificationSender interface {
	SendNotification(ctx context.Context, notification *domain.Notification) error
}

type EventPublisher interface {
	PublishEvent(ctx context.Context, event *domain.NotificationEvent) error
}
