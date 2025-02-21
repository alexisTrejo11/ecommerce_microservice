package input

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/google/uuid"
)

type MFAUseCase interface {
	SetupMFA(ctx context.Context, userID uuid.UUID) (string, string, error)
	VerifyAndEnableMFA(ctx context.Context, userID uuid.UUID, code string) (*TokenDetails, error)
	DisableMFA(ctx context.Context, userID uuid.UUID, code string) error
	VerifyMFA(ctx context.Context, userID uuid.UUID, code string) error
	GetMFA(ctx context.Context, userID uuid.UUID) (*entities.MFA, error)
}
