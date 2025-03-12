package mappers

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/models"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
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
		UpdatedAt: model.UpdatedAt,
	}
}

func (m *LessonMappers) ModelsToDomains(lessonsModels []models.LessonModel) *[]domain.Lesson {
	lessons := make([]domain.Lesson, len(lessonsModels))
	for i, lessonsModel := range lessonsModels {
		lessons[i] = *m.ModelToDomain(lessonsModel)
	}

	return &lessons
}

func (m *LessonMappers) DomainToModel(domain domain.Lesson) *models.LessonModel {
	return &models.LessonModel{
		ID:        domain.ID.String(),
		Title:     domain.Title,
		VideoURL:  domain.VideoURL,
		ModuleID:  domain.ModuleId,
		Content:   domain.Content,
		Duration:  domain.Duration,
		Order:     domain.Order,
		IsPreview: domain.IsPreview,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func (m *LessonMappers) InsertDTOToDomain(insertDTO dtos.LessonInsertDTO) *domain.Lesson {
	return &domain.Lesson{
		ID:        uuid.New(),
		Title:     insertDTO.Title,
		ModuleId:  insertDTO.ModuleId,
		VideoURL:  insertDTO.VideoURL,
		Content:   insertDTO.Content,
		Duration:  insertDTO.Duration,
		Order:     insertDTO.Order,
		IsPreview: insertDTO.IsPreview,
	}
}

func (m *LessonMappers) DomainToDTO(domain domain.Lesson) *dtos.LessonDTO {
	return &dtos.LessonDTO{
		ID:        domain.ID,
		Title:     domain.Title,
		VideoURL:  domain.VideoURL,
		Content:   domain.Content,
		Duration:  domain.Duration,
		Order:     domain.Order,
		IsPreview: domain.IsPreview,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func (m *LessonMappers) DomainsToDTOs(lessons []domain.Lesson) *[]dtos.LessonDTO {
	lessondDTOs := make([]dtos.LessonDTO, len(lessons))
	for i, lesson := range lessons {
		lessondDTOs[i] = *m.DomainToDTO(lesson)
	}

	return &lessondDTOs
}
