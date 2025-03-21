package dtos

import (
	"time"

	"github.com/google/uuid"
)

// CertificateInsertDTO defines the structure for creating a new certificate.
// swagger:model CertificateInsertDTO
type CertificateInsertDTO struct {
	// The ID of the enrollment associated with the certificate.
	// Required: true
	// Example: 550e8400-e29b-41d4-a716-446655440000
	EnrollmentID uuid.UUID `json:"enrollment_id" validate:"required"`

	// The date and time when the certificate was issued.
	// Required: true
	// Format: date-time
	// Example: 2023-10-01T12:34:56
	IssuedAt string `json:"issued_at" validate:"required,datetime=2006-01-02T15:04:05"`

	// The URL where the certificate can be accessed.
	// Required: true
	// Example: https://example.com/certificates/12345
	CertificateURL string `json:"certificate_url" validate:"required,url"`

	// The date and time when the certificate expires (optional).
	// Format: date-time
	// Example: 2024-10-01T12:34:56
	ExpiresAt string `json:"expires_at,omitempty" validate:"omitempty,datetime=2006-01-02T15:04:05"`
}

// CertificateDTO defines the structure for representing a certificate.
// swagger:model CertificateDTO
type CertificateDTO struct {
	// The unique identifier of the certificate.
	// Example: 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `json:"id"`

	// The ID of the enrollment associated with the certificate.
	// Example: 550e8400-e29b-41d4-a716-446655440001
	EnrollmentID uuid.UUID `json:"enrollment_id"`

	// The date and time when the certificate was issued.
	// Format: date-time
	// Example: 2023-10-01T12:34:56
	IssuedAt time.Time `json:"issued_at"`

	// The URL where the certificate can be accessed.
	// Example: https://example.com/certificates/12345
	CertificateURL string `json:"certificate_url"`

	// The date and time when the certificate expires (optional).
	// Format: date-time
	// Example: 2024-10-01T12:34:56
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}
