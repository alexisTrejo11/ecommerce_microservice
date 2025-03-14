package usecase

import (
	"context"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/mapper"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/utils"
	"github.com/google/uuid"
)

type NotificationUseCaseImpl struct {
	repository output.NotificationRepository
}

func NewNotificationUseCase(repository output.NotificationRepository) input.NotificationUseCase {
	return &NotificationUseCaseImpl{
		repository: repository,
	}
}

// Implement Strategy ?
func (us *NotificationUseCaseImpl) CreateNotification(
	ctx context.Context,
	dto dtos.NotificationMessageDTO) (*dtos.NotificationDTO, error) {

	return nil, nil
}

func (us *NotificationUseCaseImpl) ScheduleNotification(ctx context.Context, notificationID uuid.UUID, scheduledTime time.Time) error {
	return nil
}

func (us *NotificationUseCaseImpl) CancelNotification(ctx context.Context, notificationID uuid.UUID) error {
	err := us.repository.DeleteByID(ctx, notificationID)
	if err != nil {
		return err
	}

	return nil
}

func (us *NotificationUseCaseImpl) GetUserNotifications(ctx context.Context, userID uuid.UUID, page utils.Page) ([]*dtos.NotificationDTO, int64, error) {
	notifications, total, err := us.repository.GetByUserID(ctx, userID, page)
	if err != nil {
		return nil, 0, err
	}

	return mapper.ToNotificationDTOList(notifications), total, nil
}

// Run async schedules
func (us *NotificationUseCaseImpl) ProcessPendingNotifications(ctx context.Context) error {
	return nil
}

func (us *NotificationUseCaseImpl) GetNotification(ctx context.Context, notificationID uuid.UUID) (*dtos.NotificationDTO, error) {
	notification, err := us.repository.GetByID(ctx, notificationID)
	if err != nil {
		return nil, err
	}

	return mapper.ToNotificationDTO(notification), nil
}
