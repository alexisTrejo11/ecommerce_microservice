package models

import (
	"time"
)

type AddressModel struct {
	ID           uint       `json:"id" gorm:"primary_key"`
	UserID       string     `json:"user_id" gorm:"type:char(36)"`
	AddressLine1 string     `json:"address_line1" gorm:"not null"`
	AddressLine2 string     `json:"address_line2"`
	City         string     `json:"city" gorm:"not null"`
	State        string     `json:"state" gorm:"not null"`
	PostalCode   string     `json:"postal_code" gorm:"not null"`
	Country      string     `json:"country" gorm:"not null"`
	IsDefault    bool       `json:"is_default" gorm:"default:false"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    *time.Time `json:"-" gorm:"index"`
}
