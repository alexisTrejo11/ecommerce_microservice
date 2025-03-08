package dtos

import (
	"time"

	"github.com/google/uuid"
)

type PurchaseDetails struct {
	UserID         uuid.UUID
	TotalAmount    float64
	TotalDisscount float64
	PurchaseDate   time.Time
	Items          []ItemDTO
}

type ItemDTO struct {
	Name      string  `json:"name"`
	UnitPrice float64 `json:"unit_price"`
	Quantity  int     `json:"quantity"`
	Discount  float64 `json:"discount"`
}

type OrderDetails struct {
	UserID    uuid.UUID
	ItemCount uint
	CreatedAt time.Time
	Items     []ItemDTO
}
