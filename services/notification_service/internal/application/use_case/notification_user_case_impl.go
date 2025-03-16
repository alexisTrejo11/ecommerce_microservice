package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/mapper"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/utils"
	logging "github.com/alexisTrejo11/ecommerce_microservice/notification-service/pkg/log"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/pkg/notification"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/pkg/sms"
	"github.com/google/uuid"
)

type NotificationUseCaseImpl struct {
	repository          output.NotificationRepository
	emailUseCase        input.EmailUseCase
	smsService          sms.SMSService
	eventFactory        utils.NotificationEventFactory
	notificationContext *notification.NotificationContext
}

func NewNotificationUseCase(
	repository output.NotificationRepository,
	emailUseCase input.EmailUseCase,
	smsService sms.SMSService,
) input.NotificationUseCase {
	return &NotificationUseCaseImpl{
		repository:   repository,
		emailUseCase: emailUseCase,
		smsService:   smsService,
	}
}

func (us *NotificationUseCaseImpl) CreateNotification(ctx context.Context, dto dtos.NotificationMessageDTO) (*dtos.NotificationDTO, error) {
	notification := mapper.ToDomainFromMessageDTO(&dto)

	go us.sendNotification(ctx, *notification, domain.EventNotificationCreated, dto)

	// Proccess Events

	return mapper.ToNotificationDTO(notification), nil
}

func (us *NotificationUseCaseImpl) ScheduleNotification(ctx context.Context, notificationID uuid.UUID, scheduledTime time.Time) error {
	notification, err := us.repository.GetByID(ctx, notificationID)
	if err != nil {
		return err
	}

	if err := notification.ScheduleFor(scheduledTime); err != nil {
		return err
	}

	us.repository.Save(ctx, notification)
	// Proccess Events

	return nil
}

func (us *NotificationUseCaseImpl) CancelNotification(ctx context.Context, notificationID uuid.UUID) error {
	notification, err := us.repository.GetByID(ctx, notificationID)
	if err != nil {
		return err
	}

	if err := notification.Cancel(); err != nil {
		return err
	}

	us.repository.Save(ctx, notification)

	// Proccess Events

	return nil
}

func (us *NotificationUseCaseImpl) GetUserNotifications(ctx context.Context, userID uuid.UUID, page utils.Page) ([]*dtos.NotificationDTO, int64, error) {
	notifications, total, err := us.repository.GetByUserID(ctx, userID, page)
	if err != nil {
		return nil, 0, err
	}

	return mapper.ToNotificationDTOList(*notifications), total, nil
}

func (us *NotificationUseCaseImpl) ProcessPendingNotifications(ctx context.Context) error {
	notifications, err := us.repository.GetPendingNotifications(ctx, utils.Page{PageNumber: 1, PageSize: 1000})
	if err != nil {
		return err
	}

	// Proccess Events

	fmt.Printf("notifications: %v\n", notifications)

	return nil
}

func (us *NotificationUseCaseImpl) GetNotification(ctx context.Context, notificationID uuid.UUID) (*dtos.NotificationDTO, error) {
	notification, err := us.repository.GetByID(ctx, notificationID)
	if err != nil {
		return nil, err
	}

	return mapper.ToNotificationDTO(notification), nil
}

func (us *NotificationUseCaseImpl) sendNotification(ctx context.Context, notification domain.Notification, eventType domain.EventType, dto dtos.NotificationMessageDTO) {
	logging.Logger.Info().Str("action", "create_notification_event").
		Str("event_type", string(eventType)).
		Str("notification_id", notification.ID).
		Msg("Creating notification event")

	_, err := us.eventFactory.CreateEvent(eventType, notification)
	if err != nil {
		logging.LogError("create_notification_event", "Error creating event", map[string]interface{}{
			"error":           err.Error(),
			"event_type":      string(eventType),
			"notification_id": notification.ID,
		})
		return
	}

	logging.Logger.Info().
		Str("action", "event_created").
		Str("event_type", string(eventType)).
		Str("notification_id", notification.ID).
		Msg("Notification event created successfully")

	err = us.notificationContext.SendNotification(ctx, notification, dto)
	if err != nil {
		logging.LogError("send_notification", "Error sending notification", map[string]interface{}{
			"error":           err.Error(),
			"notification_id": notification.ID,
			"type":            notification.Type,
		})
	} else {
		logging.LogSuccess("send_notification", "Notification successfully sent", map[string]interface{}{
			"notification_id": notification.ID,
			"type":            notification.Type,
		})
	}
}
