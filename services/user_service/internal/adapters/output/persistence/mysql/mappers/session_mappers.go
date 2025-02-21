package mappers

import (
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/api/dto"
	models "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SessionMappers struct{}

func getDeletedAt(deletedAt *time.Time) gorm.DeletedAt {
	if deletedAt != nil {
		return gorm.DeletedAt{Time: *deletedAt}
	}
	return gorm.DeletedAt{}
}

func (m SessionMappers) EntityToModel(session entities.Session) *models.SessionModel {
	return &models.SessionModel{
		ID:           session.ID.String(),
		UserID:       session.UserID.String(),
		RefreshToken: session.RefreshToken,
		UserAgent:    session.UserAgent,
		ClientIP:     session.ClientIP,
		ExpiresAt:    session.ExpiresAt,
		CreatedAt:    session.CreatedAt,
		UpdatedAt:    session.UpdatedAt,
		DeletedAt:    getDeletedAt(session.DeletedAt),
	}
}

func (m SessionMappers) ModelToDomain(sessionModel models.SessionModel) *entities.Session {
	id, _ := uuid.Parse(sessionModel.ID)
	userId, _ := uuid.Parse(sessionModel.UserID)
	return &entities.Session{
		ID:           id,
		UserID:       userId,
		RefreshToken: sessionModel.RefreshToken,
		UserAgent:    sessionModel.UserAgent,
		ClientIP:     sessionModel.ClientIP,
		ExpiresAt:    sessionModel.ExpiresAt,
		CreatedAt:    sessionModel.CreatedAt,
		UpdatedAt:    sessionModel.UpdatedAt,
		DeletedAt:    &sessionModel.DeletedAt.Time,
	}
}

func (m SessionMappers) ToSessionDTO(session entities.Session) dto.SessionDTO {
	return dto.SessionDTO{
		ID:        session.ID.String(),
		UserID:    session.UserID.String(),
		UserAgent: session.UserAgent,
		ClientIP:  session.ClientIP,
		ExpiresAt: session.ExpiresAt.Format(time.RFC3339),
		CreatedAt: session.CreatedAt,
	}
}

func (m SessionMappers) DtoToDomain(dto dto.SessionDTO) entities.Session {
	id, _ := uuid.Parse(dto.ID)
	userId, _ := uuid.Parse(dto.UserID)
	expiresAt, _ := time.Parse(time.RFC3339, dto.ExpiresAt)

	return entities.Session{
		ID:        id,
		UserID:    userId,
		UserAgent: dto.UserAgent,
		ClientIP:  dto.ClientIP,
		ExpiresAt: expiresAt,
	}
}
