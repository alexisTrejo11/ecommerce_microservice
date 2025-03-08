package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartModel struct {
	ID        string          `gorm:"type:char(36);primaryKey"`
	UserID    string          `gorm:"type:char(36);not null"`
	Items     []CartItemModel `gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartItemModel struct {
	ID        string  `gorm:"type:char(36);primaryKey"`
	CartID    string  `gorm:"type:char(36);not null"`
	ProductID string  `gorm:"type:char(36);not null"`
	Name      string  `gorm:"size:255;not null"`
	UnitPrice float64 `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Discount  float64 `gorm:"default:0"`
	AddedAt   time.Time
}

// Antes de crear un registro, genera un UUID manualmente
func (c *CartModel) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	return
}

func (ci *CartItemModel) BeforeCreate(tx *gorm.DB) (err error) {
	if ci.ID == "" {
		ci.ID = uuid.New().String()
	}
	ci.AddedAt = time.Now()
	return
}
