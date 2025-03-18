package repository

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/models"
)

type CompletedLessonRepository interface {
	Create(ctx context.Context, completedLesson *models.CompletedLesson) error
	GetByEnrollmentAndLesson(ctx context.Context, enrollmentID, lessonID uint) (*models.CompletedLesson, error)
	ListByEnrollment(ctx context.Context, enrollmentID uint) ([]models.CompletedLesson, error)
	Delete(ctx context.Context, id uint) error
}
