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
	ResourceRepository output.ResourceRepository
	mappers            mappers.ResourceMapper
}

func NewResourceUseCase(ResourceRepository output.ResourceRepository) input.ResourceUseCase {
	return &ResourceUseCaseImpl{
		ResourceRepository: ResourceRepository,
	}
}

func (us *ResourceUseCaseImpl) GetResourceById(ctx context.Context, id uuid.UUID) (*dtos.ResourceDTO, error) {
	Resource, err := us.ResourceRepository.GetById(ctx, id.String())
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*Resource), nil
}

// TODO: Add Buisness logic
func (us *ResourceUseCaseImpl) CreateResource(ctx context.Context, insertDTO dtos.ResourceInsertDTO) (*dtos.ResourceDTO, error) {
	domain := us.mappers.InsertDTOToDomain(insertDTO)

	domainCreated, err := us.ResourceRepository.Create(ctx, *domain)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*domainCreated), nil
}

// TODO Implement Correct Update
func (us *ResourceUseCaseImpl) UpdateResource(ctx context.Context, id uuid.UUID, insertDTO dtos.ResourceInsertDTO) (*dtos.ResourceDTO, error) {
	domain := us.mappers.InsertDTOToDomain(insertDTO)
	domain.ID = id

	domainCreated, err := us.ResourceRepository.Update(ctx, id, *domain)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*domainCreated), nil
}

func (us *ResourceUseCaseImpl) DeleteResource(ctx context.Context, id uuid.UUID) error {
	if err := us.ResourceRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
