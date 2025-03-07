package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartModel struct {
	ID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID       `gorm:"type:uuid;not null"`
	Items     []CartItemModel `gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartItemModel struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CartID    uuid.UUID `gorm:"type:uuid;not null"`
	ProductID uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `gorm:"size:255;not null"`
	UnitPrice float64   `gorm:"not null"`
	Quantity  int       `gorm:"not null"`
	Discount  float64   `gorm:"default:0"`
	AddedAt   time.Time
}

func (c *CartModel) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	return
}

func (ci *CartItemModel) BeforeCreate(tx *gorm.DB) (err error) {
	if ci.ID == uuid.Nil {
		ci.ID = uuid.New()
	}
	ci.AddedAt = time.Now()
	return
}
