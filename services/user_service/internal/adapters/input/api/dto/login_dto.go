package dto

type LoginDTO struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	UserAgent string `json:"user_agent"`
	ClientIP  string `json:"client_ip"`
}
