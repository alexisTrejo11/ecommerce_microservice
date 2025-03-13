package entity

import (
	"time"

	"github.com/google/uuid"
)

type TransactionStatus string

const (
	PendingTransaction   TransactionStatus = "PENDING"
	SucceededTransaction TransactionStatus = "SUCCEDED"
	FailedTransaction    TransactionStatus = "FAILED"
	RefundedTransaction  TransactionStatus = "REFUNDED"
	CanceledTransaction  TransactionStatus = "CANCELED"
)

type TransactionType string

const (
	PaymentTransaction TransactionType = "PAYMENT"
	RefundTransaction  TransactionType = "REFUND"
)

type Transaction struct {
	ID                 uuid.UUID
	PaymentID          uuid.UUID
	UserID             uuid.UUID
	Amount             float64
	TransactionType    TransactionType
	Status             TransactionStatus
	Reference          string // ID Stripe
	Description        string
	Fee                float64
	NetAmount          float64
	StripeBalanceTxnID string
	CreatedAt          time.Time
	UpdatedAt          time.Time
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
