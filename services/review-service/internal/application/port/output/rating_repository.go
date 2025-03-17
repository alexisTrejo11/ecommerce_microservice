package output

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/domain"
	"github.com/google/uuid"
)

type ReviewRepository interface {
	Save(ctx context.Context, review *domain.Review) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Review, error)
	GetByCourseID(ctx context.Context, id uuid.UUID) (*[]domain.Review, error)
	GetByUserID(ctx context.Context, id uuid.UUID) (*[]domain.Review, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
}
