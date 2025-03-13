package input

import (
	"context"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/domain"
)

type NotificationUseCase interface {
	CreateNotification(ctx context.Context, userID string, notificationType domain.NotificationType, title, content string, metadata map[string]string) (*domain.Notification, error)
	ScheduleNotification(ctx context.Context, notificationID string, scheduledTime time.Time) error
	CancelNotification(ctx context.Context, notificationID string) error
	GetNotification(ctx context.Context, notificationID string) (*domain.Notification, error)
	GetUserNotifications(ctx context.Context, userID string, limit, offset int) ([]*domain.Notification, int64, error)
	ProcessPendingNotifications(ctx context.Context) error
}
