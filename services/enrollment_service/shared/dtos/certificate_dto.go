package dtos

import (
	"time"

	"github.com/google/uuid"
)

type CertificateInsertDTO struct {
	EnrollmentID   uuid.UUID `json:"enrollment_id" validate:"required"`
	IssuedAt       string    `json:"issued_at" validate:"required,datetime=2006-01-02T15:04:05"`
	CertificateURL string    `json:"certificate_url" validate:"required,url"`
	ExpiresAt      string    `json:"expires_at,omitempty" validate:"omitempty,datetime=2006-01-02T15:04:05"`
}

type CertificateDTO struct {
	ID             uuid.UUID  `json:"id"`
	EnrollmentID   uuid.UUID  `json:"enrollment_id"`
	IssuedAt       time.Time  `json:"issued_at"`
	CertificateURL string     `json:"certificate_url"`
	ExpiresAt      *time.Time `json:"expires_at,omitempty"`
}
