package services

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/models"
)

type ProgressService interface {
	MarkLessonComplete(ctx context.Context, enrollmentID, lessonID uint) error
	MarkLessonIncomplete(ctx context.Context, enrollmentID, lessonID uint) error
	GetCompletedLessons(ctx context.Context, enrollmentID uint) ([]models.CompletedLesson, error)
	CalculateProgress(ctx context.Context, enrollmentID uint) (float64, error)
	IsLessonCompleted(ctx context.Context, enrollmentID, lessonID uint) (bool, error)
}
