package mappers

import (
	"strings"

	models "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/google/uuid"
)

type MfaMappers struct{}

func (m *MfaMappers) DomainToModel(domain entities.MFA) *models.MFAModel {
	return &models.MFAModel{
		ID:          domain.ID,
		UserID:      domain.UserID.String(),
		Enabled:     domain.Enabled,
		BackupCodes: strings.Join(domain.BackupCodes, ","),
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
	}
}

func (m *MfaMappers) ModelToDomain(model models.MFAModel) *entities.MFA {
	userId, _ := uuid.Parse(model.UserID)
	return &entities.MFA{
		ID:          model.ID,
		UserID:      userId,
		Enabled:     model.Enabled,
		BackupCodes: strings.Split(model.BackupCodes, ","),
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		DeletedAt:   model.DeletedAt,
	}
}
