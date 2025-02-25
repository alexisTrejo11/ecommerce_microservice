package dto

// LoginDTO represents the data required for user authentication.
// It includes the user's email and password, along with optional metadata such as user agent and client IP.
type LoginDTO struct {
	// The user's email address, which must be a valid email format.
	// Example: user@example.com
	// @Param email body string true "Email address" validate:"required,email"
	Email string `json:"email" validate:"required,email"`

	// The user's password, which must be at least 8 characters long.
	// Example: password123
	// @Param password body string true "Password" validate:"required,min=8"
	Password string `json:"password" validate:"required,min=8"`

	// The user agent string of the client making the request.
	// Example: Mozilla/5.0 (Windows NT 10.0; Win64; x64)
	// @Param user_agent body string false "User Agent"
	UserAgent string `json:"user_agent"`

	// The IP address of the client making the request.
	// Example: 192.168.1.1
	// @Param client_ip body string false "Client IP"
	ClientIP string `json:"client_ip"`
}
