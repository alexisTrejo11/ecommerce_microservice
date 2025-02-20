package repository

import (
	"context"
	"errors"

	models "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddressRepository struct {
	db            *gorm.DB
	addressMapper mappers.AddressMappers
}

func NewAddressRepository(db *gorm.DB) output.AddressRepository {
	return &AddressRepository{db: db}
}

func (r *AddressRepository) Create(ctx context.Context, address *entities.Address) error {
	AddressModel := r.addressMapper.DomainToModel(*address)
	if err := r.db.Create(&AddressModel).Error; err != nil {
		return err
	}
	return nil
}

func (r *AddressRepository) FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]*entities.Address, error) {
	var addressModels []models.AddressModel

	userIDStr := userID.String()

	if err := r.db.Where("user_id = ?", userIDStr).Find(&addressModels).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	var addresses []*entities.Address
	for _, addrs := range addressModels {
		addresses = append(addresses, r.addressMapper.ModelToDomain(&addrs))
	}

	return addresses, nil
}

func (r *AddressRepository) FindDefaultByUserID(ctx context.Context, userID uuid.UUID) (*entities.Address, error) {
	return nil, nil
}

func (r *AddressRepository) SetDefault(ctx context.Context, id uint, userID uuid.UUID) error {
	return nil
}

func (r *AddressRepository) FindByID(ctx context.Context, id uint) (*entities.Address, error) {
	var addressModel models.AddressModel
	if err := r.db.First(&addressModel, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	address := r.addressMapper.ModelToDomain(&addressModel)
	return address, nil
}

func (r *AddressRepository) FindByEmail(ctx context.Context, email string) (*entities.Address, error) {
	var addressModel models.AddressModel
	if err := r.db.First(&addressModel, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	address := r.addressMapper.ModelToDomain(&addressModel)
	return address, nil
}

func (r *AddressRepository) FindByAddressname(ctx context.Context, Addressname string) (*entities.Address, error) {
	var addressModel models.AddressModel
	if err := r.db.First(&addressModel, "Addressname = ?", Addressname).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	address := r.addressMapper.ModelToDomain(&addressModel)
	return address, nil
}

func (r *AddressRepository) Update(ctx context.Context, address *entities.Address) error {
	addressModel := r.addressMapper.DomainToModel(*address)

	if err := r.db.Save(&addressModel).Error; err != nil {
		return err
	}
	return nil
}

func (r *AddressRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status entities.UserStatus) error {
	if err := r.db.Model(&models.AddressModel{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func (r *AddressRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.Where("id = ?", id).Delete(&models.AddressModel{}).Error; err != nil {
		return err
	}
	return nil
}
