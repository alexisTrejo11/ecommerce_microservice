package input

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/google/uuid"
)

type SessionUseCase interface {
	GetUserSessions(ctx context.Context, userID uuid.UUID) ([]*entities.Session, error)
	DeleteSession(ctx context.Context, id uuid.UUID, userID uuid.UUID) error
}
