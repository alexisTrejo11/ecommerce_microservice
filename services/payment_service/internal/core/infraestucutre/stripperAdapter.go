package infrastructure

import (
	"context"
	"fmt"

	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/entity"
	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/valueobject"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

type StripeAdapter struct {
	apiKey string
}

func NewStripeAdapter(apiKey string) *StripeAdapter {
	stripe.Key = apiKey
	return &StripeAdapter{apiKey: apiKey}
}

func (s *StripeAdapter) CreatePaymentIntent(
	ctx context.Context,
	amount int64,
	currency string,
	customerID string,
	paymentMethodID string,
) (string, error) {
	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(amount),
		Currency:      stripe.String(currency),
		Customer:      stripe.String(customerID),
		PaymentMethod: stripe.String(paymentMethodID),
		Confirm:       stripe.Bool(false),
	}

	intent, err := paymentintent.New(params)
	if err != nil {
		return "", fmt.Errorf("failed to create payment intent in Stripe: %w", err)
	}

	return intent.ID, nil
}

func (s *StripeAdapter) ConfirmPaymentIntent(ctx context.Context, paymentIntentID string) error {
	params := &stripe.PaymentIntentConfirmParams{}
	_, err := paymentintent.Confirm(paymentIntentID, params)
	if err != nil {
		return fmt.Errorf("failed to confirm payment intent in Stripe: %w", err)
	}

	return nil
}

func (s *StripeAdapter) GetPaymentIntent(ctx context.Context, paymentIntentID string) (*entity.PaymentIntent, error) {
	// Obtener el PaymentIntent desde Stripe
	intent, err := paymentintent.Get(paymentIntentID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch payment intent from Stripe: %w", err)
	}

	amount, err := valueobject.NewMoney(float64(intent.Amount), valueobject.Currency(intent.Currency))
	if err != nil {
		return nil, fmt.Errorf("failed to create Money object: %w", err)
	}

	status := mapStripeStatusToEntity(intent.Status)

	courseID := ""
	if val, exists := intent.Metadata["course_id"]; exists {
		courseID = val
	}

	paymentIntent, err := entity.NewPaymentIntent(
		intent.ID,
		intent.Customer.ID,
		amount,
		intent.PaymentMethod.ID,
		courseID,
		status,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create PaymentIntent: %w", err)
	}

	return paymentIntent, nil
}
func mapStripeStatusToEntity(stripeStatus stripe.PaymentIntentStatus) entity.PaymentIntentStatus {
	switch stripeStatus {
	case stripe.PaymentIntentStatusRequiresPaymentMethod:
		return entity.RequiresPaymentMethod
	case stripe.PaymentIntentStatusProcessing:
		return entity.Processing
	case stripe.PaymentIntentStatusSucceeded:
		return entity.Succeeded
	case stripe.PaymentIntentStatusCanceled:
		return entity.Canceled
	default:
		return entity.RequiresPaymentMethod
	}
}
