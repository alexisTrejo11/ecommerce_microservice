package domain

import (
	"time"

	"github.com/google/uuid"
)

type TransactionType string

const (
	PaymentTransaction TransactionType = "PAYMENT"
	RefundTransaction  TransactionType = "REFUND"
)

type Transaction struct {
	ID              uuid.UUID
	PaymentID       uuid.UUID
	UserID          uuid.UUID
	Amount          float64
	TransactionType TransactionType
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewTransaction(paymentID uuid.UUID, userID uuid.UUID, amount float64, transactionType TransactionType) *Transaction {
	return &Transaction{
		ID:              uuid.New(),
		PaymentID:       paymentID,
		UserID:          userID,
		Amount:          amount,
		TransactionType: transactionType,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}
