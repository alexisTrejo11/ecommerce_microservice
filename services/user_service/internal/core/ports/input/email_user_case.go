package input

import (
	"context"

	"github.com/google/uuid"
)

type EmailUseCase interface {
	SendVerificationEmail(ctx context.Context, userID uuid.UUID, token string) error
	VerifyEmail(ctx context.Context, token string) error
}
