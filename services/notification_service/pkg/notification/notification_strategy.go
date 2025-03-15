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
