package appErr

import "fmt"

type ApplicationError struct {
	Code    string
	Message string
	Err     error
}

func (e *ApplicationError) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s, Details: %v", e.Code, e.Message, e.Err)
}

func NewApplicationError(code string, message string, err error) *ApplicationError {
	return &ApplicationError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

var (
	ErrNotFoundDB         = NewApplicationError("NOT_FOUND", "The requested entity was not found", nil)
	ErrDB                 = NewApplicationError("DATABASE_ERROR", "An error occurred while accessing the database", nil)
	ErrInvalidOperationDB = NewApplicationError("INVALID_OPERATION", "The requested operation is invalid", nil)

	ErrEnrollmentNotFoundDB  = NewApplicationError("ENROLLENT_NOT_FOUND", "The requested enrollment was not found", nil)
	ErrCertificateNotFoundDB = NewApplicationError("CERTIFICATE_NOT_FOUND", "The requested module was not found", nil)
	ErrProgressNotFoundDB    = NewApplicationError("PROGRESS_NOT_FOUND", "The requested lesson was not found", nil)
)
