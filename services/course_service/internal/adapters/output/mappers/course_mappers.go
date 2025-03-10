package mappers

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/models"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/google/uuid"
)

type CourseMappers struct{}

// SLUG, MODULES
func (m *CourseMappers) DomainToModel(course domain.Course) models.CourseModel {
	return models.CourseModel{
		ID:              course.Id.String(),
		Title:           course.Name,
		Description:     course.Description,
		Category:        string(course.Category),
		Level:           string(course.Level),
		Price:           course.Price,
		IsFree:          course.IsFree,
		Rating:          float64(course.Rating),
		InstructorID:    course.InstructorId,
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
		Id:          uuid.MustParse(model.ID),
		Name:        model.Title,
		Description: model.Description,
		//Category:        CourseCategory(model.Category),
		//Level:           CourseLevel(model.Level),
		Price:           model.Price,
		IsFree:          model.IsFree,
		Rating:          int(model.Rating),
		InstructorId:    model.InstructorID,
		ThumbnailURL:    model.ThumbnailURL,
		Language:        model.Language,
		ReviewCount:     model.ReviewCount,
		EnrollmentCount: model.EnrollmentCount,
		PublishedAt:     model.PublishedAt,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
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
