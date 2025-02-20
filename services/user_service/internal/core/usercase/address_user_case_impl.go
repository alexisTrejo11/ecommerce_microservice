package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/api/dto"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output"
	"github.com/google/uuid"
)

type AddressUseCasesImpl struct {
	addressRepository output.AddressRepository
	addressMappers    mappers.AddressMappers
}

func NewAddressUseCase(addressRepository output.AddressRepository) input.AddressUseCase {
	return &AddressUseCasesImpl{addressRepository: addressRepository}
}

func (uc *AddressUseCasesImpl) GetUserAddresses(ctx context.Context, userID uuid.UUID) ([]*dto.AddressDTO, error) {
	addresses, err := uc.addressRepository.FindAllByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error fetching addresses: %w", err)
	}

	addressDTOs := make([]*dto.AddressDTO, 0)
	for _, addr := range addresses {
		addressDTOs = append(addressDTOs, uc.addressMappers.EntityToDTO(*addr))
	}

	return addressDTOs, nil
}

func (uc *AddressUseCasesImpl) AddAddress(ctx context.Context, addressDTO *dto.AddressInsertDTO) error {
	address := uc.addressMappers.InsertDtoToEntity(*addressDTO)
	address.UserID = addressDTO.UserID

	err := address.Validate()
	if err != nil {
		return err
	}

	address.PrepareForCreate()
	if err := uc.addressRepository.Create(ctx, address); err != nil {
		return fmt.Errorf("error saving address: %w", err)
	}

	return nil
}

func (uc *AddressUseCasesImpl) UpdateAddress(ctx context.Context, id uint, addressDTO *dto.AddressInsertDTO) error {
	existingAddress, err := uc.addressRepository.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("error finding address: %w", err)
	}

	address := uc.addressMappers.InsertDtoToEntity(*addressDTO)
	address.ID = id
	address.UserID = addressDTO.UserID
	address.CreatedAt = existingAddress.CreatedAt

	err = address.Validate()
	if err != nil {
		return err
	}

	address.PrepareForUpdate()
	if err := uc.addressRepository.Update(ctx, address); err != nil {
		return fmt.Errorf("error updating address: %w", err)
	}

	return nil
}

func (uc *AddressUseCasesImpl) DeleteAddress(ctx context.Context, id uint, userID uuid.UUID) error {
	address, err := uc.addressRepository.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("error finding address: %w", err)
	}

	if address.UserID != userID {
		return errors.New("unauthorized to delete this address")
	}

	if err := uc.addressRepository.Delete(ctx, id); err != nil {
		return fmt.Errorf("error deleting address: %w", err)
	}

	return nil
}

func (uc *AddressUseCasesImpl) SetDefaultAddress(ctx context.Context, id uint, userID uuid.UUID) error {
	addresses, err := uc.addressRepository.FindAllByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("error fetching addresses: %w", err)
	}

	var targetAddress *entities.Address
	for _, addr := range addresses {
		if addr.ID == id {
			targetAddress = addr
			break
		}
	}

	if targetAddress == nil {
		return errors.New("address not found or does not belong to the user")
	}

	for _, addr := range addresses {
		addr.IsDefault = (addr.ID == id)
		if err := uc.addressRepository.Update(ctx, addr); err != nil {
			return fmt.Errorf("error updating address: %w", err)
		}
	}

	return nil
}
