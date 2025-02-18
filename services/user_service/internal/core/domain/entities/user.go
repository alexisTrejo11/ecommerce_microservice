package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uuid.UUID  `json:"id" gorm:"type:uuid;primary_key"`
	Email        string     `json:"email" gorm:"unique;not null"`
	Username     string     `json:"username" gorm:"unique;not null"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	PasswordHash string     `json:"-" gorm:"not null"`
	Phone        string     `json:"phone"`
	RoleID       uint       `json:"role_id" gorm:"not null"`
	Role         *Role      `json:"role,omitempty" gorm:"foreignKey:RoleID"`
	Addresses    []Address  `json:"addresses,omitempty" gorm:"foreignKey:UserID"`
	Status       UserStatus `json:"status" gorm:"type:int;default:1"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    *time.Time `json:"-" gorm:"index"`
}

type UserStatus int

const (
	UserStatusPending  UserStatus = 0
	UserStatusActive   UserStatus = 1
	UserStatusInactive UserStatus = 2
	UserStatusBanned   UserStatus = 3
)

// PasswordReset represents password reset requests
type PasswordReset struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key"`
	UserID    uuid.UUID  `json:"user_id" gorm:"type:uuid;not null"`
	User      *User      `json:"-" gorm:"foreignKey:UserID"`
	Token     string     `json:"-" gorm:"not null;unique"`
	ExpiresAt time.Time  `json:"expires_at" gorm:"not null"`
	Used      bool       `json:"used" gorm:"default:false"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"-" gorm:"index"`
}

func CreateUser(email, username, password string, roleID uint) (*User, error) {
	id := uuid.New()

	// Hash the password
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create and return new user
	return &User{
		ID:           id,
		Email:        email,
		Username:     username,
		PasswordHash: hashedPassword,
		RoleID:       roleID,
		Status:       UserStatusPending,
	}, nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
}

func (u *User) UpdatePassword(password string) error {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}
	u.PasswordHash = hashedPassword
	return nil
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
