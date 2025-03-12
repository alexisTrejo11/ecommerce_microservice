package mappers

import (
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/models"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/google/uuid"
)

type CourseMappers struct {
	moduleMapper ModuleMapper
}

func (m *CourseMappers) DomainToModel(course domain.Course) models.CourseModel {
	return models.CourseModel{
		ID:              course.Id.String(),
		Title:           course.Name,
		Description:     course.Description,
		Category:        string(course.Category),
		Level:           string(course.Level),
		Slug:            course.Slug,
		Price:           course.Price,
		IsFree:          course.IsFree,
		Rating:          float64(course.Rating),
		InstructorID:    course.InstructorId,
		Tags:            course.Tags,
		ThumbnailURL:    course.ThumbnailURL,
		Language:        course.Language,
		ReviewCount:     course.ReviewCount,
		EnrollmentCount: course.EnrollmentCount,
		PublishedAt:     course.PublishedAt,
		CreatedAt:       course.CreatedAt,
		UpdatedAt:       course.UpdatedAt,
	}
}

func (m *CourseMappers) ModelToDomain(model models.CourseModel) *domain.Course {
	return &domain.Course{
		Id:              uuid.MustParse(model.ID),
		Name:            model.Title,
		Description:     model.Description,
		Category:        domain.CourseCategory(model.Category),
		Level:           domain.CourseLevel(model.Level),
		Price:           model.Price,
		IsFree:          model.IsFree,
		Rating:          int(model.Rating),
		Slug:            model.Slug,
		InstructorId:    model.InstructorID,
		ThumbnailURL:    model.ThumbnailURL,
		Language:        model.Language,
		ReviewCount:     model.ReviewCount,
		Tags:            model.Tags,
		EnrollmentCount: model.EnrollmentCount,
		PublishedAt:     model.PublishedAt,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
		Modules:         m.moduleMapper.ModelsToDomains(model.Modules),
	}
}

func (m *CourseMappers) ModelsToDomains(models []models.CourseModel) *[]domain.Course {
	courses := make([]domain.Course, len(models))
	for i, model := range models {
		course := m.ModelToDomain(model)
		courses[i] = *course
	}

	return &courses
}

func (m *CourseMappers) DomainToDTO(course domain.Course) *dtos.CourseDTO {
	return &dtos.CourseDTO{
		ID:              course.Id.String(),
		Title:           course.Name,
		Slug:            course.Slug,
		Description:     course.Description,
		ThumbnailURL:    course.ThumbnailURL,
		Category:        string(course.Category),
		Level:           string(course.Level),
		Language:        course.Language,
		InstructorID:    course.InstructorId.String(),
		Tags:            course.Tags,
		Price:           course.Price,
		IsFree:          course.IsFree,
		IsPublished:     course.PublishedAt != nil,
		PublishedAt:     course.PublishedAt,
		EnrollmentCount: course.EnrollmentCount,
		Rating:          float64(course.Rating),
		ReviewCount:     course.ReviewCount,
		CreatedAt:       course.CreatedAt,
		UpdatedAt:       course.UpdatedAt,
		Modules:         m.moduleMapper.DomainsToDTOs(course.Modules),
	}
}

func (m *CourseMappers) DomainsToDTOs(courses []domain.Course) []dtos.CourseDTO {
	dtosList := make([]dtos.CourseDTO, 0, len(courses))

	for _, course := range courses {
		dto := m.DomainToDTO(course)
		dtosList = append(dtosList, *dto)
	}

	return dtosList
}

func (m *CourseMappers) InsertDTOToDomain(dto dtos.CourseInsertDTO) *domain.Course {
	return &domain.Course{
		Id:              uuid.New(),
		Name:            dto.Title,
		Description:     dto.Description,
		Category:        domain.CourseCategory(dto.Category),
		Level:           domain.CourseLevel(dto.Level),
		Language:        dto.Language,
		InstructorId:    dto.InstructorID,
		ThumbnailURL:    dto.ThumbnailURL,
		Tags:            dto.Tags,
		Price:           dto.Price,
		IsFree:          dto.IsFree,
		Rating:          0,
		ReviewCount:     0,
		EnrollmentCount: 0,
		PublishedAt:     nil,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		//Modules:         modules,
	}
}
