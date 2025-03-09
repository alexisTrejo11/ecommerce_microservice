package repository

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/adapters/output/models"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/core/domain"
	"gorm.io/gorm"
)

type CartItemRepository struct {
	db *gorm.DB
}

func NewCartItemRepository(db *gorm.DB) *CartItemRepository {
	return &CartItemRepository{db: db}
}

func (r *CartItemRepository) GetItemsByCartId(ctx context.Context, cartId string) ([]models.CartItemModel, error) {
	var itemModels []models.CartItemModel
	if err := r.db.WithContext(ctx).Where("cart_id = ?", cartId).Find(&itemModels).Error; err != nil {
		return nil, err
	}

	return itemModels, nil
}

func (r *CartItemRepository) GetItemByCartAndProduct(ctx context.Context, cartId, productId string) (*models.CartItemModel, error) {
	var existingItem models.CartItemModel
	err := r.db.WithContext(ctx).Where("cart_id = ? AND product_id = ?", cartId, productId).First(&existingItem).Error
	if err != nil {
		return nil, err
	}
	return &existingItem, nil
}

func (r *CartItemRepository) CreateItem(ctx context.Context, item domain.CartItem) error {
	var existingItem models.CartItemModel
	if err := r.db.WithContext(ctx).Where("cart_id = ? AND product_id = ?", item.ID, item.ProductID).First(&existingItem).Error; err != nil {
		if err := r.db.WithContext(ctx).Create(&item).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *CartItemRepository) DeleteItems(ctx context.Context, cart domain.Cart) error {
	var currentItems []models.CartItemModel
	if err := r.db.WithContext(ctx).Where("cart_id = ?", cart.ID).Find(&currentItems).Error; err != nil {
		return err
	}

	productIDs := make(map[string]struct{})
	for _, item := range cart.Items {
		productIDs[item.ProductID.String()] = struct{}{}
	}

	for _, item := range currentItems {
		if _, exists := productIDs[item.ProductID]; !exists {
			if err := r.db.WithContext(ctx).Where("id = ?", item.ID).Delete(&models.CartItemModel{}).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *CartItemRepository) UpdateItem(ctx context.Context, item domain.CartItem, existingItem models.CartItemModel) error {
	existingItem.Quantity = item.Quantity
	existingItem.Name = item.Name
	existingItem.UnitPrice = item.UnitPrice
	existingItem.Discount = item.Discount
	if err := r.db.WithContext(ctx).Save(&existingItem).Error; err != nil {
		return err
	}

	return nil
}
