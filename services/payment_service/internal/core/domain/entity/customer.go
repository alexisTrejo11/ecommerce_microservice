package entity

import (
	"errors"

	"github.com/alexisTrejo11/ecommerce_microservice/payment-service/internal/core/domain/valueobject"
)

type Customer struct {
	id               string
	email            string
	name             string
	paymentMethods   []valueobject.PaymentMethod
	defaultPaymentID string
}

func NewCustomer(id, email, name string) (*Customer, error) {
	if id == "" || email == "" || name == "" {
		return nil, errors.New("id, email, and name cannot be empty")
	}

	return &Customer{
		id:             id,
		email:          email,
		name:           name,
		paymentMethods: []valueobject.PaymentMethod{},
	}, nil
}

func (c *Customer) AddPaymentMethod(paymentMethod string) error {
	if paymentMethod == "" {
		return errors.New("paymentMethod cannot be empty")
	}
	for _, pm := range c.paymentMethods {
		if string(pm) == paymentMethod {
			return errors.New("payment method already exists")
		}
	}
	c.paymentMethods = append(c.paymentMethods, valueobject.PaymentMethod(paymentMethod))
	return nil
}

func (c *Customer) SetDefaultPaymentMethod(paymentMethod string) error {
	if !contains(c.PaymentMethods(), paymentMethod) {
		return errors.New("payment method not found")
	}
	c.defaultPaymentID = paymentMethod
	return nil
}

func (c *Customer) ID() string    { return c.id }
func (c *Customer) Email() string { return c.email }
func (c *Customer) Name() string  { return c.name }
func (c *Customer) PaymentMethods() []string {
	methods := make([]string, len(c.paymentMethods))
	for i, pm := range c.paymentMethods {
		methods[i] = string(pm)
	}
	return methods
}
func (c *Customer) DefaultPaymentID() string { return c.defaultPaymentID }

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
