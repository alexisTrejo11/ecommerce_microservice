package mappers

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/models"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/google/uuid"
)

type ModuleMapper struct{}

func (m *ModuleMapper) InsertDTOToDomain(insertDTO dtos.ModuleInsertDTO) *domain.Module {
	return &domain.Module{
		ID:       uuid.New(),
		Title:    insertDTO.Title,
		CourseID: insertDTO.CourseID,
		Order:    insertDTO.Order,
	}
}

func (m *ModuleMapper) ModelToDomain(model models.ModuleModel) *domain.Module {
	return &domain.Module{
		ID:       uuid.MustParse(model.ID),
		Title:    model.Title,
		Order:    model.Order,
		CourseID: model.CourseID,
		//	Lessons: lessons,
	}
}

func (m *ModuleMapper) DomainToModel(domain domain.Module) *models.ModuleModel {
	return &models.ModuleModel{
		ID:       domain.ID.String(),
		Title:    domain.Title,
		Order:    domain.Order,
		CourseID: domain.CourseID,
		//	Lessons:  lessons,
	}
}

func (m *ModuleMapper) DomainToDTO(domain domain.Module) *dtos.ModuleDTO {
	return &dtos.ModuleDTO{
		ID:       domain.ID,
		Title:    domain.Title,
		Order:    domain.Order,
		CourseID: domain.CourseID,
		//	Lessons:  lessons,
	}
}
