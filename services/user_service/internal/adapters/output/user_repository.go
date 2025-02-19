package repository

import (
	"context"
	"errors"

	models "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql"
	mysql "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *entities.User) error {
	userModel := mysql.FromEntity(user)
	if err := r.db.Create(&userModel).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	var userModel models.UserModel
	if err := r.db.First(&userModel, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	user := userModel.ToEntity()
	return user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	var userModel models.UserModel
	if err := r.db.First(&userModel, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	user := userModel.ToEntity()
	return user, nil
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*entities.User, error) {
	var userModel models.UserModel
	if err := r.db.First(&userModel, "username = ?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	user := userModel.ToEntity()
	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *entities.User) error {
	userModel := mysql.FromEntity(user)
	if err := r.db.Save(&userModel).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status entities.UserStatus) error {
	if err := r.db.Model(&models.UserModel{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.Where("id = ?", id).Delete(&models.UserModel{}).Error; err != nil {
		return err
	}
	return nil
}
