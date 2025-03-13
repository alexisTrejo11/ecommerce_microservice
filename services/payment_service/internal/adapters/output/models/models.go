package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentModel struct {
	ID              uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID          uuid.UUID `gorm:"type:char(36);index;not null"`
	Amount          float64   `gorm:"type:decimal(10,2);not null"`
	Status          string    `gorm:"type:varchar(20);default:'pending';check:status IN ('pending', 'completed', 'failed')"`
	StripePaymentID string    `gorm:"type:varchar(255)"`
	InvoiceID       uuid.UUID `gorm:"type:char(36);index"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt     `gorm:"index"`
	User            UserModel          `gorm:"foreignKey:UserID"`
	Invoice         InvoiceModel       `gorm:"foreignKey:InvoiceID"`
	Transactions    []TransactionModel `gorm:"foreignKey:PaymentID;constraint:OnDelete:CASCADE"`
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
	UserID          uuid.UUID `gorm:"type:char(36);index;not null"`
	Amount          float64   `gorm:"type:decimal(10,2);not null"`
	TransactionType string    `gorm:"type:varchar(20);not null;check:transaction_type IN ('credit', 'debit', 'refund')"`
	Reference       string    `gorm:"type:varchar(255)"`
	Description     string    `gorm:"type:text"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Payment         *PaymentModel  `gorm:"foreignKey:PaymentID"`
	User            *UserModel     `gorm:"foreignKey:UserID"`
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
	ID            uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID        uuid.UUID `gorm:"type:char(36);index;not null"`
	TotalAmount   float64   `gorm:"type:decimal(10,2);not null"`
	Status        string    `gorm:"type:varchar(20);default:'pending';check:status IN ('pending', 'paid', 'cancelled')"`
	InvoiceNumber string    `gorm:"type:varchar(50);unique"`
	DueDate       time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	User          UserModel      `gorm:"foreignKey:UserID"`
	Payments      []PaymentModel `gorm:"foreignKey:InvoiceID;constraint:OnDelete:SET NULL"`
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

type UserModel struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	Email     string    `gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (UserModel) TableName() string {
	return "users"
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
