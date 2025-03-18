package input

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/pkg/dtos"
	"github.com/google/uuid"
)

type ReviewUseCase interface {
	GetReviewById(ctx context.Context, id uuid.UUID) (*dtos.ReviewDTO, error)
	GetReviewsByUserId(ctx context.Context, userID uuid.UUID) (*[]dtos.ReviewDTO, error)
	GetReviewsByCourseId(ctx context.Context, course uuid.UUID) (*[]dtos.ReviewDTO, error)
	CreateReview(ctx context.Context, userID uuid.UUID, insertDTO dtos.ReviewInsertDTO) (*dtos.ReviewDTO, error)
	UpdateReview(ctx context.Context, userID uuid.UUID, id uuid.UUID, insertDTO dtos.ReviewInsertDTO) (*dtos.ReviewDTO, error)
	DeleteReview(ctx context.Context, userID uuid.UUID, id uuid.UUID) error

	GetCourseRating(ctx context.Context, course uuid.UUID) (float64, error)
	UpdateCourseReviewData(ctx context.Context, course uuid.UUID) (*dtos.ReviewDTO, error)
}
