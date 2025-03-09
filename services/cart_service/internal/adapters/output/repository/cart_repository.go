package repository

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/adapters/output/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/adapters/output/models"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/core/application/ports/output"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/core/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartRepository struct {
	db             *gorm.DB
	itemRepository CartItemRepository
	mapper         mappers.CartMapper
}

func NewCartRepository(db *gorm.DB, itemRepository CartItemRepository) output.CartRepository {
	return &CartRepository{
		db:             db,
		itemRepository: itemRepository,
	}
}

func (r *CartRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.Cart, error) {
	var cartModel models.CartModel

	if err := r.db.WithContext(ctx).First(&cartModel, id).Error; err != nil {
		return nil, err
	}

	r.appendItems(ctx, &cartModel)

	return r.mapper.ModelToDomain(cartModel), nil
}

func (r *CartRepository) GetByUserID(ctx context.Context, userId uuid.UUID) (*domain.Cart, error) {
	var cartModel models.CartModel
	if err := r.db.WithContext(ctx).
		Where("user_id = ?", userId).
		First(&cartModel).Error; err != nil {
		return nil, err
	}

	r.appendItems(ctx, &cartModel)

	return r.mapper.ModelToDomain(cartModel), nil
}

func (r *CartRepository) CreateCart(ctx context.Context, cart domain.Cart) (*domain.Cart, error) {
	cartModel := r.mapper.DomainToModel(cart)

	if err := r.db.WithContext(ctx).Create(&cartModel).Error; err != nil {
		return nil, err
	}

	return r.mapper.ModelToDomain(*cartModel), nil
}

func (r *CartRepository) UpdateCart(ctx context.Context, cart domain.Cart) (*domain.Cart, error) {
	cartModel := r.mapper.DomainToModel(cart)

	if err := r.db.WithContext(ctx).
		Where("id = ?", cartModel.ID).
		Updates(cartModel).Error; err != nil {
		return nil, err
	}

	if err := r.updateItems(ctx, cart); err != nil {
		return nil, err
	}

	r.appendItems(ctx, cartModel)

	return r.mapper.ModelToDomain(*cartModel), nil
}

func (r *CartRepository) DeleteCart(ctx context.Context, userId uuid.UUID) error {
	var cartModel models.CartModel
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).First(&cartModel).Error
	if err != nil {
		return err
	}

	err = r.db.WithContext(ctx).Delete(&cartModel).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *CartRepository) appendItems(ctx context.Context, cartModel *models.CartModel) error {
	cartItemsModel, err := r.itemRepository.GetItemsByCartId(ctx, cartModel.ID)
	if err != nil {
		return nil
	}

	cartModel.Items = cartItemsModel
	return nil
}

func (r *CartRepository) updateItems(ctx context.Context, cart domain.Cart) error {
	// Check if are some candidates to be deleted
	r.itemRepository.DeleteItems(ctx, cart)

	// Create or Update depding in each item case
	for _, item := range cart.Items {
		var existingItem models.CartItemModel
		_, err := r.itemRepository.GetItemByCartAndProduct(ctx, item.ID.String(), item.ProductID.String())
		if err != nil {
			r.itemRepository.CreateItem(ctx, item)
		} else {
			r.itemRepository.UpdateItem(ctx, item, existingItem)
		}
	}

	return nil
}
