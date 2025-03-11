package usecase

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/google/uuid"
)

type ModuleUseCaseImpl struct {
	ModuleRepository output.ModuleRepository
	mappers          mappers.ModuleMapper
}

func NewModuleUseCase(ModuleRepository output.ModuleRepository) input.ModuleUseCase {
	return &ModuleUseCaseImpl{
		ModuleRepository: ModuleRepository,
	}
}

func (us *ModuleUseCaseImpl) GetModuleById(ctx context.Context, id uuid.UUID) (*dtos.ModuleDTO, error) {
	Module, err := us.ModuleRepository.GetById(ctx, id.String())
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*Module), nil
}

// TODO: Add Buisness logic
func (us *ModuleUseCaseImpl) CreateModule(ctx context.Context, insertDTO dtos.ModuleInsertDTO) (*dtos.ModuleDTO, error) {
	domain := us.mappers.InsertDTOToDomain(insertDTO)

	domainCreated, err := us.ModuleRepository.Create(ctx, *domain)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*domainCreated), nil
}

// TODO Implement Correct Update
func (us *ModuleUseCaseImpl) UpdateModule(ctx context.Context, id uuid.UUID, insertDTO dtos.ModuleInsertDTO) (*dtos.ModuleDTO, error) {
	domain := us.mappers.InsertDTOToDomain(insertDTO)
	domain.ID = id

	domainCreated, err := us.ModuleRepository.Update(ctx, id, *domain)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*domainCreated), nil
}

func (us *ModuleUseCaseImpl) DeleteModule(ctx context.Context, id uuid.UUID) error {
	if err := us.ModuleRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
