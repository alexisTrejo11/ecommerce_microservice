package su_service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	subscription "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/model"
	su_repository "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/repository"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
	mapper "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/mappers"
	"github.com/google/uuid"
)

type SubscriptionService interface {
	CreateSubscription(ctx context.Context, subscriptionDTO dtos.SubscriptionInsertDTO) (*dtos.SubscriptionDTO, error)
	GetSubscriptionByUser(ctx context.Context, userID uuid.UUID) (*dtos.SubscriptionDTO, error)
	UpdateSubscriptionType(ctx context.Context, subscriptionID uuid.UUID, subType subscription.SubscriptionType) error
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
	if err := s.validateCreation(ctx, subscriptionDTO); err != nil {
		return nil, err
	}

	subscription := mapper.ToSubscription(subscriptionDTO)
	if err := s.repo.Save(ctx, &subscription); err != nil {
		fmt.Println("Hola")
		return nil, err
	}

	fmt.Println("Hol2")
	createdDTO := mapper.ToSubscriptionDTO(subscription)
	fmt.Println(createdDTO.ID)
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

func (s *SubscriptionServiceImpl) UpdateSubscriptionType(ctx context.Context, subscriptionID uuid.UUID, subType subscription.SubscriptionType) error {
	subscription, err := s.repo.GetByID(ctx, subscriptionID)
	if err != nil {
		return err
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

	if dto.Type != subscription.FREE_TRIAL {
		return nil
	}

	if s.isUserAlreadyUseHisFreeTrial() {
		return errors.New("user already claimed their free trial")
	}

	return nil
}

func (s *SubscriptionServiceImpl) validateNotSubscriptionConflict(ctx context.Context, userID uuid.UUID) error {
	subscription, err := s.repo.GetValidByUserID(ctx, userID)
	if err == nil && subscription != nil {
		return errors.New("this user already has an active subscriptions")
	}

	return nil
}

// Implement
func (s *SubscriptionServiceImpl) isUserAlreadyUseHisFreeTrial() bool {
	if s.repo == nil {
		log.Println("Repository is nil in isUserAlreadyUseHisFreeTrial")
		return false
	}

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
