package usecases

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/email"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/shared/tokens"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	templatePath = "services/user_service/pkg/email/templates"
)

type EmailUseCase struct {
	mailClient   *email.MailClient
	userRepo     output.UserRepository
	tokenService output.TokenService
	frontendURL  string
}

func NewEmailUseCase(mailClient *email.MailClient, userRepo output.UserRepository, frontendURL string, tokenService output.TokenService) input.EmailUseCase {
	return &EmailUseCase{
		mailClient:   mailClient,
		userRepo:     userRepo,
		frontendURL:  frontendURL,
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

func (uc *EmailUseCase) VerifyEmail(ctx context.Context, token string) error {
	claims, err := uc.tokenService.VerifyToken(token, tokens.VerifyTokenENUM)
	if err != nil {
		return err
	}

	user, err := uc.userRepo.FindByEmail(ctx, claims.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("invalid or expired token")
		}
		return fmt.Errorf("error retrieving user by token: %w", err)
	}

	// Centralize this error
	if claims.UserID != user.ID {
		return errors.New("not allowed to get this data")
	}

	user.ActivateAccount()

	if err := uc.userRepo.Update(ctx, user); err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}

	return nil
}

func (uc *EmailUseCase) SendPasswordResetEmail(ctx context.Context, email string) error {
	user, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("error retrieving user by email: %w", err)
	}

	resetToken := uuid.New().String()
	resetLink := fmt.Sprintf("%s/reset-password?token=%s", uc.frontendURL, resetToken)

	// Save Token
	_, _, err = uc.tokenService.GenerateTokens(user.ID, user.Email, user.Role.Name)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(filepath.Join(templatePath, "password_reset.html"))
	if err != nil {
		return fmt.Errorf("error loading template: %w", err)
	}

	var bodyBuffer bytes.Buffer
	data := struct {
		ResetLink string
	}{
		ResetLink: resetLink,
	}

	if err := tmpl.ExecuteTemplate(&bodyBuffer, "password_reset.html", data); err != nil {
		return fmt.Errorf("error rendering template: %w", err)
	}

	if err := uc.mailClient.SendHTML(user.Email, "Reset Your Password", bodyBuffer.String()); err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}
