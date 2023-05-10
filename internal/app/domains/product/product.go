// Package aggregate
// File: product.go
// Product is an aggregate that represents a product
package product

import (
	"time"

	"github.com/MCPTechnology/go_microservices/internal/errs"
	"github.com/MCPTechnology/go_microservices/internal/validation"
	"github.com/google/uuid"
)

// Product is an aggregate that combines item with a price and quantity
type Product struct {
	// ID is the entity identifier
	ID          uuid.UUID `json:"uid"`
	Name        string    `json:"name" validate:"required,min=3,max=20"`
	Description string    `json:"description" validate:"required,min=5,max=150"`
	Price       float64   `json:"price" validate:"required,gt=0"`
	Quantity    int       `json:"quantity" validate:"required,gt=0"`
	UpdatedAt   string    `json:"-"`
	CreatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}

// New product will return an error if name of description is empty
func NewProduct(name string, description string, price float64, quantity int) (Product, error) {
	p := Product{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   "",
	}
	err := validation.Validate(p)
	if err != nil {
		return Product{}, errs.WrapError(errs.ErrValidation, err...)
	}
	return p, nil
}

func (p Product) GetID() uuid.UUID {
	return p.ID
}

func (p Product) GetPrice() float64 {
	return p.Price
}

func (p Product) GetQuantity() int {
	return p.Quantity
}

func (p Product) DeepEqual(target Product) bool {
	return (p.Name == target.Name && p.Description == target.Description && p.Price == target.Price && p.Quantity == target.Quantity)
}
