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

type CourseUseCaseImpl struct {
	courseRepository output.CourseRepository
	mappers          mappers.CourseMappers
}

func NewCourseUseCase(courseRepository output.CourseRepository) input.CourseUseCase {
	return &CourseUseCaseImpl{
		courseRepository: courseRepository,
	}
}

func (us *CourseUseCaseImpl) GetCourseById(ctx context.Context, id uuid.UUID) (*dtos.CourseDTO, error) {
	course, err := us.courseRepository.GetById(ctx, id.String())
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*course), nil
}

func (us *CourseUseCaseImpl) GetCoursesByCategory(ctx context.Context, category domain.CourseCategory) (*[]dtos.CourseDTO, error) {
	courses, err := us.courseRepository.GetByCategory(ctx, string(category))
	if err != nil {
		return nil, err
	}

	domainDTOs := us.mappers.DomainsToDTOs(*courses)
	return &domainDTOs, nil
}

func (us *CourseUseCaseImpl) GetCoursesByInstructorId(ctx context.Context, instructorId string) (*[]dtos.CourseDTO, error) {
	courses, err := us.courseRepository.GetByInstructorId(ctx, instructorId)
	if err != nil {
		return nil, err
	}

	domainDTOs := us.mappers.DomainsToDTOs(*courses)
	return &domainDTOs, nil
}

func (us *CourseUseCaseImpl) CourseSearch(ctx context.Context) (*[]dtos.CourseDTO, error) {
	return nil, nil
}

func (us *CourseUseCaseImpl) CreateCourse(ctx context.Context, insertDTO dtos.CourseInsertDTO) (*dtos.CourseDTO, error) {
	domain := us.mappers.InsertDTOToDomain(insertDTO)
	domain.GenerateSlug()

	domainCreated, err := us.courseRepository.Create(ctx, *domain)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*domainCreated), nil
}

func (us *CourseUseCaseImpl) UpdateCourse(ctx context.Context, id uuid.UUID, insertDTO dtos.CourseInsertDTO) (*dtos.CourseDTO, error) {
	domain := us.mappers.InsertDTOToDomain(insertDTO)
	domain.GenerateSlug()

	domain.Id = id
	domainCreated, err := us.courseRepository.Update(ctx, id, *domain)
	if err != nil {
		return nil, err
	}

	return us.mappers.DomainToDTO(*domainCreated), nil
}

func (us *CourseUseCaseImpl) DeleteCourse(ctx context.Context, id uuid.UUID) error {
	if err := us.courseRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
