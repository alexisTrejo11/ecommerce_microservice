package dtos

import "time"

type NotificationMessageDTO struct {
	ID        string            `json:"id"`
	UserID    string            `json:"user_id"`
	Type      string            `json:"type"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
}

type NotificationDTO struct {
	ID          string            `json:"id"`
	UserID      string            `json:"user_id"`
	Type        string            `json:"type"`
	Title       string            `json:"title"`
	Content     string            `json:"content"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	Status      string            `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	SentAt      *time.Time        `json:"sent_at,omitempty"`
	ScheduledAt *time.Time        `json:"scheduled_at,omitempty"`
}
