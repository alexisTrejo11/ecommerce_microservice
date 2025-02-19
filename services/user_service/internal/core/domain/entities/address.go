package entities

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	ID           uint
	UserID       uuid.UUID
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	PostalCode   string
	Country      string
	IsDefault    bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
