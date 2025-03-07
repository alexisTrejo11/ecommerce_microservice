package mappers

import (
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/adapters/output/models"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/application/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/internal/domain"
	"github.com/alexisTrejo11/ecommerce_microservice/cart-service/pkg/facadeService"
	"github.com/google/uuid"
)

type CartMapper struct {
	itemMapper CartItemMapper
}

func (m *CartMapper) DomainToModel(domain domain.Cart) *models.CartModel {
	return &models.CartModel{
		ID:        domain.ID,
		UserID:    domain.UserID,
		Items:     m.itemMapper.domainsToModels(domain.Items),
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func (m *CartMapper) ModelToDomain(model models.CartModel) *domain.Cart {
	return &domain.Cart{
		ID:        model.ID,
		UserID:    model.UserID,
		Items:     m.itemMapper.modelsToDomains(model.Items),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func (m *CartMapper) DomainToDTO(model domain.Cart) *dtos.CartDTO {
	return &dtos.CartDTO{
		ID:     model.ID,
		UserID: model.UserID,
		Items:  m.itemMapper.domainsToDTOs(model.Items),
	}
}

type CartItemMapper struct{}

func (m *CartItemMapper) domainsToModels(items []domain.CartItem) []models.CartItemModel {
	models := make([]models.CartItemModel, len(items))
	for _, item := range items {
		models = append(models, *m.domainToModel(item))
	}
	return models
}

func (m *CartItemMapper) modelsToDomains(models []models.CartItemModel) []domain.CartItem {
	domains := make([]domain.CartItem, len(models))
	for _, model := range models {
		domains = append(domains, *m.modelToDomain(model))
	}

	return domains
}

func (m *CartItemMapper) domainToModel(domain domain.CartItem) *models.CartItemModel {
	return &models.CartItemModel{
		ID:        domain.ID,
		CartID:    domain.CartID,
		ProductID: domain.ProductID,
		UnitPrice: domain.UnitPrice,
		Quantity:  domain.Quantity,
		Discount:  domain.Discount,
		AddedAt:   domain.AddedAt,
	}
}

func (m *CartItemMapper) modelToDomain(model models.CartItemModel) *domain.CartItem {
	return &domain.CartItem{
		ID:        model.ID,
		CartID:    model.CartID,
		ProductID: model.ProductID,
		UnitPrice: model.UnitPrice,
		Quantity:  model.Quantity,
		Discount:  model.Discount,
		AddedAt:   model.AddedAt,
	}
}

func (m *CartItemMapper) productToDomain(product facadeService.Product, quantity uint, cartID uuid.UUID) *domain.CartItem {
	return &domain.CartItem{
		CartID:    cartID,
		ProductID: product.Id,
		UnitPrice: product.Price,
		Quantity:  int(quantity),
		Discount:  product.Disccount,
		AddedAt:   time.Now(),
	}
}

func (m *CartItemMapper) ProductToItemList(products []dtos.CartItemFetchedDTO, cartID uuid.UUID) []domain.CartItem {
	items := make([]domain.CartItem, len(products))

	for _, product := range products {
		item := m.productToDomain(product.ProductData, uint(product.Quantity), cartID)
		items = append(items, *item)
	}
	return items
}

func (m *CartItemMapper) domainsToDTOs(models []domain.CartItem) []dtos.CartItemDTO {
	domains := make([]dtos.CartItemDTO, len(models))
	for _, model := range models {
		domains = append(domains, *m.domainToDTO(model))
	}

	return domains
}

func (m *CartItemMapper) domainToDTO(domain domain.CartItem) *dtos.CartItemDTO {
	return &dtos.CartItemDTO{
		ID:        domain.ID,
		CartID:    domain.CartID,
		ProductID: domain.ProductID,
		UnitPrice: domain.UnitPrice,
		Quantity:  domain.Quantity,
		Discount:  domain.Discount,
	}
}
