package input

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/api/dto"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/google/uuid"
)

type AddressUseCase interface {
	GetUserAddresses(ctx context.Context, userID uuid.UUID) ([]*dto.AddressDTO, error)
	AddAddress(ctx context.Context, address *dto.AddressInsertDTO) error
	UpdateAddress(ctx context.Context, id uint, address *dto.AddressInsertDTO) error
	DeleteAddress(ctx context.Context, id uint, userID uuid.UUID) error
	SetDefaultAddress(ctx context.Context, id uint, userID uuid.UUID) error
}

type SessionUseCase interface {
	GetUserSessions(ctx context.Context, userID uuid.UUID) ([]*entities.Session, error)
	DeleteSession(ctx context.Context, id uuid.UUID, userID uuid.UUID) error
}
