package dtos

import (
	"time"

	subscription "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/model"
	"github.com/google/uuid"
)

// SubscriptionDTO defines the structure for representing a subscription.
// swagger:model SubscriptionDTO
type SubscriptionDTO struct {
	// The unique identifier of the subscription.
	// Example: 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `json:"id"`

	// The ID of the user associated with the subscription.
	// Example: 550e8400-e29b-41d4-a716-446655440001
	UserID uuid.UUID `json:"user_id"`

	// The name of the subscription plan.
	// Example: Premium Monthly
	PlanName string `json:"plan_name"`

	// The start date of the subscription.
	// Format: date-time
	// Example: 2023-10-01T00:00:00Z
	StartDate time.Time `json:"start_date"`

	// The end date of the subscription.
	// Format: date-time
	// Example: 2023-11-01T00:00:00Z
	EndDate time.Time `json:"end_date"`

	// The status of the subscription.
	// Enum: active, inactive, cancelled
	// Example: active
	Status subscription.SubscriptionStatus `json:"status"`

	// The type of the subscription.
	// Enum: monthly, yearly
	// Example: monthly
	Type subscription.SubscriptionType `json:"type"`

	// The ID of the payment associated with the subscription.
	// Example: 550e8400-e29b-41d4-a716-446655440002
	PaymentID uuid.UUID `json:"payment_id"`
}

// SubscriptionInsertDTO defines the structure for creating a new subscription.
// swagger:model SubscriptionInsertDTO
type SubscriptionInsertDTO struct {
	// The ID of the user associated with the subscription.
	// Required: true
	// Example: 550e8400-e29b-41d4-a716-446655440001
	UserID uuid.UUID `json:"user_id" validate:"required,uuid"`

	// The name of the subscription plan.
	// Required: true
	// Minimum length: 3
	// Maximum length: 50
	// Example: Premium Monthly
	PlanName string `json:"plan_name" validate:"required,min=3,max=50"`

	// The status of the subscription.
	// Required: true
	// Enum: active, inactive, cancelled
	// Example: active
	Status subscription.SubscriptionStatus `json:"status" validate:"required"`

	// The type of the subscription.
	// Required: true
	// Enum: monthly, yearly
	// Example: monthly
	Type subscription.SubscriptionType `json:"type" validate:"required"`

	// The ID of the payment associated with the subscription.
	// Required: true
	// Example: 550e8400-e29b-41d4-a716-446655440002
	PaymentID uuid.UUID `json:"payment_id" validate:"required,uuid"`
}
