package mapper

import (
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/dtos"
)

func ToNotificationDTO(n *domain.Notification) *dtos.NotificationDTO {
	return &dtos.NotificationDTO{
		ID:          n.ID,
		UserID:      n.UserID,
		Type:        string(n.Type),
		Title:       n.Title,
		Content:     n.Content,
		Metadata:    n.Metadata,
		Status:      string(n.Status),
		CreatedAt:   n.CreatedAt,
		UpdatedAt:   n.UpdatedAt,
		SentAt:      n.SentAt,
		ScheduledAt: n.ScheduledAt,
	}
}

func ToNotificationDTOList(notifications []*domain.Notification) []*dtos.NotificationDTO {
	dtosList := make([]*dtos.NotificationDTO, 0, len(notifications))
	for _, n := range notifications {
		dtosList = append(dtosList, ToNotificationDTO(n))
	}
	return dtosList
}

func ToNotificationMessageDTO(n *domain.Notification) *dtos.NotificationMessageDTO {
	return &dtos.NotificationMessageDTO{
		ID:        n.ID,
		UserID:    n.UserID,
		Type:      string(n.Type),
		Title:     n.Title,
		Content:   n.Content,
		Metadata:  n.Metadata,
		CreatedAt: n.CreatedAt,
	}
}

func ToDomainFromMessageDTO(dto *dtos.NotificationMessageDTO) *domain.Notification {
	return &domain.Notification{
		ID:        dto.ID,
		UserID:    dto.UserID,
		Type:      domain.NotificationType(dto.Type),
		Title:     dto.Title,
		Content:   dto.Content,
		Metadata:  dto.Metadata,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: time.Now(),
		Status:    domain.StatusPending,
	}
}
