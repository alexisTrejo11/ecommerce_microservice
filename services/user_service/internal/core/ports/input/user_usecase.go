package input

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/google/uuid"
)

type UserUseCase interface {
	GetUser(ctx context.Context, id uuid.UUID) (*entities.User, error)
	UpdateUser(ctx context.Context, user *entities.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	ChangePassword(ctx context.Context, id uuid.UUID, currentPassword, newPassword string) error
	GetUserAddresses(ctx context.Context, userID uuid.UUID) ([]*entities.Address, error)
	AddAddress(ctx context.Context, address *entities.Address) error
	UpdateAddress(ctx context.Context, address *entities.Address) error
	DeleteAddress(ctx context.Context, id uint, userID uuid.UUID) error
	SetDefaultAddress(ctx context.Context, id uint, userID uuid.UUID) error
	GetUserSessions(ctx context.Context, userID uuid.UUID) ([]*entities.Session, error)
	DeleteSession(ctx context.Context, id uuid.UUID, userID uuid.UUID) error
}
