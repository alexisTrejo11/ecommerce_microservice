package notification

import (
	"context"
	"fmt"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/dtos"
)

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
