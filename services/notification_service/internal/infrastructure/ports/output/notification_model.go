package repository

import (
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/domain"
)

type MongoNotification struct {
	ID          string            `bson:"_id"`
	UserID      string            `bson:"user_id"`
	Type        string            `bson:"type"`
	Title       string            `bson:"title"`
	Content     string            `bson:"content"`
	Metadata    map[string]string `bson:"metadata"`
	Status      string            `bson:"status"`
	CreatedAt   time.Time         `bson:"created_at"`
	UpdatedAt   time.Time         `bson:"updated_at"`
	SentAt      *time.Time        `bson:"sent_at,omitempty"`
	ScheduledAt *time.Time        `bson:"scheduled_at,omitempty"`
}

func ToMongoModel(notification *domain.Notification) *MongoNotification {
	return &MongoNotification{
		ID:          notification.ID,
		UserID:      notification.UserID,
		Type:        string(notification.Type),
		Title:       notification.Title,
		Content:     notification.Content,
		Metadata:    notification.Metadata,
		Status:      string(notification.Status),
		CreatedAt:   notification.CreatedAt,
		UpdatedAt:   notification.UpdatedAt,
		SentAt:      notification.SentAt,
		ScheduledAt: notification.ScheduledAt,
	}
}

func ToDomainModel(mongo *MongoNotification) *domain.Notification {
	return &domain.Notification{
		ID:          mongo.ID,
		UserID:      mongo.UserID,
		Type:        domain.NotificationType(mongo.Type),
		Title:       mongo.Title,
		Content:     mongo.Content,
		Metadata:    mongo.Metadata,
		Status:      domain.NotificationStatus(mongo.Status),
		CreatedAt:   mongo.CreatedAt,
		UpdatedAt:   mongo.UpdatedAt,
		SentAt:      mongo.SentAt,
		ScheduledAt: mongo.ScheduledAt,
	}
}

func ToDomainModelList(mongoList *[]MongoNotification) *[]domain.Notification {
	notificationList := make([]domain.Notification, 0, len(*mongoList))
	for _, mongo := range *mongoList {
		notification := ToDomainModel(&mongo)
		notificationList = append(notificationList, *notification)
	}

	return &notificationList
}
