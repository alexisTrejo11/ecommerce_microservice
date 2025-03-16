package domain

import "fmt"

var (
	ErrUserIDRequired         = NewNotificationError("NOTIFICATION_INVALID_USER_ID", "User ID cannot be empty")
	ErrTitleRequired          = NewNotificationError("NOTIFICATION_INVALID_TITLE", "Title cannot be empty")
	ErrContentRequired        = NewNotificationError("NOTIFICATION_INVALID_CONTENT", "Content cannot be empty")
	ErrCannotSchedulePast     = NewNotificationError("NOTIFICATION_INVALID_SCHEDULE", "Cannot schedule a notification for the past")
	ErrCannotCancelNonPending = NewNotificationError("NOTIFICATION_INVALID_CANCEL", "Only pending notifications can be cancelled")
)

type NotificationError struct {
	Code    string
	Message string
}

func (e *NotificationError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func NewNotificationError(code string, message string) *NotificationError {
	return &NotificationError{Code: code, Message: message}
}
