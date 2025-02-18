package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID  `json:"id" gorm:"type:uuid;primary_key"`
	UserID       uuid.UUID  `json:"user_id" gorm:"type:uuid;not null"`
	User         *User      `json:"-" gorm:"foreignKey:UserID"`
	RefreshToken string     `json:"-" gorm:"not null;unique"`
	UserAgent    string     `json:"user_agent"`
	ClientIP     string     `json:"client_ip"`
	ExpiresAt    time.Time  `json:"expires_at" gorm:"not null"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    *time.Time `json:"-" gorm:"index"`
}
