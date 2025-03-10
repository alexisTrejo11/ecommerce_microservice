package output

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/google/uuid"
)

type ModuleRepository interface {
	GetById(id string) (*domain.Module, error)
	Create(newModule domain.Module) (*domain.Module, error)
	Update(id uuid.UUID, updatedModule domain.Module) (*domain.Module, error)
	Delete(id uuid.UUID) error
}
