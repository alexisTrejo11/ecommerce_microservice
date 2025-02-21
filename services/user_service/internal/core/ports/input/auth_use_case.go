package input

import (
	"context"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/api/dto"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/google/uuid"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
	SessionID    uuid.UUID
}

type AuthUseCase interface {
	Register(ctx context.Context, singupDto dto.SignupDTO) (*entities.User, string, error)
	Login(ctx context.Context, loginDTO dto.LoginDTO) (*TokenDetails, error)
	RefreshTokens(ctx context.Context, refreshToken, userAgent, clientIP string) (*TokenDetails, error)
	ResendCode(ctx context.Context, codeType string, userID uuid.UUID) error
	Logout(ctx context.Context, refreshToken string, userID uuid.UUID) error
	LogoutAll(ctx context.Context, userID uuid.UUID) error
	ResetPassword(ctx context.Context, token, newPassword string) error
	ActivateAccount(ctx context.Context, token string) error
	//SetupMFA(ctx context.Context, userID uuid.UUID) (string, string, error)
	//VerifyAndEnableMFA(ctx context.Context, userID uuid.UUID, code string) ([]string, error)
	//DisableMFA(ctx context.Context, userID uuid.UUID, code string) error
	//VerifyMFA(ctx context.Context, userID uuid.UUID, code string) error
}
