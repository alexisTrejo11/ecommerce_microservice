package su_repository

import (
	"context"

	suscription "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/suscription/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	GetByID(ctx context.Context, subscriptionID uuid.UUID) (*suscription.Subscription, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*suscription.Subscription, error)
	GetByIdAndUserID(ctx context.Context, subscriptionID, userID uuid.UUID) (*suscription.Subscription, error)
	Save(ctx context.Context, subscription *suscription.Subscription) error
	Delete(ctx context.Context, subscriptionID uuid.UUID) error
}

type SubscriptionRepositoryImpl struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &SubscriptionRepositoryImpl{db: db}
}

func (r *SubscriptionRepositoryImpl) GetByID(ctx context.Context, subscriptionID uuid.UUID) (*suscription.Subscription, error) {
	var subscription suscription.Subscription
	if err := r.db.WithContext(ctx).Where("id = ?", subscriptionID).First(&subscription).Error; err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (r *SubscriptionRepositoryImpl) GetByIdAndUserID(ctx context.Context, subscriptionID, userID uuid.UUID) (*suscription.Subscription, error) {
	var subscription suscription.Subscription
	if err := r.db.WithContext(ctx).Where("id = ? AND user_id", subscriptionID, userID).First(&subscription).Error; err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (r *SubscriptionRepositoryImpl) GetByUserID(ctx context.Context, userID uuid.UUID) (*suscription.Subscription, error) {
	var subscription suscription.Subscription
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&subscription).Error; err != nil {
		return nil, err
	}
	return &subscription, nil
}

func (r *SubscriptionRepositoryImpl) Save(ctx context.Context, subscription *suscription.Subscription) error {
	return r.db.WithContext(ctx).Save(subscription).Error
}

func (r *SubscriptionRepositoryImpl) Delete(ctx context.Context, subscriptionID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&suscription.Subscription{}, subscriptionID).Error
}
