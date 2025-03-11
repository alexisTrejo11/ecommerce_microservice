package usecase

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/google/uuid"
)

type LessonUseCaseImpl struct {
	LessonRepository output.LessonRepository
	mappers          mappers.LessonMappers
}

func NewLessonUseCase(LessonRepository output.LessonRepository) input.LessonUseCase {
	return &LessonUseCaseImpl{
		LessonRepository: LessonRepository,
	}
}

func (us *LessonUseCaseImpl) GetLessonById(ctx context.Context, id uuid.UUID) (*dtos.LessonDTO, error) {
	Lesson, err := us.LessonRepository.GetById(ctx, id.String())
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*Lesson), nil
}

// TODO: Add Buisness logic
func (us *LessonUseCaseImpl) CreateLesson(ctx context.Context, insertDTO dtos.LessonInsertDTO) (*dtos.LessonDTO, error) {
	domain := us.mappers.InsertDTOToDomain(insertDTO)

	domainCreated, err := us.LessonRepository.Create(ctx, *domain)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*domainCreated), nil
}

// TODO Implement Correct Update
func (us *LessonUseCaseImpl) UpdateLesson(ctx context.Context, id uuid.UUID, insertDTO dtos.LessonInsertDTO) (*dtos.LessonDTO, error) {
	domain := us.mappers.InsertDTOToDomain(insertDTO)
	domain.ID = id

	domainCreated, err := us.LessonRepository.Update(ctx, id, *domain)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*domainCreated), nil
}

func (us *LessonUseCaseImpl) DeleteLesson(ctx context.Context, id uuid.UUID) error {
	if err := us.LessonRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
