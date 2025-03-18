package services

import (
	"context"
	"math"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/repository"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
	mapper "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/mappers"
)

type ProgressServiceImpl struct {
	repository repository.ProgressRepository
}

func NewProgressService(repository repository.ProgressRepository) ProgressService {
	return &ProgressServiceImpl{
		repository: repository,
	}
}

func (s *ProgressServiceImpl) GetCompletedLessons(ctx context.Context, enrollmentID uint) ([]dtos.CompletedLessonDTO, error) {
	completedLessons, err := s.repository.GetByEnrollment(ctx, enrollmentID)
	if err != nil {
		return nil, err
	}

	return mapper.ToCompletedLessonDTOs(completedLessons), nil
}

func (s *ProgressServiceImpl) MarkLessonComplete(ctx context.Context, enrollmentID, lessonID uint) error {
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

func (s *ProgressServiceImpl) MarkLessonIncomplete(ctx context.Context, enrollmentID, lessonID uint) error {
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

func (s *ProgressServiceImpl) CalculateProgress(ctx context.Context, enrollmentID uint) (error, float64) {
	progressList, err := s.repository.GetByEnrollment(ctx, enrollmentID)
	if err != nil {
		return nil, 0
	}

	numberOfLesson := len(progressList)
	lessonsCompleted := 0

	if numberOfLesson == 0 {
		return nil, 0
	}

	for _, lesson := range progressList {
		if lesson.IsCompleted {
			lessonsCompleted++
		}
	}

	progressPercentage := (float64(lessonsCompleted) / float64(numberOfLesson)) * 100
	percentageRounded := math.Round(progressPercentage*100) / 100

	return nil, percentageRounded
}

func (s *ProgressServiceImpl) IsLessonCompleted(ctx context.Context, enrollmentID, lessonID uint) (error, bool) {
	lessonProgress, err := s.repository.GetByEnrollmentAndLesson(ctx, enrollmentID, lessonID)
	if err != nil {
		return nil, false
	}

	return nil, lessonProgress.IsCompleted
}
