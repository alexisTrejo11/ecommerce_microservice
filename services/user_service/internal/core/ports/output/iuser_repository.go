package output

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/google/uuid"
)

type IUserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	FindByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	FindByUsername(ctx context.Context, username string) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) error
	UpdateStatus(ctx context.Context, id uuid.UUID, status entities.UserStatus) error
	Delete(ctx context.Context, id uuid.UUID) error
	//FindByMetadata(ctx context.Context, key, value string) (*entities.User, error)
	//SetMetadata(ctx context.Context, id uuid.UUID, key, value string) error
	//GetMetadata(ctx context.Context, id uuid.UUID, key string) (string, error)
	//DeleteMetadata(ctx context.Context, id uuid.UUID, key string) error
	//FindRoleByID(ctx context.Context, id uint) (*entities.Role, error)
	//ListRoles(ctx context.Context) ([]*entities.Role, error)
}

type MFARepository interface {
	Create(ctx context.Context, mfa *entities.MFA) error
	FindByUserID(ctx context.Context, userID uuid.UUID) (*entities.MFA, error)
	Update(ctx context.Context, mfa *entities.MFA) error
	Delete(ctx context.Context, userID uuid.UUID) error
}

type PasswordResetRepository interface {
	Create(ctx context.Context, reset *entities.PasswordReset) error
	FindByToken(ctx context.Context, token string) (*entities.PasswordReset, error)
	FindByUserID(ctx context.Context, userID uuid.UUID) (*entities.PasswordReset, error)
	MarkAsUsed(ctx context.Context, id uuid.UUID) error
	DeleteExpired(ctx context.Context) error
}

type AddressRepository interface {
	Create(ctx context.Context, address *entities.Address) error
	FindByID(ctx context.Context, id uint) (*entities.Address, error)
	FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]*entities.Address, error)
	FindDefaultByUserID(ctx context.Context, userID uuid.UUID) (*entities.Address, error)
	Update(ctx context.Context, address *entities.Address) error
	SetDefault(ctx context.Context, id uint, userID uuid.UUID) error
	Delete(ctx context.Context, id uint) error
}
