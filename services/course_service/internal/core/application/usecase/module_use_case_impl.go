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
	moduleRepository output.ModuleRepository
	courseRepository output.CourseRepository
	mappers          mappers.ModuleMapper
}

func NewModuleUseCase(ModuleRepository output.ModuleRepository, courseRepository output.CourseRepository) input.ModuleUseCase {
	return &ModuleUseCaseImpl{
		moduleRepository: ModuleRepository,
		courseRepository: courseRepository,
	}
}

func (us *ModuleUseCaseImpl) GetModuleById(ctx context.Context, id uuid.UUID) (*dtos.ModuleDTO, error) {
	module, err := us.moduleRepository.GetById(ctx, id.String())
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*module), nil
}

// TODO: Add Buisness logic
func (us *ModuleUseCaseImpl) CreateModule(ctx context.Context, insertDTO dtos.ModuleInsertDTO) (*dtos.ModuleDTO, error) {
	module := us.mappers.InsertDTOToDomain(insertDTO)

	if _, err := us.courseRepository.GetById(ctx, module.CourseID.String()); err != nil {
		return nil, err
	}

	moduleCreated, err := us.moduleRepository.Create(ctx, *module)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*moduleCreated), nil
}

// TODO Implement Correct Update
func (us *ModuleUseCaseImpl) UpdateModule(ctx context.Context, id uuid.UUID, insertDTO dtos.ModuleInsertDTO) (*dtos.ModuleDTO, error) {
	module := us.mappers.InsertDTOToDomain(insertDTO)
	module.ID = id

	if _, err := us.courseRepository.GetById(ctx, module.CourseID.String()); err != nil {
		return nil, err
	}

	domainUpdated, err := us.moduleRepository.Update(ctx, id, *module)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*domainUpdated), nil
}

func (us *ModuleUseCaseImpl) DeleteModule(ctx context.Context, id uuid.UUID) error {
	_, err := us.moduleRepository.GetById(ctx, id.String())
	if err != nil {
		return err
	}

	if err := us.moduleRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
