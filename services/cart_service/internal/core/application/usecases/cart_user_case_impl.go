package usecases

import (
	"context"
	"fmt"

	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/adapters/output/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/core/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/core/application/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/core/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/shared/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/pkg/facadeService"
	"github.com/google/uuid"
)

type CartUseCaseImpl struct {
	repository     output.CartRepository
	itemMappers    mappers.CartItemMapper
	cartMappers    mappers.CartMapper
	productService facadeService.ProductFacadeService
}

func NewCartUseCase(
	repository output.CartRepository,
	productService facadeService.ProductFacadeService) input.CartUseCase {
	return &CartUseCaseImpl{
		repository:     repository,
		productService: productService,
	}
}

// TODO: Return CARTDTO
func (us *CartUseCaseImpl) CreateCart(ctx context.Context, userID uuid.UUID) error {
	newCart := domain.NewCart(userID)

	if _, err := us.repository.CreateCart(ctx, *newCart); err != nil {
		return err
	}

	return nil
}

func (us *CartUseCaseImpl) Buy(ctx context.Context, userID uuid.UUID, excludeItemsIDs []*uuid.UUID) error {
	cart, err := us.repository.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	subTotal, err := cart.Buy(excludeItemsIDs)
	if err != nil {
		return err
	}

	// Conect to Payment Service
	fmt.Printf("Sending a request to Payment Service: UserId %s, SubTotal %f\n", userID, subTotal)

	return nil
}

func (us *CartUseCaseImpl) AddItems(ctx context.Context, userID uuid.UUID, insertDTOS []dtos.CartItemInserDTO) error {
	cart, err := us.repository.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	productData, err := us.fetchProductData(insertDTOS)
	if err != nil {
		return err
	}

	items := us.itemMappers.ProductToItemList(*productData, cart.ID)

	cart.AddItems(items)

	us.repository.UpdateCart(ctx, *cart)

	return nil
}
func (us *CartUseCaseImpl) RemoveItems(ctx context.Context, userID uuid.UUID, itemIDs []uuid.UUID) error {
	cart, err := us.repository.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	err = cart.RemoveItems(itemIDs)
	if err != nil {
		return err
	}

	us.repository.UpdateCart(ctx, *cart)

	return nil
}

func (us *CartUseCaseImpl) GetCartByUserId(ctx context.Context, userID uuid.UUID) (*dtos.CartDTO, error) {
	cart, err := us.repository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return us.cartMappers.DomainToDTO(*cart), nil
}

func (us *CartUseCaseImpl) GetCartById(ctx context.Context, id uuid.UUID) (*dtos.CartDTO, error) {
	cart, err := us.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return us.cartMappers.DomainToDTO(*cart), nil
}

func (us *CartUseCaseImpl) DeleteCart(ctx context.Context, userID uuid.UUID) error {
	err := us.repository.DeleteCart(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

func (us *CartUseCaseImpl) fetchProductData(insertDTOS []dtos.CartItemInserDTO) (*[]dtos.CartItemFetchedDTO, error) {
	var productData []dtos.CartItemFetchedDTO
	var failedProducts []uuid.UUID

	for _, dto := range insertDTOS {
		product, err := us.productService.GetProductById(dto.ProductID)
		if err != nil || !product.IsAvalaible {
			failedProducts = append(failedProducts, dto.ProductID)
			continue
		}

		productData = append(productData, dtos.CartItemFetchedDTO{
			Quantity:    dto.Quantity,
			ProductData: *product,
		})
	}

	if len(failedProducts) > 0 {
		return &productData, fmt.Errorf("products not avalaible: %v", failedProducts)
	}

	return &productData, nil
}
