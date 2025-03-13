package entity

import (
	"errors"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/valueobject"
)

type InvoiceStatus string

const (
	Draft InvoiceStatus = "draft"
	Open  InvoiceStatus = "open"
	Paid  InvoiceStatus = "paid"
	Void  InvoiceStatus = "void"
)

type Invoice struct {
	id             string
	subscriptionID string
	amountDue      valueobject.Money
	status         InvoiceStatus
	issuedAt       time.Time
	paidAt         *time.Time
}

func NewInvoice(id, subscriptionID string, amountDue valueobject.Money) (*Invoice, error) {
	if id == "" || subscriptionID == "" {
		return nil, errors.New("id and subscriptionID cannot be empty")
	}
	if amountDue.Amount <= 0 {
		return nil, errors.New("amountDue must be greater than zero")
	}

	return &Invoice{
		id:             id,
		subscriptionID: subscriptionID,
		amountDue:      amountDue,
		status:         Draft,
		issuedAt:       time.Now(),
	}, nil
}

func (i *Invoice) Issue() error {
	if i.status != Draft {
		return errors.New("invoice can only be issued from draft status")
	}
	i.status = Open
	return nil
}

func (i *Invoice) MarkAsPaid() error {
	if i.status != Open {
		return errors.New("invoice can only be marked as paid if it is open")
	}
	now := time.Now()
	i.paidAt = &now
	i.status = Paid
	return nil
}

func (i *Invoice) Void() error {
	if i.status == Paid {
		return errors.New("cannot void a paid invoice")
	}
	i.status = Void
	return nil
}

func (i *Invoice) ID() string                   { return i.id }
func (i *Invoice) SubscriptionID() string       { return i.subscriptionID }
func (i *Invoice) AmountDue() valueobject.Money { return i.amountDue }
func (i *Invoice) Status() InvoiceStatus        { return i.status }
func (i *Invoice) IssuedAt() time.Time          { return i.issuedAt }
func (i *Invoice) PaidAt() *time.Time           { return i.paidAt }
