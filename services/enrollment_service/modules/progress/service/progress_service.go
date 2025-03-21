package services

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
	"github.com/google/uuid"
)

type ProgressService interface {
	CreateCourseTrackRecord(ctx context.Context, enrollmentID uuid.UUID) error
	MarkLessonComplete(ctx context.Context, enrollmentID, lessonID uuid.UUID) error
	MarkLessonIncomplete(ctx context.Context, enrollmentID, lessonID uuid.UUID) error
	GetCourseProgress(ctx context.Context, enrollmentID uuid.UUID) ([]dtos.CompletedLessonDTO, error)
	CalculateProgress(ctx context.Context, enrollmentID uuid.UUID) (float64, error)
	IsLessonCompleted(ctx context.Context, enrollmentID, lessonID uuid.UUID) (error, bool)
}
