package usecases

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/email"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EmailUseCase struct {
	mailClient   *email.MailClient
	userRepo     output.UserRepository
	tokenService output.TokenService
}

func NewEmailUseCase(
	mailClient *email.MailClient,
	userRepo output.UserRepository,
	tokenService output.TokenService) input.EmailUseCase {
	return &EmailUseCase{
		mailClient:   mailClient,
		userRepo:     userRepo,
		tokenService: tokenService,
	}
}

func (uc *EmailUseCase) SendVerificationEmail(ctx context.Context, userID uuid.UUID, token string) error {
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}

	templateData, err := email.TemplateFS.ReadFile("templates/verification_email.html")
	if err != nil {
		return fmt.Errorf("error reading email template: %w", err)
	}

	emailBody := strings.Replace(string(templateData), "{{TOKEN}}", token, 1)

	if err := uc.mailClient.SendHTML(user.Email, "Verify Your Email", emailBody); err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}

func (uc *EmailUseCase) SendPasswordResetEmail(ctx context.Context, userID uuid.UUID, token string) error {
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("error retrieving user by ID: %w", err)
	}

	templateData, err := email.TemplateFS.ReadFile("templates/password_reset.html")
	if err != nil {
		return fmt.Errorf("error reading email template: %w", err)
	}

	emailBody := strings.Replace(string(templateData), "{{TOKEN}}", token, 1)

	if err := uc.mailClient.SendHTML(user.Email, "Reset Your Password", emailBody); err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}
