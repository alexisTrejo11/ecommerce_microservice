package entities

import (
	"errors"
	"fmt"
	"regexp"
	"time"
	"unicode"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string
	Email        string
	Username     string
	FirstName    string
	LastName     string
	PasswordHash string
	Phone        string
	RoleID       uint
	Role         *Role
	Addresses    []Address
	Permisions   []string
	Status       UserStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

type UserStatus int

const (
	UserStatusPending  UserStatus = 0
	UserStatusActive   UserStatus = 1
	UserStatusInactive UserStatus = 2
	UserStatusBanned   UserStatus = 3
)

type PasswordReset struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	User      *User
	Token     string
	ExpiresAt time.Time
	Used      bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
}

func (u *User) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func (u *User) Validate() error {
	if err := u.ValidateEmail(); err != nil {
		return err
	}
	if err := u.ValidateUsername(); err != nil {
		return err
	}
	if u.PasswordHash == "" {
		return errors.New("password hash cannot be empty")
	}
	if u.RoleID == 0 {
		return errors.New("role ID is required")
	}
	if u.Phone != "" {
		if err := u.ValidatePhone(); err != nil {
			return err
		}
	}
	if err := ValidatePasswordStrength(u.PasswordHash); err != nil {
		return err
	}

	return nil
}

func (u *User) ValidateEmail() error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if matched, _ := regexp.MatchString(emailRegex, u.Email); !matched {
		return errors.New("invalid email format")
	}
	return nil
}

func (u *User) ValidateUsername() error {
	if len(u.Username) < 3 || len(u.Username) > 20 {
		return errors.New("username must be between 3 and 20 characters")
	}
	usernameRegex := `^[a-zA-Z0-9_]+$`
	if matched, _ := regexp.MatchString(usernameRegex, u.Username); !matched {
		return errors.New("username can only contain alphanumeric characters and underscores")
	}
	return nil
}

func (u *User) ValidatePhone() error {
	phoneRegex := `^\+?[1-9]\d{1,14}$` // format E.164
	if matched, _ := regexp.MatchString(phoneRegex, u.Phone); !matched {
		return errors.New("invalid phone number format")
	}
	return nil
}

func ValidatePasswordStrength(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return errors.New("password must contain at least one number")
	}
	if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}
	return nil
}

func (u *User) CheckLoginEligibility() error {
	switch u.Status {
	case UserStatusActive:
		return nil
	case UserStatusPending:
		return errors.New("account pending activation")
	case UserStatusInactive:
		return errors.New("account deactivated")
	case UserStatusBanned:
		return errors.New("account banned")
	default:
		return errors.New("invalid account status")
	}
}

func (u *User) Activate() error {
	if u.Status != UserStatusPending {
		return errors.New("only pending accounts can be activated")
	}
	u.Status = UserStatusActive
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) Deactivate() error {
	if u.Status == UserStatusBanned {
		return errors.New("banned accounts cannot be deactivated")
	}
	u.Status = UserStatusInactive
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) Ban(reason string) error {
	if u.Status == UserStatusBanned {
		return errors.New("account is already banned")
	}
	u.Status = UserStatusBanned
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) UpdateEmail(newEmail string) error {
	if u.Email == newEmail {
		return nil
	}
	previousEmail := u.Email
	u.Email = newEmail
	if err := u.ValidateEmail(); err != nil {
		u.Email = previousEmail
		return err
	}
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) UpdatePhone(newPhone string) error {
	if newPhone == "" {
		u.Phone = ""
		return nil
	}

	previousPhone := u.Phone
	u.Phone = newPhone
	if err := u.ValidatePhone(); err != nil {
		u.Phone = previousPhone
		return err
	}
	u.UpdatedAt = time.Now()
	return nil
}

func (pr *PasswordReset) IsValid() bool {
	return !pr.Used && time.Now().Before(pr.ExpiresAt)
}

func (pr *PasswordReset) MarkAsUsed() error {
	if pr.Used {
		return errors.New("token already used")
	}
	if time.Now().After(pr.ExpiresAt) {
		return errors.New("token expired")
	}
	pr.Used = true
	pr.UpdatedAt = time.Now()
	return nil
}

func (u *User) UpdatePassword(oldPassword, newPassword string) error {
	if err := u.ComparePassword(oldPassword); err != nil {
		return errors.New("invalid current password")
	}

	if err := ValidatePasswordStrength(newPassword); err != nil {
		return err
	}

	hashedPassword, err := u.HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	u.PasswordHash = hashedPassword
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) ActivateAccount() {
	u.Status = UserStatusActive
}

func (u *User) BanAccount() {
	u.Status = UserStatusBanned
}

func (u *User) SetAsInactive() {
	u.Status = UserStatusInactive
}
