package output

import (
	"context"
)

type IEmailService interface {
	SendVerificationEmail(ctx context.Context, email, name, verificationURL string) error
	SendPasswordResetEmail(ctx context.Context, email, name, resetURL string) error
	SendMFASetupEmail(ctx context.Context, email, name string, backupCodes []string) error
	SendLoginNotificationEmail(ctx context.Context, email, name, time, location, device string) error
}
