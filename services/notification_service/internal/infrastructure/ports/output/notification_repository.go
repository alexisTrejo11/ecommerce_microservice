package repository

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationRepositoryImpl struct {
	collection *mongo.Collection
}

func NewNotificationRepository(mongoClient *mongo.Client) output.NotificationRepository {
	collection := mongoClient.Database("notifications").Collection("notifications")

	return &NotificationRepositoryImpl{
		collection: collection,
	}
}

func (r *NotificationRepositoryImpl) Save(ctx context.Context, notification *domain.Notification) error {
	_, err := r.collection.InsertOne(ctx, notification)
	return err
}

func (r *NotificationRepositoryImpl) GetByID(ctx context.Context, notificationID string) (*domain.Notification, error) {
	return nil, nil
}

func (r *NotificationRepositoryImpl) GetByUserID(ctx context.Context, userID string, limit, offset int) ([]*domain.Notification, int64, error) {
	return nil, 0, nil
}

func (r *NotificationRepositoryImpl) GetPendingNotifications(ctx context.Context, limit int) ([]*domain.Notification, error) {
	return nil, nil
}
func (r *NotificationRepositoryImpl) DeleteByID(ctx context.Context, notificationID string) error {
	return nil
}
