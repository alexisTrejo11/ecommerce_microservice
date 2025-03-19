package services

import (
	"context"
	"math"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/repository"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
	mapper "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/mappers"
	"github.com/google/uuid"
)

type ProgressServiceImpl struct {
	repository repository.ProgressRepository
}

func NewProgressService(repository repository.ProgressRepository) ProgressService {
	return &ProgressServiceImpl{
		repository: repository,
	}
}

func (s *ProgressServiceImpl) GetCourseProgress(ctx context.Context, enrollmentID uuid.UUID) ([]dtos.CompletedLessonDTO, error) {
	completedLessons, err := s.repository.GetByEnrollment(ctx, enrollmentID)
	if err != nil {
		return nil, err
	}

	return mapper.ToCompletedLessonDTOs(completedLessons), nil
}

func (s *ProgressServiceImpl) MarkLessonComplete(ctx context.Context, enrollmentID, lessonID uuid.UUID) error {
	progress, err := s.repository.GetByEnrollmentAndLesson(ctx, enrollmentID, lessonID)
	if err != nil {
		return err
	}

	progress.MarkAsCompleted()

	if err := s.repository.Save(ctx, progress); err != nil {
		return err
	}

	return nil
}

func (s *ProgressServiceImpl) MarkLessonIncomplete(ctx context.Context, enrollmentID, lessonID uuid.UUID) error {
	lessonProgress, err := s.repository.GetByEnrollmentAndLesson(ctx, enrollmentID, lessonID)
	if err != nil {
		return err
	}

	lessonProgress.MarkAsIncomplete()

	if err := s.repository.Save(ctx, lessonProgress); err != nil {
		return err
	}

	return nil
}

func (s *ProgressServiceImpl) CalculateProgress(ctx context.Context, enrollmentID uuid.UUID) (float64, error) {
	progressList, err := s.repository.GetByEnrollment(ctx, enrollmentID)
	if err != nil {
		return 0, nil
	}

	numberOfLesson := len(progressList)
	lessonsCompleted := 0

	if numberOfLesson == 0 {
		return 0, nil
	}

	for _, lesson := range progressList {
		if lesson.IsCompleted {
			lessonsCompleted++
		}
	}

	progressPercentage := (float64(lessonsCompleted) / float64(numberOfLesson)) * 100
	percentageRounded := math.Round(progressPercentage*100) / 100

	return percentageRounded, nil
}

func (s *ProgressServiceImpl) IsLessonCompleted(ctx context.Context, enrollmentID, lessonID uuid.UUID) (error, bool) {
	lessonProgress, err := s.repository.GetByEnrollmentAndLesson(ctx, enrollmentID, lessonID)
	if err != nil {
		return nil, false
	}

	return nil, lessonProgress.IsCompleted
}
