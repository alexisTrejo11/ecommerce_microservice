package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/port/input"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/port/output"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/port/output/repository"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/shared/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/shared/mapper"
	"github.com/google/uuid"
)

type ReviewUseCaseImpl struct {
	repository output.ReviewRepository
	mapper     mapper.ReviewMapper
}

func NewReviewUseCase(repository output.ReviewRepository) input.ReviewUseCase {
	return &ReviewUseCaseImpl{
		repository: repository,
		mapper:     mapper.ReviewMapper{},
	}
}

func (uc *ReviewUseCaseImpl) GetReviewById(ctx context.Context, id uuid.UUID) (*dtos.ReviewDTO, error) {
	review, err := uc.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return uc.mapper.DomainToDTO(review), nil
}

func (uc *ReviewUseCaseImpl) GetReviewsByUserId(ctx context.Context, userID uuid.UUID) (*[]dtos.ReviewDTO, error) {
	reviews, err := uc.repository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	dtos := uc.mapper.DomainsToDTOList(*reviews)
	return &dtos, nil
}

func (uc *ReviewUseCaseImpl) GetReviewsByCourseId(ctx context.Context, courseID uuid.UUID) (*[]dtos.ReviewDTO, error) {
	reviews, err := uc.repository.GetByCourseID(ctx, courseID)
	if err != nil {
		return nil, err
	}

	dtos := uc.mapper.DomainsToDTOList(*reviews)
	return &dtos, nil
}

func (uc *ReviewUseCaseImpl) CreateReview(ctx context.Context, insertDTO dtos.ReviewInsertDTO) (*dtos.ReviewDTO, error) {
	if err := uc.validateNotDuplicatedReview(ctx, insertDTO.CourseID, insertDTO.UserID); err != nil {
		return nil, err
	}

	if err := uc.validateUserEnrollment(ctx, insertDTO.CourseID, insertDTO.UserID); err != nil {
		return nil, err
	}

	review, err := uc.mapper.InsertDTOToDomain(insertDTO)
	if err != nil {
		return nil, err
	}

	err = uc.repository.Save(ctx, review)
	if err != nil {
		return nil, err
	}

	return uc.mapper.DomainToDTO(review), nil
}

func (uc *ReviewUseCaseImpl) UpdateReview(ctx context.Context, id uuid.UUID, insertDTO dtos.ReviewInsertDTO) (*dtos.ReviewDTO, error) {
	existingReview, err := uc.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := uc.validateUserEnrollment(ctx, insertDTO.CourseID, insertDTO.UserID); err != nil {
		return nil, err
	}

	err = existingReview.Update(
		id,
		insertDTO.UserID,
		insertDTO.CourseID,
		insertDTO.Rating,
		insertDTO.Comment,
	)
	if err != nil {
		return nil, err
	}

	err = uc.repository.Save(ctx, existingReview)
	if err != nil {
		return nil, err
	}

	return uc.mapper.DomainToDTO(existingReview), nil
}

func (uc *ReviewUseCaseImpl) DeleteReview(ctx context.Context, id uuid.UUID) error {
	return uc.repository.DeleteByID(ctx, id)
}

func (uc *ReviewUseCaseImpl) GetCourseRating(ctx context.Context, courseID uuid.UUID) (float64, error) {
	reviews, err := uc.repository.GetByCourseID(ctx, courseID)
	if err != nil {
		return 0, err
	}

	rating := domain.CalculateRating(*reviews)
	return rating, nil
}

func (uc *ReviewUseCaseImpl) UpdateCourseReviewData(ctx context.Context, courseID uuid.UUID) (*dtos.ReviewDTO, error) {
	reviews, err := uc.repository.GetByCourseID(ctx, courseID)
	if err != nil {
		return nil, err
	}

	rating := domain.CalculateRating(*reviews)
	fmt.Printf("rating: %v\n", rating)
	// Update In Course Service

	return nil, nil
}

func (uc *ReviewUseCaseImpl) validateNotDuplicatedReview(
	ctx context.Context,
	courseID uuid.UUID,
	userID uuid.UUID) error {

	_, err := uc.repository.GetByCourseIDAndUserID(ctx, courseID, userID)
	if err != nil {
		if errors.Is(err, repository.ErrReviewNotFound) {
			return nil
		}
		return err
	}
	return errors.New("user already has a review on this course")
}

func (uc *ReviewUseCaseImpl) validateUserEnrollment(
	ctx context.Context,
	courseID uuid.UUID,
	userID uuid.UUID) error {

	// To Be Implemented
	fmt.Printf("courseID: %v\n", courseID)
	fmt.Printf("userID: %v\n", userID)
	fmt.Printf("ctx: %v\n", ctx)

	return nil
}
