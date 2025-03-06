package domain

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Items     []CartItem
	CreatedAt time.Time
	UpdateAt  time.Time
}

func NewCart(userID uuid.UUID) *Cart {
	return &Cart{
		ID:        uuid.New(),
		UserID:    userID,
		Items:     []CartItem{},
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
}
