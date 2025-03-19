package services

import (
	"context"

	enrollment "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/model"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/repository"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
	mapper "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/mappers"
	"github.com/google/uuid"
)

type EnrollmentServiceImpl struct {
	repository repository.EnrollmentRepository
}

func NewEnrollmentService(repository repository.EnrollmentRepository) EnrollmentService {
	return &EnrollmentServiceImpl{
		repository: repository,
	}
}

func (r *EnrollmentServiceImpl) GetEnrollmentByID(ctx context.Context, id uuid.UUID) (*dtos.EnrollmentDTO, error) {
	enrollment, err := r.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	enrollmentDTO := mapper.ToEnrollmentDTO(*enrollment)
	return &enrollmentDTO, nil
}
func (r *EnrollmentServiceImpl) GetEnrollmentByUserAndCourse(ctx context.Context, userID, courseID uuid.UUID) (*dtos.EnrollmentDTO, error) {
	enrollment, err := r.repository.GetByUserAndCourse(ctx, userID, courseID)
	if err != nil {
		return nil, err
	}

	enrollmentDTO := mapper.ToEnrollmentDTO(*enrollment)
	return &enrollmentDTO, nil
}

func (r *EnrollmentServiceImpl) GetUserEnrollments(ctx context.Context, userID uuid.UUID, page, limit int) ([]dtos.EnrollmentDTO, int64, error) {
	enrollments, total, err := r.repository.GetByUser(ctx, userID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	return mapper.ToEnrollmentDTOs(enrollments), total, nil
}

func (r *EnrollmentServiceImpl) GetCourseEnrollments(ctx context.Context, courseID uuid.UUID, page, limit int) ([]dtos.EnrollmentDTO, int64, error) {
	enrollments, total, err := r.repository.GetByCourse(ctx, courseID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	return mapper.ToEnrollmentDTOs(enrollments), total, nil
}
func (r *EnrollmentServiceImpl) EnrollUserInCourse(ctx context.Context, userID, courseID uuid.UUID) (*dtos.EnrollmentDTO, error) {
	newEnrollment := enrollment.NewEnrollment(userID, courseID)
	if err := r.repository.Create(ctx, newEnrollment); err != nil {
		return nil, err
	}

	enrollmentDTO := mapper.ToEnrollmentDTO(*newEnrollment)
	return &enrollmentDTO, nil
}

func (r *EnrollmentServiceImpl) CancelEnrollment(ctx context.Context, enrollmentID uuid.UUID) error {
	if err := r.repository.Delete(ctx, enrollmentID); err != nil {
		return err
	}

	return nil
}

func (r *EnrollmentServiceImpl) MarkEnrollmentComplete(ctx context.Context, enrollmentID uuid.UUID) error {
	existingEnrollment, err := r.repository.GetByID(ctx, enrollmentID)
	if err != nil {
		return err
	}

	existingEnrollment.MarkAsCompleted()

	if err := r.repository.Update(ctx, existingEnrollment); err != nil {
		return err
	}

	return nil
}

// Error is Not found
func (r *EnrollmentServiceImpl) IsUserEnrolledInCourse(ctx context.Context, userID, courseID uuid.UUID) bool {
	enrollment, err := r.repository.GetByUserAndCourse(ctx, userID, courseID)
	return enrollment == nil && err != nil
}

func UpdateEnrollmentProgress() {

}
