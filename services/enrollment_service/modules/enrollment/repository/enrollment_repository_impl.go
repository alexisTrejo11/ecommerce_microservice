package repository

import (
	"context"
	"fmt"

	enrollment "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EnrollmentRepositoryImpl struct {
	db *gorm.DB
}

func NewEnrollmentRepository(db *gorm.DB) EnrollmentRepository {
	return &EnrollmentRepositoryImpl{db: db}
}

func (r *EnrollmentRepositoryImpl) Create(ctx context.Context, enrollment *enrollment.Enrollment) error {
	if err := r.db.WithContext(ctx).Create(&enrollment).Error; err != nil {
		return err
	}

	return nil
}

func (r *EnrollmentRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*enrollment.Enrollment, error) {
	var enrollment enrollment.Enrollment
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&enrollment).Error; err != nil {
		return nil, err
	}

	fmt.Println(enrollment)

	return &enrollment, nil
}

func (r *EnrollmentRepositoryImpl) GetByUserAndCourse(ctx context.Context, userID, courseID uuid.UUID) (*enrollment.Enrollment, error) {
	var enrollment enrollment.Enrollment
	if err := r.db.WithContext(ctx).Where("user_id = ? AND course_id = ?", userID, courseID).First(&enrollment).Error; err != nil {
		return nil, err
	}

	return &enrollment, nil
}

func (r *EnrollmentRepositoryImpl) Update(ctx context.Context, enrollment *enrollment.Enrollment) error {
	if err := r.db.WithContext(ctx).Save(enrollment).Error; err != nil {
		return err
	}

	return nil
}

func (r *EnrollmentRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	var enrollment enrollment.Enrollment

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&enrollment).Error; err != nil {
		return err
	}

	err := r.db.WithContext(ctx).Delete(&enrollment).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *EnrollmentRepositoryImpl) GetByUser(ctx context.Context, userID uuid.UUID, page, limit int) ([]enrollment.Enrollment, int64, error) {
	var enrollments []enrollment.Enrollment
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&enrollments).Error; err != nil {
		return nil, 0, err
	}

	return enrollments, int64(len(enrollments)), nil

}

func (r *EnrollmentRepositoryImpl) GetByCourse(ctx context.Context, courseID uuid.UUID, page, limit int) ([]enrollment.Enrollment, int64, error) {
	var enrollments []enrollment.Enrollment
	if err := r.db.WithContext(ctx).Where("course_id = ?", courseID).Find(&enrollments).Error; err != nil {
		return nil, 0, err
	}

	return enrollments, int64(len(enrollments)), nil
}
