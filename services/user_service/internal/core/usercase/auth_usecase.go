package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/api/dto"
	repository "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	userRepo     output.IUserRepository
	tokenService output.ITokenService
	//sessionRepo       output.SessionRepository
	//mfaRepo           output.MFARepository
	//passwordResetRepo output.PasswordResetRepository
	//emailService      output.EmailService
	//config *Config
}

type Config struct {
	AccessTokenExpiry  int
	RefreshTokenExpiry int
	FrontendURL        string
	TOTPIssuer         string
}

func NewAuthUseCase(
	userRepo output.IUserRepository,
	tokenService output.ITokenService,
	//sessionRepo output.SessionRepository,
	//mfaRepo output.MFARepository,
	//passwordResetRepo output.PasswordResetRepository,
	//emailService output.EmailService,
	//config *Config,
) input.AuthUseCase {
	return &AuthUseCase{
		userRepo:     userRepo,
		tokenService: tokenService,
		//sessionRepo:       sessionRepo,
		//mfaRepo:           mfaRepo,
		//passwordResetRepo: passwordResetRepo,
		//emailService:      emailService,
		//config: config,
	}
}

func (uc *AuthUseCase) Register(ctx context.Context, signupDto dto.SignupDTO) (*entities.User, error) {
	// 1. Validar que el DTO no esté vacío
	if signupDto.Email == "" || signupDto.Username == "" || signupDto.Password == "" {
		return nil, errors.New("email, username, and password are required")
	}

	// 2. Verificar si el email ya está registrado
	existingUser, err := uc.userRepo.FindByEmail(ctx, signupDto.Email)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return nil, fmt.Errorf("error checking email availability: %w", err)
	}
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// 3. Verificar si el username ya está en uso
	existingUser, err = uc.userRepo.FindByUsername(ctx, signupDto.Username)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return nil, fmt.Errorf("error checking username availability: %w", err)
	}
	if existingUser != nil {
		return nil, errors.New("username already taken")
	}

	// 4. Validar la fortaleza de la contraseña
	if err := entities.ValidatePasswordStrength(signupDto.Password); err != nil {
		return nil, fmt.Errorf("weak password: %w", err)
	}

	newUser, err := entities.CreateUser(
		signupDto.Email,
		signupDto.Username,
		signupDto.Password,
		2,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	// 6. Validar la entidad User completa
	if err := newUser.Validate(); err != nil {
		return nil, fmt.Errorf("invalid user data: %w", err)
	}

	// 7. Persistir el usuario en la base de datos
	if err := uc.userRepo.Create(ctx, newUser); err != nil {
		return nil, fmt.Errorf("error saving user: %w", err)
	}

	return newUser, nil
}

func (uc *AuthUseCase) Login(ctx context.Context, loginDTO dto.LoginDTO) (*input.TokenDetails, error) {
	existingUser, err := uc.userRepo.FindByEmail(ctx, loginDTO.Email)
	if err != nil {
		return nil, errors.New("no user found with given credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(loginDTO.Password))
	if err != nil {
		return nil, errors.New("no user found with given credentials")
	}

	access_token, refreshToken, err := uc.tokenService.GenerateTokens(existingUser.ID, existingUser.Email, existingUser.Username)
	if err != nil {
		return nil, err
	}

	token_details := &input.TokenDetails{
		RefreshToken: refreshToken,
		AccessToken:  access_token,
	}

	return token_details, nil
}
