package models

import (
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/api/dto"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
)

type UserModel struct {
	ID           string         `json:"id" gorm:"type:char(36);primary_key"`
	Email        string         `json:"email" gorm:"unique;not null"`
	Username     string         `json:"username" gorm:"unique;not null"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	PasswordHash string         `json:"-" gorm:"not null"`
	Phone        string         `json:"phone"`
	RoleID       uint           `json:"role_id" gorm:"not null"`
	Role         *RoleModel     `json:"role,omitempty" gorm:"foreignKey:RoleID"`
	Addresses    []AddressModel `json:"addresses,omitempty" gorm:"foreignKey:UserID"`
	Status       UserStatus     `json:"status" gorm:"type:int;default:0"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    *time.Time     `json:"-" gorm:"index"`
}

func (m UserModel) ToEntity() *entities.User {
	return &entities.User{
		ID:           m.ID,
		Email:        m.Email,
		Username:     m.Username,
		FirstName:    m.FirstName,
		LastName:     m.LastName,
		PasswordHash: m.PasswordHash,
		Phone:        m.Phone,
		RoleID:       m.RoleID,
		Role:         nil,                           //m.Role.ToEntity(),
		Addresses:    nil,                           //ToAddressEntities(m.Addresses),
		Status:       entities.UserStatus(m.Status), //m.Status,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
		DeletedAt:    m.DeletedAt,
	}
}

// Mapea UserModel a UserDTO
func (m *UserModel) ToDTO() *dto.UserDTO {
	return &dto.UserDTO{
		ID:        m.ID,
		Email:     m.Email,
		Username:  m.Username,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Phone:     m.Phone,
		RoleID:    m.RoleID,
		RoleName:  m.Role.Name,
		Status:    int(m.Status),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: m.DeletedAt,
	}
}

type UserStatus int

const (
	UserStatusPending  UserStatus = 0
	UserStatusActive   UserStatus = 1
	UserStatusInactive UserStatus = 2
	UserStatusBanned   UserStatus = 3
)

type PasswordResetModel struct {
	ID        string     `json:"id" gorm:"type:char(36);primary_key"`
	UserID    string     `json:"user_id" gorm:"type:char(36);not null"`
	User      *UserModel `json:"-" gorm:"foreignKey:UserID"`
	Token     string     `json:"-" gorm:"not null;unique"`
	ExpiresAt time.Time  `json:"expires_at" gorm:"not null"`
	Used      bool       `json:"used" gorm:"default:false"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"-" gorm:"index"`
}
