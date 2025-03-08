package dtos

import (
	"github.com/google/uuid"
)

type CartDTO struct {
	ID     uuid.UUID     `json:"id"`
	UserID uuid.UUID     `json:"user_id"`
	Items  []CartItemDTO `json:"items"`
}
