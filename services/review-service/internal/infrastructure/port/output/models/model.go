package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewModel struct {
	ID         uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID     uuid.UUID `gorm:"type:uuid;not null"`
	CourseID   uuid.UUID `gorm:"type:uuid;not null"`
	Rating     int       `gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Comment    string    `gorm:"type:text;default:null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	IsApproved bool      `gorm:"default:false"`
}

func (r *ReviewModel) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}

func (r *ReviewModel) TableName() string {
	return "reviews"
}
