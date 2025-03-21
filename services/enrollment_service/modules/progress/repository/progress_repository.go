package repository

import (
	"context"

	progress "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/model"
	"github.com/google/uuid"
)

type ProgressRepository interface {
	Save(ctx context.Context, completedLesson *progress.CompletedLesson) error
	BulkCreate(ctx context.Context, completedLessons *[]progress.CompletedLesson) error
	GetByEnrollmentAndLesson(ctx context.Context, enrollmentID, lessonID uuid.UUID) (*progress.CompletedLesson, error)
	GetByEnrollment(ctx context.Context, enrollmentID uuid.UUID) ([]progress.CompletedLesson, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
