package repository

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/models"
)

type ModuleRepository interface {
	Create(ctx context.Context, module *models.Module) error
	GetByID(ctx context.Context, id uint) (*models.Module, error)
	Update(ctx context.Context, module *models.Module) error
	Delete(ctx context.Context, id uint) error
	ListByCourse(ctx context.Context, courseID uint) ([]models.Module, error)
	UpdateOrderNumbers(ctx context.Context, moduleIDs []uint) error
}
