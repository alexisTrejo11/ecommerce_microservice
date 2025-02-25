package dto

import (
	"github.com/google/uuid"
)

type AddressDTO struct {
	ID           uint      `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	AddressLine1 string    `json:"address_line_1"`
	AddressLine2 *string   `json:"address_line_2,omitempty"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	PostalCode   string    `json:"postal_code"`
	Country      string    `json:"country"`
	IsDefault    bool      `json:"is_default"`
}

type AddressInsertDTO struct {
	UserID       uuid.UUID `json:"user_id"`
	AddressLine1 string    `json:"address_line_1" validate:"required,min=5,max=100"`
	AddressLine2 *string   `json:"address_line_2,omitempty" validate:"omitempty,min=5,max=100"`
	City         string    `json:"city" validate:"required,min=2,max=50"`
	State        string    `json:"state" validate:"required,min=2,max=50"`
	PostalCode   string    `json:"postal_code" validate:"required"`
	Country      string    `json:"country" validate:"required,len=2"`
	IsDefault    bool      `json:"is_default"`
}
