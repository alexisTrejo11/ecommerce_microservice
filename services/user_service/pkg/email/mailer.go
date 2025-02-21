package email

import (
	"crypto/tls"
	"fmt"

	"github.com/go-gomail/gomail"
)

type MailClient struct {
	dialer *gomail.Dialer
	from   string
}

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUser     string
	SMTPPassword string
	FromEmail    string
}

func NewMailClient(cfg EmailConfig) *MailClient {
	d := gomail.NewDialer(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUser, cfg.SMTPPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return &MailClient{
		dialer: d,
		from:   cfg.FromEmail,
	}
}

func (mc *MailClient) SendHTML(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", mc.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if err := mc.dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}
	return nil
}
