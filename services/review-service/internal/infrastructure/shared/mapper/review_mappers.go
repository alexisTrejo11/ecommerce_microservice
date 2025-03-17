package mapper

import (
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/port/output/models"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/shared/dtos"
)

type ReviewMapper struct {
}

func (m *ReviewMapper) InsertDTOToDomain(dto dtos.ReviewInsertDTO) *domain.Review {
	return domain.NewReview(
		dto.UserID,
		dto.CourseID,
		dto.Rating,
		dto.Comment,
	)
}

func (m *ReviewMapper) DomainToDTO(review *domain.Review) *dtos.ReviewDTO {
	dto := &dtos.ReviewDTO{}
	dto.SetID(review.GetID())
	dto.SetUserID(review.GetUserID())
	dto.SetCourseID(review.GetCourseID())
	dto.SetRating(review.GetRating())
	dto.SetComment(review.GetComment())
	dto.SetCreatedAt(review.GetCreatedAt())
	dto.SetUpdatedAt(review.GetUpdatedAt())
	dto.SetIsApproved(review.IsApproved())
	return dto
}
func (m *ReviewMapper) ModelToDomain(model *models.ReviewModel) *domain.Review {
	return domain.NewReviewFromDTO(
		model.ID,
		model.UserID,
		model.CourseID,
		model.Rating,
		model.Comment,
		model.CreatedAt,
		model.UpdatedAt,
		model.IsApproved,
	)
}

func (m *ReviewMapper) DomainToModel(domainReview *domain.Review) *models.ReviewModel {
	return &models.ReviewModel{
		ID:         domainReview.GetID(),
		UserID:     domainReview.GetUserID(),
		CourseID:   domainReview.GetCourseID(),
		Rating:     domainReview.GetRating(),
		Comment:    domainReview.GetComment(),
		CreatedAt:  domainReview.GetCreatedAt(),
		UpdatedAt:  domainReview.GetUpdatedAt(),
		IsApproved: domainReview.IsApproved(),
	}
}

func (m *ReviewMapper) ModelsToDomainList(models []models.ReviewModel) []domain.Review {
	var domainReviews []domain.Review
	for _, model := range models {
		domainReview := m.ModelToDomain(&model)
		domainReviews = append(domainReviews, *domainReview)
	}
	return domainReviews
}

func (m *ReviewMapper) DomainsToDTOList(reviews []domain.Review) []dtos.ReviewDTO {
	dtos := make([]dtos.ReviewDTO, len(reviews))
	for i, review := range reviews {
		dtos[i] = *m.DomainToDTO(&review)
	}
	return dtos
}
