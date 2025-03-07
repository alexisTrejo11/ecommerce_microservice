package usecases

import (
	"context"
	"fmt"

	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/adapters/output/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/application/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/application/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/pkg/facadeService"
	"github.com/google/uuid"
)

type CartUseCaseImpl struct {
	repository     output.CartRepository
	itemMappers    mappers.CartItemMapper
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

func (us *CartUseCaseImpl) AddItems(ctx context.Context, userID uuid.UUID, dtos []dtos.CartItemFetchedDTO) error {
	cart, err := us.repository.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	items := us.itemMappers.ProductToItemList(dtos, cart.ID)

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

func (us *CartUseCaseImpl) GetCart(ctx context.Context, userID uuid.UUID) (*domain.Cart, error) {
	cart, err := us.repository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (us *CartUseCaseImpl) DeleteCart(ctx context.Context, userID uuid.UUID) error {
	err := us.repository.DeleteCart(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
