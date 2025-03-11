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
	/*
		lessons := make([]domain.Lesson, len(m.Lessons))
		for i, lessonModel := range m.Lessons {
			lessons[i] = LessonModelToDomain(lessonModel)
		}
	*/

	return &domain.Module{
		ID:    uuid.MustParse(model.ID),
		Title: model.Title,
		Order: model.Order,
		//	Lessons: lessons,
	}
}

func (m *ModuleMapper) DomainToModel(domain domain.Module) *models.ModuleModel {
	/*
		lessons := make([]models.LessonModel, len(m.Lessons))
		for i, lesson := range m.Lessons {
			lessons[i] = LessonDomainToModel(lesson, m.ID)
		}
	*/

	return &models.ModuleModel{
		ID:       domain.ID.String(),
		Title:    domain.Title,
		Order:    domain.Order,
		CourseID: domain.CourseID,
		//	Lessons:  lessons,
	}
}

func (m *ModuleMapper) DomainToDTO(domain domain.Module) *dtos.ModuleDTO {
	/*
		lessons := make([]models.LessonModel, len(m.Lessons))
		for i, lesson := range m.Lessons {
			lessons[i] = LessonDomainToModel(lesson, m.ID)
		}
	*/

	return &dtos.ModuleDTO{
		ID:       domain.ID,
		Title:    domain.Title,
		Order:    domain.Order,
		CourseID: domain.CourseID,
		//	Lessons:  lessons,
	}
}
