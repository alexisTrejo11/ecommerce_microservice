package usecase

import (
	"bytes"
	"context"
	"fmt"
	"html/template"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/pkg/email"
)

type EmailUseCase struct {
	mailClient *email.MailClient
}

func NewEmailUseCase(
	mailClient *email.MailClient) input.EmailUseCase {
	return &EmailUseCase{
		mailClient: mailClient,
	}
}

func (uc *EmailUseCase) SendEmail(ctx context.Context, emailDTO dtos.NotificationMessageDTO) error {
	templateData, err := email.TemplateFS.ReadFile("templates/template.html")
	if err != nil {
		return fmt.Errorf("error reading email template: %w", err)
	}

	tmpl, err := template.New("email").Parse(string(templateData))
	if err != nil {
		return fmt.Errorf("error parsing email template: %w", err)
	}

	var emailBody bytes.Buffer
	if err := tmpl.Execute(&emailBody, emailDTO); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	if err := uc.mailClient.SendHTML(emailDTO.UserData.Email, emailDTO.Title, emailBody.String()); err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}
