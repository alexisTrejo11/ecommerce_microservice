package dto

import (
	"time"
)

// UserDTO represents a user entity, including personal details, role information, and timestamps.
type UserDTO struct {
	// The unique identifier of the user.
	// Example: user_123
	// @Param id body string true "User ID"
	ID string `json:"id"`

	// The user's email address.
	// Example: user@example.com
	// @Param email body string true "Email address"
	Email string `json:"email"`

	// The user's username.
	// Example: johndoe
	// @Param username body string true "Username"
	Username string `json:"username"`

	// The user's first name.
	// Example: John
	// @Param first_name body string true "First Name"
	FirstName string `json:"first_name"`

	// The user's last name.
	// Example: Doe
	// @Param last_name body string true "Last Name"
	LastName string `json:"last_name"`

	// The user's phone number.
	// Example: 1234567890
	// @Param phone body string true "Phone Number"
	Phone string `json:"phone"`

	// The ID of the user's role.
	// Example: 1
	// @Param role_id body uint true "Role ID"
	RoleID uint `json:"role_id"`

	// The name of the user's role.
	// Example: Admin
	// @Param role_name body string true "Role Name"
	RoleName string `json:"role_name"`

	// The status of the user account (e.g., active, inactive).
	// Example: 1
	// @Param status body int true "User Status (e.g., 1 for active, 0 for inactive)"
	Status int `json:"status"`

	// The timestamp when the user was created.
	// Example: 2025-02-25T12:00:00Z
	// @Param created_at body string true "Creation Time (ISO 8601)"
	CreatedAt time.Time `json:"created_at"`

	// The timestamp when the user was last updated.
	// Example: 2025-02-26T15:30:00Z
	// @Param updated_at body string true "Last Updated Time (ISO 8601)"
	UpdatedAt time.Time `json:"updated_at"`

	// The timestamp when the user was deleted, if applicable.
	// Example: 2025-03-01T10:00:00Z
	// @Param deleted_at body string false "Deletion Time (ISO 8601, nullable)"
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
