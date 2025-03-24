package su_repository

import (
	"context"
	"fmt"
	"time"

	subscription "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	GetByID(ctx context.Context, subscriptionID uuid.UUID) (*subscription.Subscription, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*subscription.Subscription, error)
	GetByIdAndUserID(ctx context.Context, subscriptionID, userID uuid.UUID) (*subscription.Subscription, error)
	GetValidByUserID(ctx context.Context, userID uuid.UUID) (*subscription.Subscription, error)
	Save(ctx context.Context, subscription *subscription.Subscription) error
	SoftDelete(ctx context.Context, subscriptionID uuid.UUID) error
	ExpireSubscriptions(ctx context.Context) error
}

type SubscriptionRepositoryImpl struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &SubscriptionRepositoryImpl{db: db}
}

func (r *SubscriptionRepositoryImpl) GetByID(ctx context.Context, subscriptionID uuid.UUID) (*subscription.Subscription, error) {
	var subscription subscription.Subscription
	if err := r.db.WithContext(ctx).Where("id = ?", subscriptionID).First(&subscription).Error; err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (r *SubscriptionRepositoryImpl) GetByIdAndUserID(ctx context.Context, subscriptionID, userID uuid.UUID) (*subscription.Subscription, error) {
	var subscription subscription.Subscription
	if err := r.db.WithContext(ctx).Where("id = ? AND user_id", subscriptionID, userID).First(&subscription).Error; err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (r *SubscriptionRepositoryImpl) GetByUserID(ctx context.Context, userID uuid.UUID) (*subscription.Subscription, error) {
	var subscription subscription.Subscription

	if err := r.db.WithContext(ctx).
		Where("user_id = ? AND status != ?", userID, "EXPIRED").
		First(&subscription).Error; err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (r *SubscriptionRepositoryImpl) GetValidByUserID(ctx context.Context, userID uuid.UUID) (*subscription.Subscription, error) {
	var subscription subscription.Subscription
	if err := r.db.WithContext(ctx).Where("user_id = ? and status = 'ACTIVE'", userID).First(&subscription).Error; err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (r *SubscriptionRepositoryImpl) Save(ctx context.Context, subscription *subscription.Subscription) error {
	return r.db.WithContext(ctx).Save(subscription).Error
}

// Test Issues
func (r *SubscriptionRepositoryImpl) SoftDelete(ctx context.Context, subscriptionID uuid.UUID) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var subscription subscription.Subscription
	if err := tx.Where("id = ?", subscriptionID).First(&subscription).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&subscription).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *SubscriptionRepositoryImpl) ExpireSubscriptions(ctx context.Context) error {
	now := time.Now()

	result := r.db.WithContext(ctx).
		Model(&subscription.Subscription{}).
		Where("end_date <= ? AND status != ?", now, "EXPIRED").
		Update("status", "EXPIRED")

	if result.Error != nil {
		return result.Error
	}

	fmt.Printf("%d subscriptions have been expired.\n", result.RowsAffected)
	return nil
}
