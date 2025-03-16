package repository

import "fmt"

var (
	ErrNotificationNotFound = NewRepositoryError("NOTIFICATION_NOT_FOUND", "Notification not found")
	ErrDatabaseFailure      = NewRepositoryError("DATABASE_ERROR", "An error occurred while accessing the database")
)

type RepositoryError struct {
	Code    string
	Message string
}

func (e *RepositoryError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func NewRepositoryError(code, message string) *RepositoryError {
	return &RepositoryError{Code: code, Message: message}
}
