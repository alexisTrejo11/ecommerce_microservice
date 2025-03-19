package repository

import (
	"context"

	enrollment "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/model"
	"github.com/google/uuid"
)

type EnrollmentRepository interface {
	Create(ctx context.Context, enrollment *enrollment.Enrollment) error
	GetByID(ctx context.Context, id uuid.UUID) (*enrollment.Enrollment, error)
	GetByUserAndCourse(ctx context.Context, userID, courseID uuid.UUID) (*enrollment.Enrollment, error)
	GetByUser(ctx context.Context, userID uuid.UUID, page, limit int) ([]enrollment.Enrollment, int64, error)
	GetByCourse(ctx context.Context, courseID uuid.UUID, page, limit int) ([]enrollment.Enrollment, int64, error)
	Update(ctx context.Context, enrollment *enrollment.Enrollment) error
	Delete(ctx context.Context, id uuid.UUID) error
}
