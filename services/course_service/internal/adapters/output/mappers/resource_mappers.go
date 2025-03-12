package mappers

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/models"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/google/uuid"
)

type ResourceMapper struct{}

func (m *ResourceMapper) ModelToDomain(model models.ResourceModel) *domain.Resource {
	return domain.NewResourceFromModel(
		uuid.MustParse(model.ID),
		model.LessonID,
		model.Title,
		domain.ResourceType(model.Type),
		model.URL,
		model.CreatedAt,
		model.UpdatedAt,
	)
}

func (m *ResourceMapper) DomainToModel(resource domain.Resource) *models.ResourceModel {
	return &models.ResourceModel{
		ID:        resource.ID().String(),
		Title:     resource.Title(),
		Type:      string(resource.Type()),
		URL:       resource.URL(),
		LessonID:  resource.LessonID(),
		CreatedAt: resource.CreatedAt(),
		UpdatedAt: resource.UpdatedAt(),
	}
}

func (m *ResourceMapper) DomainToDTO(resource domain.Resource) *dtos.ResourceDTO {
	return &dtos.ResourceDTO{
		ID:        resource.ID(),
		Title:     resource.Title(),
		Type:      string(resource.Type()),
		URL:       resource.URL(),
		LessonID:  resource.LessonID(),
		CreatedAt: resource.CreatedAt(),
		UpdatedAt: resource.UpdatedAt(),
	}
}

func (m *ResourceMapper) DomainsToDTOs(resources []domain.Resource) *[]dtos.ResourceDTO {
	resourcesDTOs := make([]dtos.ResourceDTO, len(resources))
	for i, resource := range resources {
		resourcesDTOs[i] = *m.DomainToDTO(resource)
	}

	return &resourcesDTOs
}

func (m *ResourceMapper) InsertDTOToDomain(insertDTO dtos.ResourceInsertDTO) (*domain.Resource, error) {
	return domain.NewResource(
		insertDTO.LessonID,
		insertDTO.Title,
		domain.ResourceType(insertDTO.Type),
		insertDTO.URL,
	)
}
