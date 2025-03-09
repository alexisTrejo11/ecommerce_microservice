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
	db     *gorm.DB
	mapper mappers.CartMapper
}

func NewCartRepository(db *gorm.DB) output.CartRepository {
	return &CartRepository{db: db}
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

// WORKS ?
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
	var cartItemsModel []models.CartItemModel
	if err := r.db.WithContext(ctx).Where("cart_id = ?", cartModel.ID).Find(&cartItemsModel).Error; err != nil {
		return err
	}

	cartModel.Items = cartItemsModel
	return nil
}

func (r *CartRepository) updateItems(ctx context.Context, cart domain.Cart) error {
	var currentItems []models.CartItemModel
	if err := r.db.WithContext(ctx).Where("cart_id = ?", cart.ID).Find(&currentItems).Error; err != nil {
		return err
	}

	productIDs := make(map[string]struct{})
	for _, item := range cart.Items {
		productIDs[item.ProductID.String()] = struct{}{}
	}

	// Delete
	for _, item := range currentItems {
		if _, exists := productIDs[item.ProductID]; !exists {
			if err := r.db.WithContext(ctx).Where("id = ?", item.ID).Delete(&models.CartItemModel{}).Error; err != nil {
				return err
			}
		}
	}

	// Adjust
	for _, item := range cart.Items {
		var existingItem models.CartItemModel
		// Create
		if err := r.db.WithContext(ctx).Where("cart_id = ? AND product_id = ?", cart.ID, item.ProductID).First(&existingItem).Error; err != nil {
			if err := r.db.WithContext(ctx).Create(&item).Error; err != nil {
				return err
			}
		} else {
			// Update
			existingItem.Quantity = item.Quantity
			existingItem.Name = item.Name
			existingItem.UnitPrice = item.UnitPrice
			existingItem.Discount = item.Discount
			if err := r.db.WithContext(ctx).Save(&existingItem).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
