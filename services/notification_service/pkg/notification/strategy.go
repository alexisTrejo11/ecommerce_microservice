package notification

import (
	"context"
	"fmt"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/dtos"
)

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
