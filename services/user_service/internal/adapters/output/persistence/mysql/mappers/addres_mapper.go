package mappers

import (
	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/api/dto"
	models "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/google/uuid"
)

type AddressMappers struct{}

func (m AddressMappers) DtoToEntity(dto dto.AddressDTO) entities.Address {
	return entities.Address{
		ID:           dto.ID,
		UserID:       dto.UserID,
		AddressLine1: dto.AddressLine1,
		AddressLine2: *dto.AddressLine2,
		City:         dto.City,
		PostalCode:   dto.PostalCode,
		State:        dto.State,
		Country:      dto.Country,
		IsDefault:    dto.IsDefault,
	}
}

func (m AddressMappers) InsertDtoToEntity(dto dto.AddressInsertDTO) *entities.Address {
	return &entities.Address{
		AddressLine1: dto.AddressLine1,
		AddressLine2: *dto.AddressLine2,
		City:         dto.City,
		PostalCode:   dto.PostalCode,
		State:        dto.State,
		Country:      dto.Country,
		IsDefault:    dto.IsDefault,
	}
}

func (m AddressMappers) EntityToDTO(entity entities.Address) *dto.AddressDTO {
	return &dto.AddressDTO{
		ID:           entity.ID,
		UserID:       entity.UserID,
		AddressLine1: entity.AddressLine1,
		AddressLine2: &entity.AddressLine2,
		City:         entity.City,
		PostalCode:   entity.PostalCode,
		State:        entity.State,
		Country:      entity.Country,
		IsDefault:    entity.IsDefault,
	}
}

func (m AddressMappers) DomainToModel(entity entities.Address) *models.AddressModel {
	return &models.AddressModel{
		ID:           entity.ID,
		UserID:       entity.UserID.String(),
		AddressLine1: entity.AddressLine1,
		AddressLine2: entity.AddressLine2,
		City:         entity.City,
		PostalCode:   entity.PostalCode,
		State:        entity.State,
		Country:      entity.Country,
		IsDefault:    entity.IsDefault,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
}

func (m AddressMappers) ModelToDomain(model *models.AddressModel) *entities.Address {
	userID, _ := uuid.Parse(model.UserID)
	return &entities.Address{
		ID:           model.ID,
		UserID:       userID,
		AddressLine1: model.AddressLine1,
		AddressLine2: model.AddressLine2,
		City:         model.City,
		PostalCode:   model.PostalCode,
		State:        model.State,
		Country:      model.Country,
		IsDefault:    model.IsDefault,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	}
}
