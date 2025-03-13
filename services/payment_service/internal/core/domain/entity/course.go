package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Course struct {
	id          uuid.UUID
	title       string
	description string
	price       float64
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   *time.Time
}

func NewCourse(title, description string, price float64) (*Course, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	if price <= 0 {
		return nil, errors.New("price must be greater than zero")
	}

	return &Course{
		id:          uuid.New(),
		title:       title,
		description: description,
		price:       price,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}, nil
}

func (c *Course) ID() uuid.UUID {
	return c.id
}

func (c *Course) Title() string {
	return c.title
}

func (c *Course) Description() string {
	return c.description
}

func (c *Course) Price() float64 {
	return c.price
}

func (c *Course) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Course) UpdatedAt() time.Time {
	return c.updatedAt
}

func (c *Course) DeletedAt() *time.Time {
	return c.deletedAt
}

func (c *Course) UpdateTitle(newTitle string) error {
	if newTitle == "" {
		return errors.New("title cannot be empty")
	}
	c.title = newTitle
	c.updatedAt = time.Now()
	return nil
}

func (c *Course) UpdateDescription(newDescription string) {
	c.description = newDescription
	c.updatedAt = time.Now()
}

func (c *Course) UpdatePrice(newPrice float64) error {
	if newPrice <= 0 {
		return errors.New("price must be greater than zero")
	}
	c.price = newPrice
	c.updatedAt = time.Now()
	return nil
}

func (c *Course) MarkAsDeleted() {
	now := time.Now()
	c.deletedAt = &now
}
