package output

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/domain"
)

type NotificationRepository interface {
	Save(ctx context.Context, notification *domain.Notification) error
	GetByID(ctx context.Context, notificationID string) (*domain.Notification, error)
	GetByUserID(ctx context.Context, userID string, limit, offset int) ([]*domain.Notification, int64, error)
	GetPendingNotifications(ctx context.Context, limit int) ([]*domain.Notification, error)
	DeleteByID(ctx context.Context, notificationID string) error
}
