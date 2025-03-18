package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/port/input"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/application/port/output"
	rabbitmq "github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/message"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/port/output/repository"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/pkg/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/pkg/mapper"
	"github.com/google/uuid"
)

type ReviewUseCaseImpl struct {
	repository    output.ReviewRepository
	messageClient *rabbitmq.RabbitMQClient
	mapper        mapper.ReviewMapper
}

func NewReviewUseCase(repository output.ReviewRepository, messageClient *rabbitmq.RabbitMQClient) input.ReviewUseCase {
	return &ReviewUseCaseImpl{
		repository:    repository,
		messageClient: messageClient,
		mapper:        mapper.ReviewMapper{},
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

func (uc *ReviewUseCaseImpl) CreateReview(ctx context.Context, userID uuid.UUID, insertDTO dtos.ReviewInsertDTO) (*dtos.ReviewDTO, error) {
	if err := uc.validateNotDuplicatedReview(ctx, insertDTO.CourseID, userID); err != nil {
		return nil, err
	}

	if err := uc.validateUserEnrollment(ctx, insertDTO.CourseID, userID); err != nil {
		return nil, err
	}

	review, err := uc.mapper.InsertDTOToDomain(insertDTO, userID)
	if err != nil {
		return nil, err
	}

	err = uc.repository.Save(ctx, review)
	if err != nil {
		return nil, err
	}

	go uc.updateCourseAvgRating(ctx, review.GetCourseID())

	return uc.mapper.DomainToDTO(review), nil
}

func (uc *ReviewUseCaseImpl) UpdateReview(ctx context.Context, userID uuid.UUID, id uuid.UUID, insertDTO dtos.ReviewInsertDTO) (*dtos.ReviewDTO, error) {
	existingReview, err := uc.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := uc.validateUserEnrollment(ctx, insertDTO.CourseID, userID); err != nil {
		return nil, err
	}

	err = existingReview.Update(
		id,
		userID,
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

	go uc.updateCourseAvgRating(ctx, existingReview.GetCourseID())

	return uc.mapper.DomainToDTO(existingReview), nil
}

func (uc *ReviewUseCaseImpl) DeleteReview(ctx context.Context, userID, id uuid.UUID) error {
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
	go uc.updateCourseAvgRating(ctx, courseID)

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

func (uc *ReviewUseCaseImpl) updateCourseAvgRating(ctx context.Context, courseID uuid.UUID) {
	reviews, err := uc.repository.GetByCourseID(ctx, courseID)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	rating := domain.CalculateRating(*reviews)
	uc.messageClient.PublishCourseRatingUpdate(courseID.String(), rating)

	fmt.Printf("rating: send it to queue%v\n", rating)

}
