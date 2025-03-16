package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/utils"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	mongoNotification := ToMongoModel(notification)

	_, err := r.collection.InsertOne(ctx, &mongoNotification)
	if err != nil {
		return fmt.Errorf("%w: failed to insert notification", ErrDatabaseFailure)
	}

	return nil
}

func (r *NotificationRepositoryImpl) GetByID(ctx context.Context, notificationID uuid.UUID) (*domain.Notification, error) {
	var mongoNotification MongoNotification
	filter := bson.M{"_id": notificationID.String()}

	err := r.collection.FindOne(ctx, filter).Decode(&mongoNotification)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, ErrNotificationNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("%w: failed to retrieve notification", ErrDatabaseFailure)
	}

	return ToDomainModel(&mongoNotification), nil
}

func (r *NotificationRepositoryImpl) GetByUserID(ctx context.Context, userID uuid.UUID, page utils.Page) (*[]domain.Notification, int64, error) {
	filter := bson.M{"user_id": userID.String()}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, fmt.Errorf("%w: failed to count documents", ErrDatabaseFailure)
	}

	if total == 0 {
		return nil, 0, ErrNotificationNotFound
	}

	opts := getPaginationOpts(page)
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, fmt.Errorf("%w: failed to retrieve notifications", ErrDatabaseFailure)
	}
	defer cursor.Close(ctx)

	var mongoNotifications []MongoNotification
	if err := cursor.All(ctx, &mongoNotifications); err != nil {
		return nil, 0, fmt.Errorf("%w: failed to parse notifications", ErrDatabaseFailure)
	}

	domainNotifications := ToDomainModelList(&mongoNotifications)

	return domainNotifications, total, nil
}

func (r *NotificationRepositoryImpl) GetPendingNotifications(ctx context.Context, page utils.Page) (*[]domain.Notification, error) {
	filter := bson.M{"status": "PENDING"}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to retrieve pending notifications", ErrDatabaseFailure)
	}
	defer cursor.Close(ctx)

	var mongoNotificationList []MongoNotification
	if err := cursor.All(ctx, &mongoNotificationList); err != nil {
		return nil, fmt.Errorf("%w: failed to parse pending notifications", ErrDatabaseFailure)
	}

	if len(mongoNotificationList) == 0 {
		return nil, ErrNotificationNotFound
	}

	domainNotifications := ToDomainModelList(&mongoNotificationList)
	return domainNotifications, nil
}

func (r *NotificationRepositoryImpl) DeleteByID(ctx context.Context, notificationID uuid.UUID) error {
	filter := bson.M{"_id": notificationID.String()}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("%w: failed to delete notification", ErrDatabaseFailure)
	}

	if result.DeletedCount == 0 {
		return ErrNotificationNotFound
	}

	return nil
}

func getPaginationOpts(page utils.Page) *options.FindOptions {
	return options.Find().
		SetSkip(int64((page.PageNumber - 1) * page.PageSize)).
		SetLimit(int64(page.PageSize)).
		SetSort(bson.M{"created_at": -1})
}
