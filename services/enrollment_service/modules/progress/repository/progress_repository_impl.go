package repository

import (
	"context"
	"errors"

	progress "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/model"
	appErr "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/error"
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
		return appErr.ErrDB
	}
	return nil
}

func (r *ProgressRepositoryImpl) BulkCreate(ctx context.Context, completedLessons *[]progress.CompletedLesson) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&completedLessons).Error; err != nil {
		tx.Rollback()
		return appErr.ErrDB
	}

	return tx.Commit().Error
}

func (r *ProgressRepositoryImpl) GetByEnrollmentAndLesson(ctx context.Context, enrollmentID, lessonID uuid.UUID) (*progress.CompletedLesson, error) {
	var completeLesson progress.CompletedLesson
	if err := r.db.WithContext(ctx).Where("enrollment_id = ? AND lesson_id = ?", enrollmentID, lessonID).First(&completeLesson).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErr.ErrProgressNotFoundDB
		}
		return nil, appErr.ErrDB
	}
	return &completeLesson, nil
}

func (r *ProgressRepositoryImpl) GetByEnrollment(ctx context.Context, enrollmentID uuid.UUID) ([]progress.CompletedLesson, error) {
	var completeLessons []progress.CompletedLesson
	if err := r.db.WithContext(ctx).Where("enrollment_id = ?", enrollmentID).Find(&completeLessons).Error; err != nil {
		return nil, appErr.ErrDB
	}
	return completeLessons, nil
}

func (r *ProgressRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	var completeLesson progress.CompletedLesson
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&completeLesson).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return appErr.ErrProgressNotFoundDB
		}
		return appErr.ErrDB
	}

	if err := r.db.WithContext(ctx).Delete(&completeLesson).Error; err != nil {
		return appErr.ErrDB
	}
	return nil
}
