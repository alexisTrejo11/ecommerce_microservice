package su_service

import (
	"context"
	"fmt"

	suscription "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/suscription/model"
	su_repository "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/suscription/repository"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
	mapper "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/mappers"
	"github.com/google/uuid"
)

type SubscriptionService interface {
	CreateSubscription(ctx context.Context, subscriptionDTO dtos.SubscriptionInsertDTO) (*dtos.SubscriptionDTO, error)
	GetSubscriptionByUser(ctx context.Context, userID uuid.UUID) (*dtos.SubscriptionDTO, error)
	UpdateSubscriptionType(ctx context.Context, subscriptionID uuid.UUID, subType suscription.SubscriptionType) error
	CancelSubscription(ctx context.Context, userID, subscriptionID uuid.UUID) error
	DeleteSubscription(ctx context.Context, subscriptionID uuid.UUID) error
}

type SubscriptionServiceImpl struct {
	repo su_repository.SubscriptionRepository
}

func NewSubscriptionService(repo su_repository.SubscriptionRepository) SubscriptionService {
	return &SubscriptionServiceImpl{repo: repo}
}

func (s *SubscriptionServiceImpl) CreateSubscription(ctx context.Context, subscriptionDTO dtos.SubscriptionInsertDTO) (*dtos.SubscriptionDTO, error) {
	subscription := mapper.ToSubscription(subscriptionDTO)

	if err := s.repo.Save(ctx, &subscription); err != nil {
		return nil, err
	}

	createdDTO := mapper.ToSubscriptionDTO(subscription)
	return &createdDTO, nil
}

func (s *SubscriptionServiceImpl) GetSubscriptionByUser(ctx context.Context, userID uuid.UUID) (*dtos.SubscriptionDTO, error) {
	subscription, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	subscriptionDTO := mapper.ToSubscriptionDTO(*subscription)
	return &subscriptionDTO, nil
}

func (s *SubscriptionServiceImpl) UpdateSubscriptionType(ctx context.Context, subscriptionID uuid.UUID, subType suscription.SubscriptionType) error {
	subscription, err := s.repo.GetByID(ctx, subscriptionID)
	if err != nil {
		return nil
	}

	subscription.SetType(&subType)
	if err := s.repo.Save(ctx, subscription); err != nil {
		return err
	}

	return nil
}

func (s *SubscriptionServiceImpl) CancelSubscription(ctx context.Context, userID, subscriptionID uuid.UUID) error {
	subscription, err := s.repo.GetByIdAndUserID(ctx, subscriptionID, userID)
	if err != nil {
		return err
	}

	subscription.Cancel()

	if err := s.repo.Save(ctx, subscription); err != nil {
		return err
	}

	return nil
}

func (s *SubscriptionServiceImpl) DeleteSubscription(ctx context.Context, subscriptionID uuid.UUID) error {
	_, err := s.repo.GetByID(ctx, subscriptionID)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	return s.repo.Delete(ctx, subscriptionID)
}
