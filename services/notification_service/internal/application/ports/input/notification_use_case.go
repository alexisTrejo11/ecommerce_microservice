package input

import (
	"context"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/utils"
	"github.com/google/uuid"
)

type NotificationUseCase interface {
	CreateNotification(ctx context.Context, dto dtos.NotificationMessageDTO) (*dtos.NotificationDTO, error)
	ScheduleNotification(ctx context.Context, notificationID uuid.UUID, scheduledTime time.Time) error
	CancelNotification(ctx context.Context, notificationID uuid.UUID) error
	GetNotification(ctx context.Context, notificationID uuid.UUID) (*dtos.NotificationDTO, error)
	GetUserNotifications(ctx context.Context, userID uuid.UUID, page utils.Page) ([]*dtos.NotificationDTO, int64, error)
	ProcessPendingNotifications(ctx context.Context) error
}
