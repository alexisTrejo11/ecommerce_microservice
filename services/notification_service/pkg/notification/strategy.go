package notification

import (
	"context"
	"fmt"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/dtos"
)

type NotificationStrategy interface {
	Send(ctx context.Context, notification domain.Notification, dto dtos.NotificationMessageDTO) error
}

type EmailStrategy struct {
	emailUseCase EmailSender
}

type EmailSender interface {
	SendEmail(ctx context.Context, dto dtos.NotificationMessageDTO) error
}

func NewEmailStrategy(emailService EmailSender) *EmailStrategy {
	return &EmailStrategy{emailUseCase: emailService}
}

func (s *EmailStrategy) Send(ctx context.Context, notification domain.Notification, dto dtos.NotificationMessageDTO) error {
	if err := s.emailUseCase.SendEmail(ctx, dto); err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}
	return nil
}

type SMSStrategy struct {
	smsService SMSSender
}

type SMSSender interface {
	SendSMS(to string, message string, from string) (string, error)
	From() string
}

func NewSMSStrategy(smsService SMSSender) *SMSStrategy {
	return &SMSStrategy{smsService: smsService}
}

func (s *SMSStrategy) Send(ctx context.Context, notification domain.Notification, dto dtos.NotificationMessageDTO) error {
	if dto.UserData.Phone == nil {
		return fmt.Errorf("phone number not provided")
	}

	sid, err := s.smsService.SendSMS(*dto.UserData.Phone, notification.Content, s.smsService.From())
	if err != nil {
		return fmt.Errorf("error sending SMS: %w", err)
	}

	// SID SAVE
	_ = sid
	return nil
}

type InAppStrategy struct {
	repository NotificationRepository
}

type NotificationRepository interface {
	Save(ctx context.Context, notification *domain.Notification) error
}

func NewInAppStrategy(repository NotificationRepository) *InAppStrategy {
	return &InAppStrategy{repository: repository}
}

func (s *InAppStrategy) Send(ctx context.Context, notification domain.Notification, dto dtos.NotificationMessageDTO) error {
	if err := s.repository.Save(ctx, &notification); err != nil {
		return fmt.Errorf("error saving in-app notification: %w", err)
	}
	return nil
}

type PushStrategy struct {
	repository NotificationRepository
}

func NewPushStrategy(repository NotificationRepository) *PushStrategy {
	return &PushStrategy{repository: repository}
}

func (s *PushStrategy) Send(ctx context.Context, notification domain.Notification, dto dtos.NotificationMessageDTO) error {
	if err := s.repository.Save(ctx, &notification); err != nil {
		return fmt.Errorf("error saving push notification: %w", err)
	}

	// Send To Push

	return nil
}

type NotificationContext struct {
	strategies map[domain.NotificationType]NotificationStrategy
}

func NewNotificationContext(
	emailStrategy *EmailStrategy,
	smsStrategy *SMSStrategy,
	inAppStrategy *InAppStrategy,
	pushStrategy *PushStrategy,
) *NotificationContext {
	return &NotificationContext{
		strategies: map[domain.NotificationType]NotificationStrategy{
			domain.TypeEmail: emailStrategy,
			domain.TypeSMS:   smsStrategy,
			domain.TypeInApp: inAppStrategy,
			domain.TypePush:  pushStrategy,
		},
	}
}

func (c *NotificationContext) SendNotification(ctx context.Context, notification domain.Notification, dto dtos.NotificationMessageDTO) error {
	strategy, exists := c.strategies[notification.Type]
	if !exists {
		return fmt.Errorf("no strategy found for notification type: %s", notification.Type)
	}

	return strategy.Send(ctx, notification, dto)
}
