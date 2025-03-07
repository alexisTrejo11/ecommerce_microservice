package domain

import (
	"time"

	"github.com/google/uuid"
)

type CartItem struct {
	ID        uuid.UUID
	ProductID uuid.UUID
	Name      string
	UnitPrice float64
	Quantity  int
	Discount  float64
	CreatedAt time.Time
}

func NewCartItem(productID uuid.UUID, name string, unitPrice float64, quantity int, discount float64) CartItem {
	return CartItem{
		ID:        uuid.New(),
		ProductID: productID,
		Name:      name,
		UnitPrice: unitPrice,
		Quantity:  quantity,
		Discount:  discount,
		CreatedAt: time.Now(),
	}
}
