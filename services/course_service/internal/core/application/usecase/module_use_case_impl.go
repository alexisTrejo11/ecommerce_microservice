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

func (us *ModuleUseCaseImpl) GetModuleByCourseId(ctx context.Context, id uuid.UUID) (*[]dtos.ModuleDTO, error) {
	modules, err := us.moduleRepository.GetByCourseId(ctx, id.String())
	if err != nil {
		return nil, err
	}

	dtos := us.mappers.DomainsToDTOs(*modules)
	return &dtos, nil
}

func (us *ModuleUseCaseImpl) CreateModule(ctx context.Context, insertDTO dtos.ModuleInsertDTO) (*dtos.ModuleDTO, error) {
	module, err := us.mappers.InsertDTOToDomain(insertDTO)
	if err != nil {
		return nil, err
	}

	if _, err := us.courseRepository.GetById(ctx, module.CourseID().String()); err != nil {
		return nil, err
	}

	moduleCreated, err := us.moduleRepository.Create(ctx, *module)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*moduleCreated), nil
}

// Implement buisness logic for lesson
func (us *ModuleUseCaseImpl) AddLesson(ctx context.Context, id uuid.UUID, lesson domain.Lesson) error {
	existing, err := us.moduleRepository.GetById(ctx, id.String())
	if err != nil {
		return err
	}

	if err := existing.AddLesson(lesson); err != nil {
		return err
	}

	return nil
}

func (us *ModuleUseCaseImpl) UpdateModule(ctx context.Context, id uuid.UUID, insertDTO dtos.ModuleInsertDTO) (*dtos.ModuleDTO, error) {
	existing, err := us.moduleRepository.GetById(ctx, id.String())
	if err != nil {
		return nil, err
	}

	if _, err := us.courseRepository.GetById(ctx, insertDTO.CourseID.String()); err != nil {
		return nil, err
	}

	if err := existing.Update(insertDTO.Title, insertDTO.Order); err != nil {
		return nil, err
	}

	updated, err := us.moduleRepository.Update(ctx, id, *existing)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*updated), nil
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
