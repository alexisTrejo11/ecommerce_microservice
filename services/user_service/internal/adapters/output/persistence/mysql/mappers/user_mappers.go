package mappers

import (
	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/api/dto"
	models "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
)

func ToUserDTO(user *entities.User) *dto.UserDTO {
	return &dto.UserDTO{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		RoleID:    user.RoleID,
		RoleName:  user.Role.Name,
		Status:    int(user.Status),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

func FromUserDTO(dto *dto.UserDTO) *entities.User {
	return &entities.User{
		ID:        dto.ID,
		Email:     dto.Email,
		Username:  dto.Username,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Phone:     dto.Phone,
		RoleID:    dto.RoleID,
		Status:    entities.UserStatus(dto.Status),
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
		DeletedAt: dto.DeletedAt,
	}
}

func FromEntity(user *entities.User) *models.UserModel {
	return &models.UserModel{
		ID:           user.ID,
		Email:        user.Email,
		Username:     user.Username,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		PasswordHash: user.PasswordHash,
		Phone:        user.Phone,
		RoleID:       user.RoleID,
		Role:         nil, //FromRoleEntity(user.Role),
		Addresses:    nil, //ToAddressModels(user.Addresses),
		Status:       models.UserStatus(entities.UserStatusActive),
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    user.DeletedAt,
	}
}

type UserMappers struct{}

func (m UserMappers) SignupDTOToDomain(dto dto.SignupDTO) *entities.User {
	return &entities.User{
		FirstName:    dto.FirstName,
		LastName:     dto.LastName,
		Email:        dto.Email,
		PasswordHash: dto.Password,
		Username:     dto.Username,
		Phone:        dto.Phone,
	}
}
