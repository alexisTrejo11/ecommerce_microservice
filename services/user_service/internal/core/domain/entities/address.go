package entities

import (
	"fmt"
	"regexp"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/errors"
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
	Country      string // Code of country ISO 3166-1 alpha-2 (ej: "MX", "US")
	IsDefault    bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func (a *Address) Validate() error {
	if a.AddressLine1 == "" {
		return fmt.Errorf("%w: address_line1", errors.ErrRequiredField)
	}

	if a.City == "" {
		return fmt.Errorf("%w: city", errors.ErrRequiredField)
	}

	if a.PostalCode == "" {
		return fmt.Errorf("%w: postal_code", errors.ErrRequiredField)
	}

	if !isValidCountryCode(a.Country) {
		return errors.ErrInvalidCountry
	}

	if !isValidPostalCode(a.PostalCode, a.Country) {
		return errors.ErrInvalidPostalCode
	}

	return nil
}

func (a *Address) PrepareForCreate() {
	a.CreatedAt = time.Now().UTC()
	a.UpdatedAt = time.Now().UTC()
}

func (a *Address) PrepareForUpdate() {
	a.UpdatedAt = time.Now().UTC()
}

func isValidCountryCode(country string) bool {
	if len(country) != 2 {
		return false
	}

	match, _ := regexp.MatchString("^[A-Z]{2}$", country)
	return match
}

func isValidPostalCode(postalCode, country string) bool {
	switch country {
	case "MX":
		return regexp.MustCompile(`^\d{5}$`).MatchString(postalCode)
	case "US":
		return regexp.MustCompile(`^\d{5}(-\d{4})?$`).MatchString(postalCode)
	case "CA":
		return regexp.MustCompile(`^[A-Za-z]\d[A-Za-z] \d[A-Za-z]\d$`).MatchString(postalCode)
	default:
		// Generic Validation for other countries
		return len(postalCode) >= 3 && len(postalCode) <= 10
	}
}
