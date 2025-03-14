package output

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/utils"
	"github.com/google/uuid"
)

type NotificationRepository interface {
	Save(ctx context.Context, notification *domain.Notification) error
	GetByID(ctx context.Context, notificationID uuid.UUID) (*domain.Notification, error)
	GetByUserID(ctx context.Context, userID uuid.UUID, page utils.Page) (*[]domain.Notification, int64, error)
	GetPendingNotifications(ctx context.Context, page utils.Page) (*[]domain.Notification, error)
	DeleteByID(ctx context.Context, notificationID uuid.UUID) error
}
