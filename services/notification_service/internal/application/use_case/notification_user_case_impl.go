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
	"github.com/google/uuid"
)

type NotificationUseCaseImpl struct {
	repository   output.NotificationRepository
	emailUseCase input.EmailUseCase
}

func NewNotificationUseCase(repository output.NotificationRepository, emailUseCase input.EmailUseCase) input.NotificationUseCase {
	return &NotificationUseCaseImpl{
		repository:   repository,
		emailUseCase: emailUseCase,
	}
}

func (us *NotificationUseCaseImpl) CreateNotification(ctx context.Context, dto dtos.NotificationMessageDTO) (*dtos.NotificationDTO, error) {
	notification := mapper.ToDomainFromMessageDTO(&dto)

	go us.send(ctx, *notification, domain.EventNotificationCreated, dto)

	return mapper.ToNotificationDTO(notification), nil
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

	return mapper.ToNotificationDTOList(*notifications), total, nil
}

func (us *NotificationUseCaseImpl) ProcessPendingNotifications(ctx context.Context) error {
	notifications, err := us.repository.GetPendingNotifications(ctx, utils.Page{PageNumber: 1, PageSize: 1000})
	if err != nil {
		return err
	}

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

func (us *NotificationUseCaseImpl) send(ctx context.Context, notification domain.Notification, eventType domain.EventType, dto dtos.NotificationMessageDTO) {
	event, err := createEvent(eventType, notification)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("event: %v\n", event)

	switch notification.Type {
	case domain.TypeEmail:
		if err := us.emailUseCase.SendEmail(ctx, dto); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Email Successfully Sended")
	case domain.TypeSMS:
		return
	case domain.TypeInApp:
		if err := us.repository.Save(ctx, &notification); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Email Successfully Created")
	case domain.TypePush:
		return
	}

}

func createEvent(eventType domain.EventType, notification domain.Notification) (*domain.NotificationEvent, error) {
	switch eventType {
	case domain.EventNotificationCreated:
		event, err := domain.CreateNotificationCreatedEvent(&notification)
		if err != nil {
			return nil, err
		}
		return event, nil
	case domain.EventNotificationScheduled:
		event, err := domain.CreateNotificationCreatedEvent(&notification)
		if err != nil {
			return nil, err
		}
		return event, nil
	case domain.EventNotificationSent:
		event, err := domain.CreateNotificationCreatedEvent(&notification)
		if err != nil {
			return nil, err
		}
		return event, nil
	case domain.EventNotificationCancelled:
		event, err := domain.CreateNotificationCreatedEvent(&notification)
		if err != nil {
			return nil, err
		}
		return event, nil
	default:
		return nil, fmt.Errorf("not supported notification type")
	}
}
