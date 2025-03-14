package mappers

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/models"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/google/uuid"
)

type LessonMappers struct{}

func (m *LessonMappers) ModelToDomain(model models.LessonModel) *domain.Lesson {
	return domain.NewLessonFromModel(
		uuid.MustParse(model.ID),
		model.Title,
		model.VideoURL,
		model.Content,
		model.ModuleID,
		model.Duration,
		model.Order,
		model.IsPreview,
		model.CreatedAt,
		model.UpdatedAt,
	)
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
		ID:        domain.ID().String(),
		Title:     domain.Title(),
		VideoURL:  domain.VideoURL(),
		ModuleID:  domain.ModuleID(),
		Content:   domain.Content(),
		Duration:  domain.Duration(),
		Order:     domain.Order(),
		IsPreview: domain.IsPreview(),
		CreatedAt: domain.CreatedAt(),
		UpdatedAt: domain.UpdatedAt(),
	}
}

func (m *LessonMappers) InsertDTOToDomain(insertDTO dtos.LessonInsertDTO) (*domain.Lesson, error) {
	return domain.NewLesson(
		insertDTO.Title,
		insertDTO.VideoURL,
		insertDTO.Content,
		insertDTO.ModuleId,
		insertDTO.Duration,
		insertDTO.Order,
		insertDTO.IsPreview,
	)
}

func (m *LessonMappers) DomainToDTO(domain domain.Lesson) *dtos.LessonDTO {
	return &dtos.LessonDTO{
		ID:        domain.ID(),
		Title:     domain.Title(),
		VideoURL:  domain.VideoURL(),
		Content:   domain.Content(),
		Duration:  domain.Duration(),
		Order:     domain.Order(),
		IsPreview: domain.IsPreview(),
		CreatedAt: domain.CreatedAt(),
		UpdatedAt: domain.UpdatedAt(),
	}
}

func (m *LessonMappers) DomainsToDTOs(lessons []domain.Lesson) *[]dtos.LessonDTO {
	lessondDTOs := make([]dtos.LessonDTO, len(lessons))
	for i, lesson := range lessons {
		lessondDTOs[i] = *m.DomainToDTO(lesson)
	}

	return &lessondDTOs
}
