package usecase

import (
	"context"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/domain"
)

type NotificationUseCaseImpl struct {
	repository output.NotificationRepository
}

func NewNotificationUseCase(repository output.NotificationRepository) input.NotificationUseCase {
	return &NotificationUseCaseImpl{
		repository: repository,
	}
}
func (us *NotificationUseCaseImpl) CreateNotification(
	ctx context.Context,
	userID string,
	notificationType domain.NotificationType,
	title, content string,
	metadata map[string]string) (*domain.Notification, error) {

	return nil, nil
}

func (us *NotificationUseCaseImpl) ScheduleNotification(
	ctx context.Context,
	notificationID string,
	scheduledTime time.Time) error {
	return nil
}

func (us *NotificationUseCaseImpl) CancelNotification(ctx context.Context, notificationID string) error {
	return nil
}
func (us *NotificationUseCaseImpl) GetUserNotifications(ctx context.Context, userID string, limit, offset int) ([]*domain.Notification, int64, error) {
	return nil, 0, nil
}

func (us *NotificationUseCaseImpl) ProcessPendingNotifications(ctx context.Context) error {
	return nil
}

func (us *NotificationUseCaseImpl) GetNotification(ctx context.Context, notificationID string) (*domain.Notification, error) {
	return nil, nil
}
