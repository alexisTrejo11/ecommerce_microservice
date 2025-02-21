package output

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/google/uuid"
)

type SessionRepository interface {
	Create(ctx context.Context, session *entities.Session) error
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Session, error)
	FindByRefreshToken(ctx context.Context, token string) (*entities.Session, error)
	FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]*entities.Session, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteAllByUserID(ctx context.Context, userID uuid.UUID) error
	DeleteExpired(ctx context.Context) error
}
