package dtos

import (
	"time"

	subscription "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/model"
	"github.com/google/uuid"
)

type SubscriptionDTO struct {
	ID        uuid.UUID                       `json:"id"`
	UserID    uuid.UUID                       `json:"user_id"`
	PlanName  string                          `json:"plan_name"`
	StartDate time.Time                       `json:"start_date"`
	EndDate   time.Time                       `json:"end_date"`
	Status    subscription.SubscriptionStatus `json:"status"`
	Type      subscription.SubscriptionType   `json:"type"`
	PaymentID uuid.UUID                       `json:"payment_id"`
}

type SubscriptionInsertDTO struct {
	UserID    uuid.UUID                       `json:"user_id" validate:"required,uuid"`
	PlanName  string                          `json:"plan_name" validate:"required,min=3,max=50"`
	Status    subscription.SubscriptionStatus `json:"status" validate:"required"`
	Type      subscription.SubscriptionType   `json:"type" validate:"required"`
	PaymentID uuid.UUID                       `json:"payment_id" validate:"required,uuid"`
}
