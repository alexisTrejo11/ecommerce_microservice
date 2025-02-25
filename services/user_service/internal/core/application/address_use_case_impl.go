package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/http/v1/dto"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql/mappers"
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

	if len(addresses) == 0 {
		return []*dto.AddressDTO{}, nil
	}

	addressDTOs := make([]*dto.AddressDTO, len(addresses))
	for i, addr := range addresses {
		addressDTOs[i] = uc.addressMappers.EntityToDTO(*addr)
	}

	return addressDTOs, nil
}

func (uc *AddressUseCasesImpl) AddAddress(ctx context.Context, addressDTO *dto.AddressInsertDTO) error {
	address := uc.addressMappers.InsertDtoToEntity(*addressDTO)
	address.UserID = addressDTO.UserID

	if err := address.Validate(); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	address.PrepareForCreate()

	existingAddresses, err := uc.addressRepository.FindAllByUserID(ctx, addressDTO.UserID)
	if err != nil {
		return fmt.Errorf("error checking existing addresses: %w", err)
	}

	if len(existingAddresses) == 0 {
		address.IsDefault = true
	}

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

	if existingAddress.UserID != addressDTO.UserID {
		return errors.New("Forbidden")
	}

	address := uc.addressMappers.InsertDtoToEntity(*addressDTO)
	address.ID = id
	address.UserID = addressDTO.UserID
	address.CreatedAt = existingAddress.CreatedAt
	address.IsDefault = existingAddress.IsDefault

	if err := address.Validate(); err != nil {
		return fmt.Errorf("validation error: %w", err)
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
		return errors.New("Forbidden")
	}

	if address.IsDefault {
		addresses, err := uc.addressRepository.FindAllByUserID(ctx, userID)
		if err != nil {
			return fmt.Errorf("error fetching addresses: %w", err)
		}

		if len(addresses) > 1 {
			for _, addr := range addresses {
				if addr.ID != id {
					addr.IsDefault = true
					if err := uc.addressRepository.Update(ctx, addr); err != nil {
						return fmt.Errorf("error updating new default address: %w", err)
					}
					break
				}
			}
		}
	}

	if err := uc.addressRepository.Delete(ctx, id); err != nil {
		return fmt.Errorf("error deleting address: %w", err)
	}

	return nil
}

func (uc *AddressUseCasesImpl) SetDefaultAddress(ctx context.Context, id uint, userID uuid.UUID) error {
	address, err := uc.addressRepository.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("error finding address: %w", err)
	}

	if address.UserID != userID {
		return errors.New("Forbidden")
	}

	if address.IsDefault {
		return nil
	}

	addresses, err := uc.addressRepository.FindAllByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("error fetching addresses: %w", err)
	}

	// Begin transaction - this should be handled by your repository
	// Note: You would need to update your repository interface to support transactions
	// or handle this differently based on your architecture
	for _, addr := range addresses {
		wasDefault := addr.IsDefault
		addr.IsDefault = (addr.ID == id)

		if wasDefault != addr.IsDefault {
			if err := uc.addressRepository.Update(ctx, addr); err != nil {
				return fmt.Errorf("error updating address: %w", err)
			}
		}
	}

	return nil
}
