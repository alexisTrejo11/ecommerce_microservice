package dto

// SignupDTO represents the data required for user registration.
// It includes the user's email, username, full name, phone number, and password.
// This structure is used to validate and parse the incoming registration data.
type SignupDTO struct {
	// The user's email address, which must be a valid email format.
	// Example: user@example.com
	// @Param email body string true "Email address" validate:"required,email"
	Email string `json:"email" validate:"required,email"`

	// The user's username, which must be between 3 and 30 characters.
	// Example: johndoe
	// @Param username body string true "Username" validate:"required,min=3,max=30"
	Username string `json:"username" validate:"required,min=3,max=30"`

	// The user's first name, which must be between 2 and 50 characters.
	// Example: John
	// @Param first_name body string true "First Name" validate:"required,min=2,max=50"
	FirstName string `json:"first_name" validate:"required,min=2,max=50"`

	// The user's last name, which must be between 2 and 50 characters.
	// Example: Doe
	// @Param last_name body string true "Last Name" validate:"required,min=2,max=50"
	LastName string `json:"last_name" validate:"required,min=2,max=50"`

	// The user's phone number, which must be exactly 10 digits.
	// Example: 1234567890
	// @Param phone body string true "Phone Number" validate:"required,len=10"
	Phone string `json:"phone" validate:"required,len=10"`

	// The user's password, which must be at least 8 characters long.
	// Example: password123
	// @Param password body string true "Password" validate:"required,min=8"
	Password string `json:"password" validate:"required,min=8"`
}
