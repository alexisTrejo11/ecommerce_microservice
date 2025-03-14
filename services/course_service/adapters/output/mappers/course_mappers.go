package mappers

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/models"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
)

type CourseMappers struct {
	moduleMapper ModuleMapper
}

func (m *CourseMappers) DomainToModel(course domain.Course) models.CourseModel {
	return models.CourseModel{
		ID:              course.ID(),
		Title:           course.Name(),
		Description:     course.Description(),
		Category:        string(course.Category()),
		Level:           string(course.Level()),
		Slug:            course.Slug(),
		Price:           course.Price(),
		IsFree:          course.IsFree(),
		Rating:          float64(course.Rating()),
		InstructorID:    course.InstructorID(),
		Tags:            course.Tags(),
		ThumbnailURL:    course.ThumbnailURL(),
		Language:        course.Language(),
		ReviewCount:     course.ReviewCount(),
		EnrollmentCount: course.EnrollmentCount(),
		PublishedAt:     course.PublishedAt(),
		CreatedAt:       course.CreatedAt(),
		UpdatedAt:       course.UpdatedAt(),
	}
}

func (m *CourseMappers) ModelToDomain(model models.CourseModel) *domain.Course {
	course := domain.NewCourseFromModel(model)
	course.SetModules(m.moduleMapper.ModelsToDomains(model.Modules))

	return course
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
		ID:              course.ID().String(),
		Title:           course.Name(),
		Slug:            course.Slug(),
		Description:     course.Description(),
		ThumbnailURL:    course.ThumbnailURL(),
		Category:        string(course.Category()),
		Level:           string(course.Level()),
		Language:        course.Language(),
		InstructorID:    course.InstructorID().String(),
		Tags:            course.Tags(),
		Price:           course.Price(),
		IsFree:          course.IsFree(),
		IsPublished:     course.PublishedAt() != nil,
		PublishedAt:     course.PublishedAt(),
		EnrollmentCount: course.EnrollmentCount(),
		Rating:          float64(course.Rating()),
		ReviewCount:     course.ReviewCount(),
		CreatedAt:       course.CreatedAt(),
		UpdatedAt:       course.UpdatedAt(),
		Modules:         m.moduleMapper.DomainsToDTOs(course.Modules()),
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

func (m *CourseMappers) InsertDTOToDomain(dto dtos.CourseInsertDTO) (*domain.Course, error) {
	return domain.NewCourse(
		dto.Title,
		dto.Description,
		domain.CourseCategory(dto.Category),
		domain.CourseLevel(dto.Level),
		dto.Price,
		dto.IsFree,
		dto.InstructorID,
		dto.ThumbnailURL,
		dto.Language,
		dto.Tags,
	)
}

func (m *CourseMappers) FillDomainFromDTO(c *domain.Course, dto dtos.CourseInsertDTO) error {
	return c.UpdateInfo(
		dto.Title,
		dto.Description,
		domain.CourseCategory(dto.Category),
		domain.CourseLevel(dto.Level),
		dto.Price,
		dto.IsFree,
		dto.ThumbnailURL,
		dto.Language,
		dto.Tags,
	)
}
