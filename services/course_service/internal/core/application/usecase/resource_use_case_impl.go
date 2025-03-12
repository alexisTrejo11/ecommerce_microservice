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

type ResourceUseCaseImpl struct {
	resourceRepository output.ResourceRepository
	lessonRepository   output.LessonRepository
	mappers            mappers.ResourceMapper
}

func NewResourceUseCase(
	resourceRepository output.ResourceRepository,
	lessonRepository output.LessonRepository) input.ResourceUseCase {
	return &ResourceUseCaseImpl{
		resourceRepository: resourceRepository,
		lessonRepository:   lessonRepository,
	}
}

func (us *ResourceUseCaseImpl) GetResourceById(ctx context.Context, id uuid.UUID) (*dtos.ResourceDTO, error) {
	resource, err := us.resourceRepository.GetById(ctx, id.String())
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*resource), nil
}

func (us *ResourceUseCaseImpl) GetResourcesByLessonId(ctx context.Context, lessonId uuid.UUID) (*[]dtos.ResourceDTO, error) {
	resources, err := us.resourceRepository.GetByLessonId(ctx, lessonId.String())
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainsToDTOs(*resources), nil
}

func (us *ResourceUseCaseImpl) CreateResource(ctx context.Context, insertDTO dtos.ResourceInsertDTO) (*dtos.ResourceDTO, error) {
	domain, err := us.mappers.InsertDTOToDomain(insertDTO)
	if err != nil {
		return nil, err
	}

	if _, err := us.lessonRepository.GetById(ctx, insertDTO.LessonID.String()); err != nil {
		return nil, err
	}

	domainCreated, err := us.resourceRepository.Create(ctx, *domain)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*domainCreated), nil
}

func (us *ResourceUseCaseImpl) UpdateResource(ctx context.Context, id uuid.UUID, insertDTO dtos.ResourceInsertDTO) (*dtos.ResourceDTO, error) {
	exisitingResource, err := us.resourceRepository.GetById(ctx, id.String())
	if err != nil {
		return nil, err
	}

	err = exisitingResource.UpdateInfo(insertDTO.Title, insertDTO.URL, domain.ResourceType(insertDTO.Type))
	if err != nil {
		return nil, err
	}

	if _, err := us.lessonRepository.GetById(ctx, insertDTO.LessonID.String()); err != nil {
		return nil, err
	}

	domainCreated, err := us.resourceRepository.Update(ctx, id, *exisitingResource)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*domainCreated), nil
}

func (us *ResourceUseCaseImpl) DeleteResource(ctx context.Context, id uuid.UUID) error {
	if _, err := us.resourceRepository.GetById(ctx, id.String()); err != nil {
		return err
	}

	if err := us.resourceRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
