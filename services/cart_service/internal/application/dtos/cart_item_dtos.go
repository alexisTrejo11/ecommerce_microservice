package dtos

import (
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/pkg/facadeService"
	"github.com/google/uuid"
)

type CartItemInserDTO struct {
	ProductIDs uuid.UUID `json:"product_id"`
	Quantity   int       `json:"quantity"`
}

type CartItemFetchedDTO struct {
	ProductData facadeService.Product `json:"product_data"`
	Quantity    int
}

type CartItemDTO struct {
	ID        uuid.UUID `json:"id"`
	CartID    uuid.UUID `json:"cart_id"`
	ProductID uuid.UUID `json:"product_id"`
	Name      string    `json:"name"`
	UnitPrice float64   `json:"unit_price"`
	Quantity  int       `json:"quantity"`
	Discount  float64   `json:"discount"`
}
