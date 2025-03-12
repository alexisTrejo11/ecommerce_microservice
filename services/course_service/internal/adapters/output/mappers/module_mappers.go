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

func (m *ModuleMapper) InsertDTOToDomain(insertDTO dtos.ModuleInsertDTO) (*domain.Module, error) {
	return domain.NewModule(
		insertDTO.Title,
		insertDTO.CourseID,
		insertDTO.Order,
	)
}

func (m *ModuleMapper) ModelToDomain(model models.ModuleModel) (*domain.Module, error) {
	lessons := *m.lessonMappers.ModelsToDomains(model.Lessons)

	return domain.NewModuleFromModel(
		uuid.MustParse(model.ID),
		model.Title,
		model.CourseID,
		model.Order,
		lessons,
	)
}

func (m *ModuleMapper) ModelsToDomains(models []models.ModuleModel) []domain.Module {
	modules := make([]domain.Module, len(models))
	for i, model := range models {
		domain, _ := m.ModelToDomain(model)
		modules[i] = *domain
	}

	return modules
}

func (m *ModuleMapper) DomainToModel(domain domain.Module) *models.ModuleModel {
	return &models.ModuleModel{
		ID:       domain.ID().String(),
		Title:    domain.Title(),
		Order:    domain.Order(),
		CourseID: domain.CourseID(),
		//	Lessons:  lessons,
	}
}

func (m *ModuleMapper) DomainToDTO(domain domain.Module) *dtos.ModuleDTO {
	return &dtos.ModuleDTO{
		ID:       domain.ID(),
		Title:    domain.Title(),
		Order:    domain.Order(),
		CourseID: domain.CourseID(),
		Lessons:  *m.lessonMappers.DomainsToDTOs(domain.Lessons()),
	}
}

func (m *ModuleMapper) DomainsToDTOs(domains []domain.Module) []dtos.ModuleDTO {
	modules := make([]dtos.ModuleDTO, len(domains))
	for i, domain := range domains {
		modules[i] = *m.DomainToDTO(domain)
	}

	return modules
}
