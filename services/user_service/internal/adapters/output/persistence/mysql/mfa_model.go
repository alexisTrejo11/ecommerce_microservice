package models

import (
	"time"
)

type MFAModel struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	UserID      string     `json:"user_id" gorm:"type:char(36)"`
	User        *UserModel `json:"-" gorm:"foreignKey:UserID"`
	Enabled     bool       `json:"enabled" gorm:"default:false"`
	Secret      string     `json:"-"`
	BackupCodes string     `json:"backup_codes"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `json:"-" gorm:"index"`
}
