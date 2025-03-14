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
	event, err := us.eventFactory.CreateEvent(eventType, notification)
	if err != nil {
		fmt.Printf("Error creating event: %v\n", err)
	}
	fmt.Printf("Event created: %v\n", event)

	err = us.notificationContext.SendNotification(ctx, notification, dto)
	if err != nil {
		fmt.Printf("Error sending notification: %v\n", err)
	} else {
		fmt.Printf("%s notification successfully sent\n", notification.Type)
	}
}
