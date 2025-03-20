package su_repository

import (
	"context"
	"fmt"
	"time"

	suscription "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/suscription/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	GetByID(ctx context.Context, subscriptionID uuid.UUID) (*suscription.Subscription, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*suscription.Subscription, error)
	GetByIdAndUserID(ctx context.Context, subscriptionID, userID uuid.UUID) (*suscription.Subscription, error)
	GetValidByUserID(ctx context.Context, userID uuid.UUID) (*suscription.Subscription, error)
	Save(ctx context.Context, subscription *suscription.Subscription) error
	SoftDelete(ctx context.Context, subscriptionID uuid.UUID) error
	ExpireSubscriptions(ctx context.Context) error
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

	if err := r.db.WithContext(ctx).
		Where("user_id = ? AND status != ?", userID, "EXPIRED").
		First(&subscription).Error; err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (r *SubscriptionRepositoryImpl) GetValidByUserID(ctx context.Context, userID uuid.UUID) (*suscription.Subscription, error) {
	var subscription suscription.Subscription
	if err := r.db.WithContext(ctx).Where("user_id = ? and status = 'ACTIVE'", userID).First(&subscription).Error; err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (r *SubscriptionRepositoryImpl) Save(ctx context.Context, subscription *suscription.Subscription) error {
	return r.db.WithContext(ctx).Save(subscription).Error
}

func (r *SubscriptionRepositoryImpl) SoftDelete(ctx context.Context, subscriptionID uuid.UUID) error {
	var subscription suscription.Subscription

	if err := r.db.WithContext(ctx).Where("id = ?", subscriptionID).First(&subscription).Error; err != nil {
		return err
	}

	if err := r.db.WithContext(ctx).Delete(&subscription).Error; err != nil {
		return err
	}

	return nil
}

func (r *SubscriptionRepositoryImpl) ExpireSubscriptions(ctx context.Context) error {
	now := time.Now()

	result := r.db.WithContext(ctx).
		Model(&suscription.Subscription{}).
		Where("end_date <= ? AND status != ?", now, "EXPIRED").
		Update("status", "EXPIRED")

	if result.Error != nil {
		return result.Error
	}

	fmt.Printf("%d suscriptions have been expired.\n", result.RowsAffected)
	return nil
}
