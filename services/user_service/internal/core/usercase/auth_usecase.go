package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/api/dto"
	repository "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output"
	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	userRepo     output.IUserRepository
	tokenService output.ITokenService
	userMappers  mappers.UserMappers
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
	if err := validateSignupDTO(signupDto); err != nil {
		return nil, err
	}

	if err := uc.isEmailAvailable(ctx, signupDto.Email); err != nil {
		return nil, err
	}

	if err := uc.isUsernameAvailable(ctx, signupDto.Username); err != nil {
		return nil, err
	}

	if err := validatePasswordStrength(signupDto.Password); err != nil {
		return nil, err
	}

	newUser, err := uc.createUserEntity(signupDto, 2)
	if err != nil {
		return nil, err
	}

	if err := uc.saveUser(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (uc *AuthUseCase) Login(ctx context.Context, loginDTO dto.LoginDTO) (*input.TokenDetails, error) {
	user, err := uc.verifyCredentials(ctx, loginDTO.Email, loginDTO.Password)
	if err != nil {
		return nil, err
	}

	tokenDetails, err := uc.generateTokens(user.ID, user.Email, user.Username)
	if err != nil {
		return nil, err
	}

	return tokenDetails, nil
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

	if err := newUser.Validate(); err != nil {
		return nil, fmt.Errorf("invalid user data: %w", err)
	}

	hashed_passwrod, err := newUser.HashPassword(signupDTO.Password)
	if err != nil {
		return nil, fmt.Errorf("cant hash password: %w", err)
	}
	newUser.PasswordHash = hashed_passwrod

	return newUser, nil
}

func (uc *AuthUseCase) saveUser(ctx context.Context, user *entities.User) error {
	if err := uc.userRepo.Create(ctx, user); err != nil {
		return fmt.Errorf("error saving user: %w", err)
	}
	return nil
}

func (uc *AuthUseCase) verifyCredentials(ctx context.Context, email, password string) (*entities.User, error) {
	existingUser, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("no user found with given credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("no user found with given credentials")
	}

	return existingUser, nil
}

func (uc *AuthUseCase) generateTokens(userID, email, username string) (*input.TokenDetails, error) {
	accessToken, refreshToken, err := uc.tokenService.GenerateTokens(userID, email, username)
	if err != nil {
		return nil, err
	}

	return &input.TokenDetails{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}
