package usecase

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/google/uuid"
)

type ResourceUseCaseImpl struct {
	resourceRepository output.ResourceRepository
	lessonRepository   output.LessonRepository
	mappers            mappers.ResourceMapper
}

func NewResourceUseCase(resourceRepository output.ResourceRepository, lessonRepository output.LessonRepository) input.ResourceUseCase {
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

// TODO: Add Buisness logic
func (us *ResourceUseCaseImpl) CreateResource(ctx context.Context, insertDTO dtos.ResourceInsertDTO) (*dtos.ResourceDTO, error) {
	domain := us.mappers.InsertDTOToDomain(insertDTO)

	if _, err := us.lessonRepository.GetById(ctx, insertDTO.LessonID.String()); err != nil {
		return nil, err
	}

	domainCreated, err := us.resourceRepository.Create(ctx, *domain)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*domainCreated), nil
}

// TODO Implement Correct Update
func (us *ResourceUseCaseImpl) UpdateResource(ctx context.Context, id uuid.UUID, insertDTO dtos.ResourceInsertDTO) (*dtos.ResourceDTO, error) {
	exisitingResource, err := us.resourceRepository.GetById(ctx, id.String())
	if err != nil {
		return nil, err
	}

	updatedResource := us.mappers.InsertDTOToDomain(insertDTO)
	updatedResource.ID = id
	updatedResource.CreatedAt = exisitingResource.CreatedAt
	updatedResource.UpdatedAt = exisitingResource.UpdatedAt

	if _, err := us.lessonRepository.GetById(ctx, insertDTO.LessonID.String()); err != nil {
		return nil, err
	}

	domainCreated, err := us.resourceRepository.Update(ctx, id, *updatedResource)
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
