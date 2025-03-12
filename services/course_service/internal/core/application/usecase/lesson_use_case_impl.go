package usecase

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/google/uuid"
)

type LessonUseCaseImpl struct {
	lessonRepository output.LessonRepository
	moduleRepository output.ModuleRepository
	mappers          mappers.LessonMappers
}

func NewLessonUseCase(lessonRepository output.LessonRepository, moduleRepository output.ModuleRepository) input.LessonUseCase {
	return &LessonUseCaseImpl{
		lessonRepository: lessonRepository,
		moduleRepository: moduleRepository,
	}
}

func (us *LessonUseCaseImpl) GetLessonById(ctx context.Context, id uuid.UUID) (*dtos.LessonDTO, error) {
	lesson, err := us.lessonRepository.GetById(ctx, id.String())
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*lesson), nil
}

func (us *LessonUseCaseImpl) CreateLesson(ctx context.Context, insertDTO dtos.LessonInsertDTO) (*dtos.LessonDTO, error) {
	domain, err := us.mappers.InsertDTOToDomain(insertDTO)
	if err != nil {
		return nil, err
	}

	if _, err := us.moduleRepository.GetByCourseId(ctx, domain.ModuleID().String()); err != nil {
		return nil, err
	}

	domainCreated, err := us.lessonRepository.Create(ctx, *domain)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*domainCreated), nil
}

func (us *LessonUseCaseImpl) AddResource(ctx context.Context, id uuid.UUID, resource domain.Resource) (*dtos.LessonDTO, error) {
	lesson, err := us.lessonRepository.GetById(ctx, id.String())
	if err != nil {
		return nil, err
	}

	if err := lesson.AddResource(resource); err != nil {
		return nil, err
	}

	lessonUpdated, err := us.lessonRepository.Update(ctx, id, *lesson)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*lessonUpdated), nil

}

func (us *LessonUseCaseImpl) UpdateLesson(ctx context.Context, id uuid.UUID, insertDTO dtos.LessonInsertDTO) (*dtos.LessonDTO, error) {
	lesson, err := us.lessonRepository.GetById(ctx, id.String())
	if err != nil {
		return nil, err
	}

	err = lesson.UpdateContent(insertDTO.Title, insertDTO.Content, insertDTO.VideoURL, insertDTO.Duration, insertDTO.IsPreview)
	if err != nil {
		return nil, err
	}

	if _, err := us.moduleRepository.GetByCourseId(ctx, lesson.ModuleID().String()); err != nil {
		return nil, err
	}

	domainCreated, err := us.lessonRepository.Update(ctx, id, *lesson)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*domainCreated), nil
}

func (us *LessonUseCaseImpl) DeleteLesson(ctx context.Context, id uuid.UUID) error {
	if err := us.lessonRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
