package mappers

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/models"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/google/uuid"
)

type ResourceMapper struct{}

func (m *ResourceMapper) ModelToDomain(model models.ResourceModel) *domain.Resource {
	return &domain.Resource{
		ID:       uuid.MustParse(model.ID),
		Title:    model.Title,
		LessonID: model.LessonID,
		Type:     domain.ResourceType(model.Type),
		URL:      model.URL,
	}
}

func (m *ResourceMapper) DomainToModel(resource domain.Resource) *models.ResourceModel {
	return &models.ResourceModel{
		ID:       resource.ID.String(),
		Title:    resource.Title,
		Type:     string(resource.Type),
		URL:      resource.URL,
		LessonID: resource.LessonID,
	}
}

func (m *ResourceMapper) DomainToDTO(resource domain.Resource) *dtos.ResourceDTO {
	return &dtos.ResourceDTO{
		ID:       resource.ID,
		Title:    resource.Title,
		Type:     string(resource.Type),
		URL:      resource.URL,
		LessonID: resource.LessonID,
	}
}

func (m *ResourceMapper) InsertDTOToDomain(insertDTO dtos.ResourceInsertDTO) *domain.Resource {
	return &domain.Resource{
		ID:       uuid.New(),
		Title:    insertDTO.Title,
		LessonID: insertDTO.LessonID,
		Type:     domain.ResourceType(insertDTO.Type),
		URL:      insertDTO.URL,
	}
}
