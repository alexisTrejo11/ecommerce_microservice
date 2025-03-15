package notification

import (
	"context"
	"fmt"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/dtos"
)

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
