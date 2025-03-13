package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentModel struct {
	ID              uuid.UUID         `gorm:"primaryKey;type:uuid"`
	CustomerID      uuid.UUID         `gorm:"not null;type:uuid"`
	Currency        string            `gorm:"not null;size:3"`
	PaymentMethodID string            `gorm:"not null;size:255"`
	Status          string            `gorm:"not null;size:50"`
	Created         time.Time         `gorm:"autoCreateTime"`
	Metadata        map[string]string `gorm:"type:jsonb"`
	TransactionID   *uuid.UUID        `gorm:"type:uuid"`
}

func (PaymentModel) TableName() string {
	return "payments"
}

func (p *PaymentModel) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

type TransactionModel struct {
	ID              uuid.UUID `gorm:"type:char(36);primaryKey"`
	PaymentID       uuid.UUID `gorm:"type:char(36);index;not null"`
	CustomerID      uuid.UUID `gorm:"type:char(36);index;not null"`
	Amount          float64   `gorm:"type:decimal(10,2);not null"`
	TransactionType string    `gorm:"type:varchar(20);not null;check:transaction_type IN ('credit', 'debit', 'refund')"`
	Reference       string    `gorm:"type:varchar(255)"`
	Description     string    `gorm:"type:text"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Payment         *PaymentModel  `gorm:"foreignKey:PaymentID"`
	Customer        *CustomerModel `gorm:"foreignKey:CustomerID"`
}

func (TransactionModel) TableName() string {
	return "transactions"
}

func (t *TransactionModel) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}

type InvoiceModel struct {
	ID             uuid.UUID  `gorm:"primaryKey;size:36"`
	SubscriptionID string     `gorm:"not null;size:36"`
	AmountDue      int64      `gorm:"not null"`
	Currency       string     `gorm:"not null;size:3"`
	Status         string     `gorm:"not null"`
	IssuedAt       time.Time  `gorm:"autoCreateTime"`
	PaidAt         *time.Time `gorm:"default:null"`
	CreatedAt      time.Time  `gorm:"autoCreateTime"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime"`
}

func (InvoiceModel) TableName() string {
	return "invoices"
}

func (i *InvoiceModel) BeforeCreate(tx *gorm.DB) error {
	if i.ID == uuid.Nil {
		i.ID = uuid.New()
	}
	return nil
}

type CustomerModel struct {
	ID               uuid.UUID          `gorm:"primaryKey;size:36"`
	Email            string             `gorm:"unique;not null"`
	Name             string             `gorm:"not null"`
	PaymentMethods   []PaymentMethodRef `gorm:"foreignKey:CustomerID"`
	DefaultPaymentID *string            `gorm:"size:36"`
	CreatedAt        time.Time          `gorm:"autoCreateTime"`
	UpdatedAt        time.Time          `gorm:"autoUpdateTime"`
}

type PaymentMethodRef struct {
	ID         string    `gorm:"primaryKey;size:36"`
	CustomerID string    `gorm:"size:36;not null"`
	Method     string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

func (CustomerModel) TableName() string {
	return "Customers"
}

type CourseModel struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"type:decimal(10,2);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (CourseModel) TableName() string {
	return "courses"
}

type PaymentCourseModel struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	PaymentID uuid.UUID `gorm:"type:char(36);index;not null"`
	CourseID  uuid.UUID `gorm:"type:char(36);index;not null"`
	CreatedAt time.Time
	Payment   PaymentModel `gorm:"foreignKey:PaymentID"`
	Course    CourseModel  `gorm:"foreignKey:CourseID"`
}

func (PaymentCourseModel) TableName() string {
	return "payment_courses"
}

func (pc *PaymentCourseModel) BeforeCreate(tx *gorm.DB) error {
	if pc.ID == uuid.Nil {
		pc.ID = uuid.New()
	}
	return nil
}
