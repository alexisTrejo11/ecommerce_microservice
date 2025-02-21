package usecases

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output"
	"github.com/google/uuid"
)

type MFAUseCase struct {
	mfaRepository output.MFARepository
	tokenService  output.TokenService
}

func NewMFAUseCase(mfaRepository output.MFARepository, tokenService output.TokenService) input.MFAUseCase {
	return &MFAUseCase{
		mfaRepository: mfaRepository,
		tokenService:  tokenService,
	}
}

func (m *MFAUseCase) SetupMFA(ctx context.Context, userID uuid.UUID) (string, string, error) {
	mfa, _ := m.mfaRepository.FindByUserID(ctx, userID)
	if mfa != nil {
		return "", "", errors.New("user Already has MFA configurations")
	}

	newMfa := entities.NewMFA(userID)
	newMfa.GenerateBackupCodes(8)
	err := mfa.GenerateSecret()
	if err != nil {
		return "", "", err
	}

	err = m.mfaRepository.Create(ctx, newMfa)
	if err != nil {
		return "", "", err
	}

	return newMfa.Secret, "QRPathPlaceHolder", nil
}

// TODO: Centralize Creation of Access and Refresh Token
func (m *MFAUseCase) VerifyAndEnableMFA(ctx context.Context, userID uuid.UUID, code string) (*input.TokenDetails, error) {
	// To Be Implemented
	claims, _ := m.tokenService.VerifyToken(code)

	refresh, access, err := m.tokenService.GenerateTokens(claims.UserID, claims.Email, claims.Role)
	if err != nil {
		return nil, err
	}

	expirationDate, err := m.tokenService.GetTokenExpirationDate(code)
	if err != nil {
		return nil, err
	}

	return &input.TokenDetails{
		RefreshToken: refresh,
		AccessToken:  access,
		ExpiresAt:    expirationDate,
	}, nil
}

func (m *MFAUseCase) DisableMFA(ctx context.Context, userID uuid.UUID, code string) error {
	err := m.mfaRepository.Delete(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

func (m *MFAUseCase) VerifyMFA(ctx context.Context, userID uuid.UUID, code string) error {
	return nil
}

func (m *MFAUseCase) GetMFA(ctx context.Context, userID uuid.UUID) (*entities.MFA, error) {
	mfa, err := m.mfaRepository.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return mfa, nil
}
