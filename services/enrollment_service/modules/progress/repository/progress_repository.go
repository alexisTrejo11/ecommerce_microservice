package repository

import (
	"context"

	progress "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/model"
)

type ProgressRepository interface {
	Save(ctx context.Context, completedLesson *progress.CompletedLesson) error
	GetByEnrollmentAndLesson(ctx context.Context, enrollmentID, lessonID uint) (*progress.CompletedLesson, error)
	GetByEnrollment(ctx context.Context, enrollmentID uint) ([]progress.CompletedLesson, error)
	Delete(ctx context.Context, id uint) error
}
