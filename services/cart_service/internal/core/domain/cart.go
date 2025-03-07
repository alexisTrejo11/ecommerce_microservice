package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Items     []CartItem
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCart(userID uuid.UUID) *Cart {
	return &Cart{
		ID:        uuid.New(),
		UserID:    userID,
		Items:     []CartItem{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c *Cart) AddItem(item CartItem) error {
	if err := c.validateMaxLimitOfItems(); err != nil {
		return err
	}

	for i, existingItem := range c.Items {
		if existingItem.ProductID == item.ProductID {
			// Actualizar cantidad en lugar de duplicar
			c.Items[i].Quantity += item.Quantity
			c.updateAction()
			return nil
		}
	}

	c.Items = append(c.Items, item)
	c.updateAction()

	return nil
}

func (c *Cart) RemoveItem(itemID uuid.UUID) error {
	if err := c.validateNotEmptyCart(); err != nil {
		return err
	}

	itemSeen := false
	newItems := []CartItem{}

	for _, item := range c.Items {
		if item.ID != itemID {
			newItems = append(newItems, item)
		} else {
			itemSeen = true
		}
	}

	if !itemSeen {
		return errors.New("item not found")
	}

	c.Items = newItems
	c.updateAction()

	return nil
}

func (c *Cart) Buy(excludeItemsIDs []*uuid.UUID) (float64, error) {
	if err := c.validateNotEmptyCart(); err != nil {
		return 0, err
	}

	if len(excludeItemsIDs) == 0 {
		return c.completePurchase()
	}

	filteredItems := c.filterItems(excludeItemsIDs)
	originalItems := c.Items

	c.Items = filteredItems
	subTotal, err := c.calculateTotal()
	if err != nil {
		c.Items = originalItems
		return 0, err
	}

	c.completeCartAction()
	return subTotal, nil
}

func (c *Cart) GetItemCount() int {
	return len(c.Items)
}

func (c *Cart) GetItems() []CartItem {
	return c.Items
}

func (c *Cart) CalculateTotal() (float64, error) {
	return c.calculateTotal()
}

func (c *Cart) validateMaxLimitOfItems() error {
	if len(c.Items) >= 20 {
		return errors.New("cart: cannot add more than 20 items")
	}
	return nil
}

func (c *Cart) validateNotEmptyCart() error {
	if len(c.Items) <= 0 {
		return errors.New("cart: cart is empty")
	}
	return nil
}

func (c *Cart) clearCart() {
	c.Items = []CartItem{}
}

func (c *Cart) completePurchase() (float64, error) {
	subTotal, err := c.calculateTotal()
	if err != nil {
		return 0, err
	}
	c.completeCartAction()
	return subTotal, nil
}

func (c *Cart) filterItems(excludeItemsIDs []*uuid.UUID) []CartItem {
	excludeMap := c.createExcludeMap(excludeItemsIDs)
	filteredItems := make([]CartItem, 0, len(c.Items))
	for _, item := range c.Items {
		if _, found := excludeMap[item.ID]; !found {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}

func (c *Cart) createExcludeMap(excludeItemsIDs []*uuid.UUID) map[uuid.UUID]bool {
	excludeMap := make(map[uuid.UUID]bool)
	for _, id := range excludeItemsIDs {
		excludeMap[*id] = true
	}
	return excludeMap
}

func (c *Cart) completeCartAction() {
	c.clearCart()
	c.updateAction()
}

func (c *Cart) calculateTotal() (float64, error) {
	if err := c.validateNotEmptyCart(); err != nil {
		return 0, err
	}

	total := 0.0
	for _, item := range c.Items {
		total += (item.UnitPrice * float64(item.Quantity)) - item.Discount
	}

	return total, nil
}

func (c *Cart) updateAction() {
	c.UpdatedAt = time.Now()
}
