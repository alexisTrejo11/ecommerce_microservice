package suscription

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SuscriptionStatus string
type SubscriptionType string

var (
	ACTIVE          SuscriptionStatus = "ACTIVE"
	CANCELLED       SuscriptionStatus = "CANCELLED"
	EXPIRED         SuscriptionStatus = "EXPIRED"
	INVALID_PAYMENT SuscriptionStatus = "INVALID_PAYMENT"
)

var (
	MONTLY       SubscriptionType = "MOTHLY"
	FREE_TRIAL   SubscriptionType = "FREE_TRIAL"
	ANUALLY      SubscriptionType = "ANUALLY"
	THREE_MONTHS SubscriptionType = "THREE_MONTHS"
)

func (s SubscriptionType) IsValid() bool {
	switch s {
	case MONTLY, FREE_TRIAL, ANUALLY, THREE_MONTHS:
		return true
	default:
		return false
	}
}

type Subscription struct {
	ID               uuid.UUID         `gorm:"type:char(36);primaryKey"`
	UserID           uuid.UUID         `json:"user_id" gorm:"type:char(36);not null"`
	PlanName         string            `json:"plan_name" gorm:"type:varchar(100);not null"`
	StartDate        time.Time         `json:"start_date" gorm:"not null"`
	EndDate          time.Time         `json:"end_date" gorm:"not null"`
	Status           SuscriptionStatus `json:"status" gorm:"default:active;not null"`
	SubscriptionType SubscriptionType  `json:"type" gorm:"default:'ACTIVE';not null"`
	PaymentID        uuid.UUID         `json:"payment_id" gorm:"type:char(36);not null"`
	ElapsedTime      int64             `gorm:"-" json:"elapsed_time"`
	CreatedAt        time.Time         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time         `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt    `json:"deleted_at"`
}

func NewSubscription(
	userID uuid.UUID,
	planName string,
	paymentID uuid.UUID,
	status SuscriptionStatus,
	subscriptiontype SubscriptionType) *Subscription {
	now := time.Now()

	suscription := &Subscription{
		ID:               uuid.New(),
		UserID:           userID,
		PlanName:         planName,
		Status:           status,
		SubscriptionType: subscriptiontype,
		PaymentID:        paymentID,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	suscription.AssignSubscriptionDates()

	return suscription
}

func (s *Subscription) GetID() uuid.UUID             { return s.ID }
func (s *Subscription) GetUserID() uuid.UUID         { return s.UserID }
func (s *Subscription) GetPlanName() string          { return s.PlanName }
func (s *Subscription) GetStartDate() time.Time      { return s.StartDate }
func (s *Subscription) GetEndDate() time.Time        { return s.EndDate }
func (s *Subscription) GetStatus() SuscriptionStatus { return s.Status }
func (s *Subscription) GetPaymentID() uuid.UUID      { return s.PaymentID }
func (s *Subscription) GetType() SubscriptionType    { return s.SubscriptionType }
func (s *Subscription) GetCreatedAt() time.Time      { return s.CreatedAt }
func (s *Subscription) GetUpdatedAt() time.Time      { return s.UpdatedAt }

func (s *Subscription) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}

	if s.StartDate.IsZero() {
		s.StartDate = time.Now()
	}

	if s.EndDate.IsZero() {
		s.EndDate = s.StartDate.AddDate(0, 1, 0)
	}

	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}
	if s.UpdatedAt.IsZero() {
		s.UpdatedAt = time.Now()
	}

	return nil
}

func (s *Subscription) AssignSubscriptionDates() {
	startDate, endDate := calculatePeriodDates(s.SubscriptionType)

	s.StartDate = startDate
	s.EndDate = endDate
}

func (s *Subscription) GetElapsedTime() time.Duration {
	if s.StartDate.IsZero() {
		return 0
	}
	return time.Since(s.StartDate)
}

func (s *Subscription) GetElapsedTimeInDays() int {
	elapsed := s.GetElapsedTime()
	return int(elapsed.Hours() / 24)
}

func (s *Subscription) Cancel() error {
	if s.Status != ACTIVE {
		return errors.New("only active subscriptions can't be cancelled")
	}

	s.Status = CANCELLED
	s.UpdatedAt = time.Now()

	return nil
}

func calculatePeriodDates(subscrptiontype SubscriptionType) (time.Time, time.Time) {
	now := time.Now()

	durationMap := map[SubscriptionType]time.Duration{
		FREE_TRIAL:   7 * 24 * time.Hour,   // 7 days
		MONTLY:       30 * 24 * time.Hour,  // 30 days
		THREE_MONTHS: 90 * 24 * time.Hour,  // 90 days
		ANUALLY:      365 * 24 * time.Hour, // 1 yeat
	}

	duration, exists := durationMap[subscrptiontype]
	if !exists {
		duration = 7 * 24 * time.Hour
	}

	startDate := now
	endDate := now.Add(duration)
	return startDate, endDate
}

func (s *Subscription) SetType(subType *SubscriptionType) {
	s.SubscriptionType = *subType
}

func (s *Subscription) ExtendSubscription() {
	_, endDate := calculatePeriodDates(s.SubscriptionType)

	s.EndDate = endDate
	s.UpdatedAt = time.Now()
}
