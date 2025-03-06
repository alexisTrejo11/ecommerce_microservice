package domain

import (
	"time"

	"github.com/google/uuid"
)

type CartItem struct {
	ID        uuid.UUID
	CartID    uuid.UUID
	ProductID uuid.UUID
	Quantity  uint
	AddedAt   time.Time
	UpdateAt  time.Time
}
