package certificate

import (
	"time"

	"github.com/google/uuid"
)

type Certificate struct {
	ID             uuid.UUID `gorm:"type:char(36);primaryKey"`
	EnrollmentID   uuid.UUID `json:"enrollment_id"`
	IssuedAt       time.Time `json:"issued_at"`
	CertificateURL string    `json:"certificate_url"`
	ExpiresAt      time.Time `json:"expires_at,omitempty"`
}
