package repository

import (
	"context"

	progress "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProgressRepositoryImpl struct {
	db *gorm.DB
}

func NewProgressRepository(db *gorm.DB) ProgressRepository {
	return &ProgressRepositoryImpl{db: db}
}

func (r *ProgressRepositoryImpl) Save(ctx context.Context, completedLesson *progress.CompletedLesson) error {
	if err := r.db.WithContext(ctx).Save(&completedLesson).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProgressRepositoryImpl) GetByEnrollmentAndLesson(ctx context.Context, enrollmentID, lessonID uuid.UUID) (*progress.CompletedLesson, error) {
	var completeLesson progress.CompletedLesson
	r.db.WithContext(ctx).Where("enrollment id = ? AND lesson_id", enrollmentID, lessonID).First(&completeLesson)

	return &completeLesson, nil
}

func (r *ProgressRepositoryImpl) GetByEnrollment(ctx context.Context, enrollmentID uuid.UUID) ([]progress.CompletedLesson, error) {
	var completeLesson []progress.CompletedLesson
	r.db.WithContext(ctx).Where("enrollment id = ?", enrollmentID).Find(&completeLesson)

	return completeLesson, nil
}

func (r *ProgressRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	var completeLesson progress.CompletedLesson
	if err := r.db.WithContext(ctx).Where("id id = ?", id).First(&completeLesson).Error; err != nil {
		return err
	}

	if err := r.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
