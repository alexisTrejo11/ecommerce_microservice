package su_service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

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
	StartSubscriptionChecker(interval time.Duration)
}

type SubscriptionServiceImpl struct {
	repo su_repository.SubscriptionRepository
}

func NewSubscriptionService(repo su_repository.SubscriptionRepository) SubscriptionService {
	return &SubscriptionServiceImpl{repo: repo}
}

func (s *SubscriptionServiceImpl) CreateSubscription(ctx context.Context, subscriptionDTO dtos.SubscriptionInsertDTO) (*dtos.SubscriptionDTO, error) {
	subscription := mapper.ToSubscription(subscriptionDTO)

	if err := s.validateCreation(ctx, subscriptionDTO); err != nil {
		return nil, err
	}

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

	if err := subscription.Cancel(); err != nil {
		return err
	}

	if err := s.repo.Save(ctx, subscription); err != nil {
		return err
	}

	return nil
}

func (s *SubscriptionServiceImpl) DeleteSubscription(ctx context.Context, subscriptionID uuid.UUID) error {
	return s.repo.SoftDelete(ctx, subscriptionID)
}

func (s *SubscriptionServiceImpl) validateCreation(ctx context.Context, dto dtos.SubscriptionInsertDTO) error {
	if err := s.validateNotSubscriptionConflict(ctx, dto.UserID); err != nil {
		return err
	}

	if dto.Type == suscription.FREE_TRIAL {
		if s.isUserAlreadyUseHisFreeTrial() {
			return errors.New("use ralready claim his free trial")
		}
	}

	return nil
}

func (s *SubscriptionServiceImpl) validateNotSubscriptionConflict(ctx context.Context, userID uuid.UUID) error {
	suscription, err := s.repo.GetValidByUserID(ctx, userID)
	if err == nil && suscription != nil {
		return errors.New("this user already has an active suscriptions")
	}

	return nil
}

// Implement
func (s *SubscriptionServiceImpl) isUserAlreadyUseHisFreeTrial() bool {
	return true
}

func (s *SubscriptionServiceImpl) StartSubscriptionChecker(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("Checking for subscriptions to expire...")
		ctx := context.Background()
		if err := s.repo.ExpireSubscriptions(ctx); err != nil {
			log.Printf("Error while expiring subscriptions: %v\n", err)
		}
	}
}
