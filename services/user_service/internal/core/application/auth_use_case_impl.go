package usecases

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/http/v1/dto"
	repository "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/tokens"
	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	userRepo      output.UserRepository
	tokenService  output.TokenService
	userMappers   mappers.UserMappers
	sessionRepo   output.SessionRepository
	tokenFactory  tokens.TokenFactory
	mfaRepository output.MFARepository
}

type Config struct {
	AccessTokenExpiry  int
	RefreshTokenExpiry int
	FrontendURL        string
	TOTPIssuer         string
}

func NewAuthUseCase(
	userRepo output.UserRepository,
	tokenService output.TokenService,
	sessionRepo output.SessionRepository,
	mfaRepository output.MFARepository,
) input.AuthUseCase {
	return &AuthUseCase{
		userRepo:      userRepo,
		tokenService:  tokenService,
		sessionRepo:   sessionRepo,
		tokenFactory:  *tokens.NewTokenFactory(),
		mfaRepository: mfaRepository,
	}
}

func (uc *AuthUseCase) Register(ctx context.Context, signupDto dto.SignupDTO) (*entities.User, string, error) {
	if err := validateSignupDTO(signupDto); err != nil {
		return nil, "", err
	}

	if err := uc.isEmailAvailable(ctx, signupDto.Email); err != nil {
		return nil, "", err
	}

	if err := uc.isUsernameAvailable(ctx, signupDto.Username); err != nil {
		return nil, "", err
	}

	if err := validatePasswordStrength(signupDto.Password); err != nil {
		return nil, "", err
	}

	newUser, err := uc.createUserEntity(signupDto, 2)
	if err != nil {
		return nil, "", err
	}

	if err := uc.saveUser(ctx, newUser); err != nil {
		return nil, "", err
	}

	activationToken, err := uc.generateActivationToken(newUser)
	if err != nil {
		return nil, "", err
	}

	return newUser, activationToken, nil
}

func (uc *AuthUseCase) Login(ctx context.Context, loginDTO dto.LoginDTO) (*input.TokenDetails, error) {
	user, err := uc.verifyCredentials(ctx, loginDTO.Email, loginDTO.Password)
	if err != nil {
		return nil, err
	}

	if err := user.VerifyValidAccount(); err != nil {
		return nil, err
	}

	if err := uc.checkMFA(ctx, user.ID); err != nil {
		return nil, err
	}

	tokenDetails, err := uc.generateTokens(user.ID, user.Email, user.Username)
	if err != nil {
		return nil, err
	}

	session, err := uc.createSession(ctx, *tokenDetails, *user)
	if err != nil {
		return nil, err
	}

	tokenDetails.SessionID = session.ID
	return tokenDetails, nil
}

func (uc *AuthUseCase) Logout(ctx context.Context, refreshToken string, userID uuid.UUID) error {
	session, err := uc.sessionRepo.FindByRefreshToken(ctx, refreshToken)
	if err != nil {
		return err
	}

	if session.UserID != userID {
		return errors.New("not allowed to get this data")
	}

	if _, err := uc.tokenService.VerifyToken(refreshToken, tokens.RefreshTokenENUM); err != nil {
		return err
	}

	return uc.sessionRepo.Delete(ctx, session.ID)
}

func (uc *AuthUseCase) LogoutAll(ctx context.Context, userID uuid.UUID) error {
	return uc.sessionRepo.DeleteAllByUserID(ctx, userID)
}

func (uc *AuthUseCase) RefreshTokens(ctx context.Context, refreshToken, userAgent, clientIP string) (*input.TokenDetails, error) {
	session, err := uc.sessionRepo.FindByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	claims, err := uc.tokenService.VerifyToken(refreshToken, tokens.RefreshTokenENUM)
	if err != nil {
		return nil, err
	}

	accessToken, err := uc.tokenService.RefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	return &input.TokenDetails{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		ExpiresAt:    claims.ExpiresAt,
		SessionID:    session.ID,
	}, nil
}

func (uc *AuthUseCase) ResetPassword(ctx context.Context, token, newPassword string) error {
	claims, err := uc.tokenService.VerifyToken(token, tokens.VerifyTokenENUM)
	if err != nil {
		return err
	}

	if err := validatePasswordStrength(newPassword); err != nil {
		return err
	}

	user, err := uc.userRepo.FindByEmail(ctx, claims.Email)
	if err != nil {
		return err
	}

	newHashedPassword, err := user.HashPassword(newPassword)
	if err != nil {
		return err
	}

	user.PasswordHash = newHashedPassword
	return uc.userRepo.Update(ctx, user)
}

func (uc *AuthUseCase) ActivateAccount(ctx context.Context, token string) error {
	claims, err := uc.tokenService.VerifyToken(token, tokens.VerifyTokenENUM)
	if err != nil {
		return err
	}

	user, err := uc.userRepo.FindByEmail(ctx, claims.Email)
	if err != nil {
		return err
	}

	user.ActivateAccount()
	return uc.userRepo.Update(ctx, user)
}

func (uc *AuthUseCase) ResendCode(ctx context.Context, codeType string, userID uuid.UUID) error {
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	switch codeType {
	case "activation", "password_reset":
		_, _, err := uc.tokenService.GenerateTokens(user.ID, user.Email, user.Role.Name)
		return err
	default:
		return errors.New("invalid code type")
	}
}

func validateSignupDTO(signupDto dto.SignupDTO) error {
	if signupDto.Email == "" || signupDto.Username == "" || signupDto.Password == "" {
		return errors.New("email, username, and password are required")
	}
	return nil
}

func (uc *AuthUseCase) isEmailAvailable(ctx context.Context, email string) error {
	existingUser, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return fmt.Errorf("error checking email availability: %w", err)
	}
	if existingUser != nil {
		return errors.New("email already registered")
	}
	return nil
}

func (uc *AuthUseCase) isUsernameAvailable(ctx context.Context, username string) error {
	existingUser, err := uc.userRepo.FindByUsername(ctx, username)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return fmt.Errorf("error checking username availability: %w", err)
	}
	if existingUser != nil {
		return errors.New("username already taken")
	}
	return nil
}

func validatePasswordStrength(password string) error {
	if err := entities.ValidatePasswordStrength(password); err != nil {
		return fmt.Errorf("weak password: %w", err)
	}
	return nil
}

func (uc *AuthUseCase) createUserEntity(signupDTO dto.SignupDTO, roleID int) (*entities.User, error) {
	newUser := uc.userMappers.SignupDTOToDomain(signupDTO)
	newUser.RoleID = uint(roleID)
	newUser.ID = uuid.NewString()
	newUser.Status = entities.UserStatusPending

	if err := newUser.Validate(); err != nil {
		return nil, fmt.Errorf("invalid user data: %w", err)
	}

	hashedPassword, err := newUser.HashPassword(signupDTO.Password)
	if err != nil {
		return nil, fmt.Errorf("cant hash password: %w", err)
	}
	newUser.PasswordHash = hashedPassword

	return newUser, nil
}

func (uc *AuthUseCase) saveUser(ctx context.Context, user *entities.User) error {
	return uc.userRepo.Create(ctx, user)
}

func (uc *AuthUseCase) verifyCredentials(ctx context.Context, email, password string) (*entities.User, error) {
	existingUser, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("no user found with given credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("no user found with given credentials")
	}

	return existingUser, nil
}

func (uc *AuthUseCase) generateTokens(userID, email, username string) (*input.TokenDetails, error) {
	accessToken, refreshToken, err := uc.tokenService.GenerateTokens(userID, email, username)
	if err != nil {
		return nil, err
	}

	expirationDate, err := uc.tokenService.GetTokenExpirationDate(refreshToken, tokens.RefreshTokenENUM)
	if err != nil {
		return nil, err
	}

	return &input.TokenDetails{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		ExpiresAt:    expirationDate,
	}, nil
}

func (uc *AuthUseCase) generateActivationToken(user *entities.User) (string, error) {
	factory, err := uc.tokenFactory.CreateToken(tokens.VerifyTokenENUM)
	if err != nil {
		return "", err
	}

	return factory.Generate(user.Email, user.ID, user.Role.Name)
}

func (uc *AuthUseCase) createSession(ctx context.Context, tokens input.TokenDetails, user entities.User) (*entities.Session, error) {
	userId, _ := uuid.Parse(user.ID)
	session := entities.Session{
		ID:           uuid.New(),
		UserID:       userId,
		RefreshToken: tokens.RefreshToken,
		ExpiresAt:    tokens.ExpiresAt,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := uc.sessionRepo.Create(ctx, &session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (uc *AuthUseCase) checkMFA(ctx context.Context, userID string) error {
	userId, _ := uuid.Parse(userID)
	mfa, err := uc.mfaRepository.FindByUserID(ctx, userId)
	if err != nil {
		return err
	}

	if mfa != nil {
		return errors.New("user has mfa activated")
	}

	return nil
}
