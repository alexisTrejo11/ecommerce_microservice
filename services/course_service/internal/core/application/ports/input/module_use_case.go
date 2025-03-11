package input

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/google/uuid"
)

type ModuleUseCase interface {
	GetModuleById(ctx context.Context, id uuid.UUID) (*dtos.ModuleDTO, error)
	CreateModule(ctx context.Context, dto dtos.ModuleInsertDTO) (*dtos.ModuleDTO, error)
	UpdateModule(ctx context.Context, id uuid.UUID, dto dtos.ModuleInsertDTO) (*dtos.ModuleDTO, error)
	DeleteModule(ctx context.Context, id uuid.UUID) error
}
