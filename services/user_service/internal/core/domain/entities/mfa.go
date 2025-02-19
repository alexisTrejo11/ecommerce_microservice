package entities

import (
	"time"

	"github.com/google/uuid"
)

// Multi-Factor Authentication settings for a user
type MFA struct {
	ID          uint
	UserID      uuid.UUID
	User        *User
	Enabled     bool
	Secret      string
	BackupCodes []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
