package dtos

import (
	"time"

	suscription "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/suscription/model"
	"github.com/google/uuid"
)

type SubscriptionDTO struct {
	ID        uuid.UUID                     `json:"id"`
	UserID    uuid.UUID                     `json:"user_id"`
	PlanName  string                        `json:"plan_name"`
	StartDate time.Time                     `json:"start_date"`
	EndDate   time.Time                     `json:"end_date"`
	Status    suscription.SuscriptionStatus `json:"status"`
	Type      suscription.SubscriptionType  `json:"type"`
	PaymentID uuid.UUID                     `json:"payment_id"`
}

type SubscriptionInsertDTO struct {
	UserID    uuid.UUID                     `json:"user_id" validate:"required,uuid"`
	PlanName  string                        `json:"plan_name" validate:"required,min=3,max=50"`
	Status    suscription.SuscriptionStatus `json:"status" validate:"required"`
	Type      suscription.SubscriptionType  `json:"type" validate:"required"`
	PaymentID uuid.UUID                     `json:"payment_id" validate:"required,uuid"`
}
