package domain

import (
	"time"

	"github.com/google/uuid"
)

type PaymentStatus string

const (
	Pending   PaymentStatus = "PENDING"
	Completed PaymentStatus = "COMPLETED"
	Failed    PaymentStatus = "FAILED"
)

type Payment struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Amount    float64
	Status    PaymentStatus
	PaymentID string // ID Stripe
	CreatedAt time.Time
	UpdatedAt time.Time
	InvoiceID uuid.UUID
}

func NewPayment(userID uuid.UUID, amount float64) *Payment {
	return &Payment{
		ID:        uuid.New(),
		UserID:    userID,
		Amount:    amount,
		Status:    Pending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (p *Payment) Complete(paymentID string) {
	p.Status = Completed
	p.PaymentID = paymentID
	p.UpdatedAt = time.Now()
}

func (p *Payment) Fail() {
	p.Status = Failed
	p.UpdatedAt = time.Now()
}
