package mappers

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/models"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/google/uuid"
)

type LessonMappers struct{}

func (m *LessonMappers) ModelToDomain(model models.LessonModel) *domain.Lesson {
	return &domain.Lesson{
		ID:        uuid.MustParse(model.ID),
		Title:     model.Title,
		VideoURL:  model.VideoURL,
		Content:   model.Content,
		Duration:  model.Duration,
		Order:     model.Order,
		IsPreview: model.IsPreview,
		CreatedAt: model.CreatedAt,
	}
}

func (m *LessonMappers) DomainToModel(domain domain.Lesson) *models.LessonModel {
	return &models.LessonModel{
		ID:        domain.ID.String(),
		Title:     domain.Title,
		VideoURL:  domain.VideoURL,
		Content:   domain.Content,
		Duration:  domain.Duration,
		Order:     domain.Order,
		IsPreview: domain.IsPreview,
		CreatedAt: domain.CreatedAt,
	}
}
