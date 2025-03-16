package usecase

import (
	"bytes"
	"context"
	"html/template"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/pkg/email"
	logging "github.com/alexisTrejo11/ecommerce_microservice/notification-service/pkg/log"
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
	logging.Logger.Info().
		Str("action", "send_email").
		Str("email", emailDTO.UserData.Email).
		Msg("Starting email sending process")

	templateData, err := email.TemplateFS.ReadFile("templates/template.html")
	if err != nil {
		logging.LogError("read_email_template", "Error reading email template", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}

	tmpl, err := template.New("email").Parse(string(templateData))
	if err != nil {
		logging.LogError("parse_email_template", "Error parsing email template", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}

	var emailBody bytes.Buffer
	if err := tmpl.Execute(&emailBody, emailDTO); err != nil {
		logging.LogError("execute_email_template", "Error executing email template", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}

	if err := uc.mailClient.SendHTML(emailDTO.UserData.Email, emailDTO.Title, emailBody.String()); err != nil {
		logging.LogError("send_email", "Error sending email", map[string]interface{}{
			"error": err.Error(),
			"email": emailDTO.UserData.Email,
		})
		return err
	}

	logging.LogSuccess("send_email", "Email sent successfully", map[string]interface{}{
		"email": emailDTO.UserData.Email,
		"title": emailDTO.Title,
	})

	return nil
}
