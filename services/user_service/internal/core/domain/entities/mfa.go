package entities

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MFA struct {
	ID                   uint
	UserID               uuid.UUID
	User                 *User
	Enabled              bool
	Secret               string
	BackupCodes          []string
	BackupCodeExpiration time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            *time.Time
}

func NewMFA(userId uuid.UUID) *MFA {
	return &MFA{
		UserID:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (m *MFA) GenerateSecret() error {
	if m.Secret != "" {
		return errors.New("MFA already has a secret")
	}

	secret := make([]byte, 16)
	_, err := rand.Read(secret)
	if err != nil {
		return err
	}

	m.Secret = base32.StdEncoding.EncodeToString(secret)

	return nil
}

func (m *MFA) GenerateBackupCodes(n int) error {
	if len(m.BackupCodes) > 0 {
		return errors.New("backup codes already generated")
	}

	codes := []string{}
	for i := 0; i < n; i++ {
		code := generateRandomCode(8)
		codes = append(codes, code)
	}

	m.BackupCodes = codes
	m.BackupCodeExpiration = time.Now().Add(48 * time.Hour)

	return nil
}

func generateRandomCode(length int) string {
	code := make([]byte, length)
	_, err := rand.Read(code)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", code)[:length]
}

func (m *MFA) ValidateBackupCode(code string) bool {
	if time.Now().After(m.BackupCodeExpiration) {
		m.BackupCodes = []string{}
		return false
	}

	for _, storedCode := range m.BackupCodes {
		if storedCode == code {
			m.BackupCodes = removeBackupCode(m.BackupCodes, code)
			return true
		}
	}

	return false
}

func removeBackupCode(codes []string, code string) []string {
	var newCodes []string
	for _, storedCode := range codes {
		if storedCode != code {
			newCodes = append(newCodes, storedCode)
		}
	}
	return newCodes
}
