package dto

import (
	"github.com/google/uuid"
)

// AddressDTO represents a user's address, including location details and whether it's the default address.
type AddressDTO struct {
	// The unique identifier of the address.
	// Example: 1
	// @Param id body uint true "Address ID"
	ID uint `json:"id"`

	// The unique identifier of the user associated with this address.
	// Example: 550e8400-e29b-41d4-a716-446655440000
	// @Param user_id body string true "User ID (UUID)"
	UserID uuid.UUID `json:"user_id"`

	// The first line of the address.
	// Example: 123 Main Street
	// @Param address_line_1 body string true "Address Line 1"
	AddressLine1 string `json:"address_line_1"`

	// The second line of the address (optional).
	// Example: Apt 4B
	// @Param address_line_2 body string false "Address Line 2 (Optional)"
	AddressLine2 *string `json:"address_line_2,omitempty"`

	// The city of the address.
	// Example: New York
	// @Param city body string true "City"
	City string `json:"city"`

	// The state or province of the address.
	// Example: NY
	// @Param state body string true "State/Province"
	State string `json:"state"`

	// The postal or ZIP code of the address.
	// Example: 10001
	// @Param postal_code body string true "Postal Code"
	PostalCode string `json:"postal_code"`

	// The country code of the address (ISO 3166-1 alpha-2 format).
	// Example: US
	// @Param country body string true "Country (ISO 3166-1 alpha-2)"
	Country string `json:"country"`

	// Whether this address is the default for the user.
	// Example: true
	// @Param is_default body bool true "Is Default Address"
	IsDefault bool `json:"is_default"`
}

// AddressInsertDTO represents the data required to insert a new address.
type AddressInsertDTO struct {
	// The unique identifier of the user associated with this address.
	// Example: 550e8400-e29b-41d4-a716-446655440000
	// @Param user_id body string true "User ID (UUID)"
	UserID uuid.UUID `json:"user_id"`

	// The first line of the address, required with a length between 5 and 100 characters.
	// Example: 123 Main Street
	// @Param address_line_1 body string true "Address Line 1" validate:"required,min=5,max=100"
	AddressLine1 string `json:"address_line_1" validate:"required,min=5,max=100"`

	// The second line of the address, optional with a length between 5 and 100 characters.
	// Example: Apt 4B
	// @Param address_line_2 body string false "Address Line 2 (Optional)" validate:"omitempty,min=5,max=100"
	AddressLine2 *string `json:"address_line_2,omitempty" validate:"omitempty,min=5,max=100"`

	// The city of the address, required with a length between 2 and 50 characters.
	// Example: New York
	// @Param city body string true "City" validate:"required,min=2,max=50"
	City string `json:"city" validate:"required,min=2,max=50"`

	// The state or province of the address, required with a length between 2 and 50 characters.
	// Example: NY
	// @Param state body string true "State/Province" validate:"required,min=2,max=50"
	State string `json:"state" validate:"required,min=2,max=50"`

	// The postal or ZIP code of the address, required.
	// Example: 10001
	// @Param postal_code body string true "Postal Code" validate:"required"
	PostalCode string `json:"postal_code" validate:"required"`

	// The country code of the address, required in ISO 3166-1 alpha-2 format.
	// Example: US
	// @Param country body string true "Country (ISO 3166-1 alpha-2)" validate:"required,len=2"
	Country string `json:"country" validate:"required,len=2"`

	// Whether this address is the default for the user.
	// Example: true
	// @Param is_default body bool true "Is Default Address"
	IsDefault bool `json:"is_default"`
}
