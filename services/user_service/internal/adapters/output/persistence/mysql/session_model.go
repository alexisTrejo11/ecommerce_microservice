package models

import (
	"time"

	"gorm.io/gorm"
)

type SessionModel struct {
	ID           string    `json:"id" gorm:"type:char(36)"`
	UserID       string    `json:"user_id" gorm:"type:char(36)"`
	RefreshToken string    `gorm:"type:text;not null"`
	UserAgent    string    `gorm:"type:varchar(255);not null"`
	ClientIP     string    `gorm:"type:varchar(45);not null"`
	ExpiresAt    time.Time `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (SessionModel) TableName() string {
	return "sessions"
}
