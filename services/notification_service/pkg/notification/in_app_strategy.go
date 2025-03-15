package notification

import (
	"context"
	"fmt"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/dtos"
)

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
