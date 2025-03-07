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
