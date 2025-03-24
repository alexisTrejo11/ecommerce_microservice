package service_test

import (
	"context"
	"testing"

	subscription "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/model"
	su_service "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/service"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
	mocks "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/tests/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateSubscription_Success(t *testing.T) {
	mockRepo := new(mocks.MockSubscriptionRepository)
	service := su_service.NewSubscriptionService(mockRepo)

	userID := uuid.New()
	subscriptionDTO := dtos.SubscriptionInsertDTO{
		UserID:   userID,
		Type:     subscription.ANUALLY,
		PlanName: "PREMUIM",
	}

	sub := subscription.NewSubscription(userID, "PREMUIM", uuid.New(), subscription.ACTIVE, subscription.ANUALLY)

	mockRepo.On("GetValidByUserID", mock.Anything, userID).Return(nil, nil)
	mockRepo.On("Save", mock.Anything, mock.MatchedBy(func(sub *subscription.Subscription) bool {
		return sub.UserID == userID && sub.SubscriptionType == subscription.ANUALLY
	})).Return(nil)

	result, err := service.CreateSubscription(context.Background(), subscriptionDTO)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, userID, result.UserID)
	assert.Equal(t, result.PlanName, sub.PlanName)
	mockRepo.AssertExpectations(t)
}

func TestGetSubscriptionByUser_Success(t *testing.T) {
	mockRepo := new(mocks.MockSubscriptionRepository)
	service := su_service.NewSubscriptionService(mockRepo)

	userID := uuid.New()
	subscription := &subscription.Subscription{
		ID:               uuid.New(),
		UserID:           userID,
		SubscriptionType: subscription.ANUALLY,
		Status:           subscription.ACTIVE,
	}

	mockRepo.On("GetByUserID", mock.Anything, userID).Return(subscription, nil)

	result, err := service.GetSubscriptionByUser(context.Background(), userID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, userID, result.UserID)
	mockRepo.AssertExpectations(t)
}

func TestUpdateSubscriptionType_Success(t *testing.T) {
	mockRepo := new(mocks.MockSubscriptionRepository)
	service := su_service.NewSubscriptionService(mockRepo)

	subscriptionID := uuid.New()
	existingSubscription := &subscription.Subscription{
		ID:               subscriptionID,
		UserID:           uuid.New(),
		SubscriptionType: subscription.FREE_TRIAL,
		Status:           subscription.ACTIVE,
	}

	newType := subscription.ANUALLY

	mockRepo.On("GetByID", mock.Anything, subscriptionID).Return(existingSubscription, nil)
	mockRepo.On("Save", mock.Anything, mock.MatchedBy(func(sub *subscription.Subscription) bool {
		return sub.GetType() == newType
	})).Return(nil)

	err := service.UpdateSubscriptionType(context.Background(), subscriptionID, newType)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCancelSubscription_Success(t *testing.T) {
	mockRepo := new(mocks.MockSubscriptionRepository)
	service := su_service.NewSubscriptionService(mockRepo)

	userID := uuid.New()
	subscriptionID := uuid.New()
	existingSubscription := &subscription.Subscription{
		ID:               subscriptionID,
		UserID:           userID,
		SubscriptionType: subscription.ANUALLY,
		Status:           subscription.ACTIVE,
	}

	mockRepo.On("GetByIdAndUserID", mock.Anything, subscriptionID, userID).Return(existingSubscription, nil)
	mockRepo.On("Save", mock.Anything, mock.MatchedBy(func(sub *subscription.Subscription) bool {
		return sub.Status == subscription.CANCELLED
	})).Return(nil)

	err := service.CancelSubscription(context.Background(), userID, subscriptionID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteSubscription_Success(t *testing.T) {
	mockRepo := new(mocks.MockSubscriptionRepository)
	service := su_service.NewSubscriptionService(mockRepo)

	subscriptionID := uuid.New()

	mockRepo.On("SoftDelete", mock.Anything, subscriptionID).Return(nil)

	err := service.DeleteSubscription(context.Background(), subscriptionID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
