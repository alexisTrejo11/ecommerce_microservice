package input

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/dtos"
)

type EmailUseCase interface {
	SendEmail(ctx context.Context, emailDTO dtos.NotificationMessageDTO) error
}
