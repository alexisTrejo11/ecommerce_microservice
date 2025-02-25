package dto

import "time"

// SessionDTO represents a user session, including metadata such as user agent, client IP, and expiration time.
type SessionDTO struct {
	// The unique identifier of the session.
	// Example: abc123xyz
	// @Param id body string true "Session ID"
	ID string `json:"id"`

	// The ID of the user associated with this session.
	// Example: user_456
	// @Param user_id body string true "User ID"
	UserID string `json:"user_id"`

	// The user agent string of the client that initiated the session.
	// Example: Mozilla/5.0 (Windows NT 10.0; Win64; x64)
	// @Param user_agent body string false "User Agent"
	UserAgent string `json:"user_agent"`

	// The IP address of the client that initiated the session.
	// Example: 192.168.1.1
	// @Param client_ip body string false "Client IP"
	ClientIP string `json:"client_ip"`

	// The expiration timestamp of the session in ISO 8601 format.
	// Example: 2025-12-31T23:59:59Z
	// @Param expires_at body string true "Expiration Time (ISO 8601)"
	ExpiresAt string `json:"expires_at"`

	// The timestamp when the session was created.
	// Example: 2025-02-25T12:00:00Z
	// @Param created_at body string true "Creation Time (ISO 8601)"
	CreatedAt time.Time `json:"created_at"`
}
