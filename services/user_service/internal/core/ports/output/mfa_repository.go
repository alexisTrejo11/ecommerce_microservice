package output

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/google/uuid"
)

type MFARepository interface {
	Create(ctx context.Context, mfa *entities.MFA) error
	FindByUserID(ctx context.Context, userID uuid.UUID) (*entities.MFA, error)
	Update(ctx context.Context, mfa *entities.MFA) error
	Delete(ctx context.Context, userID uuid.UUID) error
}
