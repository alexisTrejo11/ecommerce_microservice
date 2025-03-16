package input

import (
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/shared/dtos"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type ReviewUseCase interface {
	GetReviewById(ctx context.Context, id uuid.UUID)
	GetReviewByUserId(ctx context.Context, userID uuid.UUID)
	GetReviewByCourseId(ctx context.Context, course uuid.UUID)
	CreateReview(ctx context.Context, insertDTO dtos.ReviewInsertDTO)
	UpateReview(ctx context.Context, insertDTO dtos.ReviewInsertDTO)
	DeleteReview(ctx context.Context, id uuid.UUID)

	GetCourseRating(ctx context.Context, course uuid.UUID)
	UpdateCourseReivewData(ctx context.Context, course uuid.UUID)
}
