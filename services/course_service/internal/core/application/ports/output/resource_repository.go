package output

import (
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/domain"
	"github.com/google/uuid"
)

type ResourceRepository interface {
	GetById(id string) (*domain.Resource, error)
	Create(newResource domain.Resource) (*domain.Resource, error)
	Update(id uuid.UUID, updatedResource domain.Resource) (*domain.Resource, error)
	Delete(id uuid.UUID) error
}
