package entity

import (
	"errors"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/valueobject"
)

type PaymentIntentStatus string

const (
	RequiresPaymentMethod PaymentIntentStatus = "requires_payment_method"
	Processing            PaymentIntentStatus = "processing"
	Succeeded             PaymentIntentStatus = "succeeded"
	Canceled              PaymentIntentStatus = "canceled"
)

type PaymentIntent struct {
	id              string
	customerID      string
	amount          valueobject.Money
	paymentMethodID string
	status          PaymentIntentStatus
	created         time.Time
	metadata        map[string]string
	transaction     Transaction
	courseID        string
}

func NewPaymentIntent(
	id string,
	customerID string,
	amount valueobject.Money,
	paymentMethodID string,
	courseID string,
) (*PaymentIntent, error) {
	if id == "" || customerID == "" || paymentMethodID == "" || courseID == "" {
		return nil, errors.New("id, customerID, paymentMethodID, and courseID cannot be empty")
	}
	if amount.Amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	return &PaymentIntent{
		id:              id,
		customerID:      customerID,
		amount:          amount,
		paymentMethodID: paymentMethodID,
		status:          RequiresPaymentMethod,
		created:         time.Now(),
		metadata:        make(map[string]string),
		transaction:     Transaction{},
		courseID:        courseID,
	}, nil
}

func (pi *PaymentIntent) CourseID() string { return pi.courseID }

func (pi *PaymentIntent) Confirm() error {
	if pi.status != RequiresPaymentMethod {
		return errors.New("invalid status: only 'requires_payment_method' can be confirmed")
	}
	pi.status = Processing
	return nil
}

func (pi *PaymentIntent) MarkAsSucceeded() error {
	if pi.status != Processing {
		return errors.New("invalid status: only 'processing' can be marked as succeeded")
	}
	pi.status = Succeeded
	return nil
}

func (pi *PaymentIntent) MarkAsCanceled() error {
	if pi.status == Succeeded {
		return errors.New("cannot cancel a succeeded payment intent")
	}
	pi.status = Canceled
	return nil
}
