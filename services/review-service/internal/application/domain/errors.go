package domain

import "fmt"

var (
	ErrRatingValue   = NewReviewError("INVALID_RATING", "Rating must be between 1 and 5")
	ErrCommentLenght = NewReviewError("COMMENT_TOO_LONG", "Comment length exceeds the maximum allowed limit of 1000 characters")
)

type ReviewError struct {
	Code    string
	Message string
}

func (e *ReviewError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func NewReviewError(code string, message string) *ReviewError {
	return &ReviewError{Code: code, Message: message}
}
