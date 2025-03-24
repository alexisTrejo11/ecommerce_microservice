package mocks

import (
	"context"

	subscription "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockSubscriptionRepository struct {
	mock.Mock
}

func (m *MockSubscriptionRepository) Save(ctx context.Context, subscription *subscription.Subscription) error {
	args := m.Called(ctx, subscription)
	return args.Error(0)
}

func (m *MockSubscriptionRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*subscription.Subscription, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*subscription.Subscription), args.Error(1)
}

func (m *MockSubscriptionRepository) GetByIdAndUserID(ctx context.Context, subscriptionID, userID uuid.UUID) (*subscription.Subscription, error) {
	args := m.Called(ctx, subscriptionID, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*subscription.Subscription), args.Error(1)
}

func (m *MockSubscriptionRepository) GetByID(ctx context.Context, subscriptionID uuid.UUID) (*subscription.Subscription, error) {
	args := m.Called(ctx, subscriptionID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*subscription.Subscription), args.Error(1)
}

func (m *MockSubscriptionRepository) SoftDelete(ctx context.Context, subscriptionID uuid.UUID) error {
	args := m.Called(ctx, subscriptionID)
	return args.Error(0)
}

func (m *MockSubscriptionRepository) ExpireSubscriptions(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockSubscriptionRepository) GetValidByUserID(ctx context.Context, userID uuid.UUID) (*subscription.Subscription, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*subscription.Subscription), args.Error(1)
}
