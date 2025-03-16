package output

import (
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/domain"
	"github.com/google/uuid"
)

type ReviewRepository interface {
	Save(review *domain.Review) error
	GetById(id uuid.UUID) (*domain.Review, error)
	GetByCourseId(id uuid.UUID) (*[]domain.Review, error)
	GetByUserId(id uuid.UUID) (*domain.Review, error)
	DeleteById(id uuid.UUID) error
}
