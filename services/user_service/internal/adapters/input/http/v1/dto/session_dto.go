package dto

import "time"

type SessionDTO struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	UserAgent string    `json:"user_agent"`
	ClientIP  string    `json:"client_ip"`
	ExpiresAt string    `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}
