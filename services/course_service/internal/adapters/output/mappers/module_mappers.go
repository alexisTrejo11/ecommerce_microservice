package mappers

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/models"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/google/uuid"
)

type ModuleMapper struct {
	lessonMappers LessonMappers
}

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
		Lessons:  *m.lessonMappers.ModelsToDomains(model.Lessons),
	}
}

func (m *ModuleMapper) ModelsToDomains(models []models.ModuleModel) []domain.Module {
	modules := make([]domain.Module, len(models))
	for i, model := range models {
		modules[i] = *m.ModelToDomain(model)
	}

	return modules
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
		Lessons:  *m.lessonMappers.DomainsToDTOs(domain.Lessons),
	}
}

func (m *ModuleMapper) DomainsToDTOs(domains []domain.Module) []dtos.ModuleDTO {
	modules := make([]dtos.ModuleDTO, len(domains))
	for i, domain := range domains {
		modules[i] = *m.DomainToDTO(domain)
	}

	return modules
}
