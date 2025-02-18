package models

import (
	"time"

	"github.com/google/uuid"
)

// Multi-Factor Authentication settings for a user
type MFA struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	UserID      uuid.UUID  `json:"user_id" gorm:"type:uuid;unique;not null"`
	User        *User      `json:"-" gorm:"foreignKey:UserID"`
	Enabled     bool       `json:"enabled" gorm:"default:false"`
	Secret      string     `json:"-"`
	BackupCodes []string   `json:"-" gorm:"type:text[]"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `json:"-" gorm:"index"`
}
