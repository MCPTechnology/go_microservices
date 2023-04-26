// Package aggregate
// File: product.go
// Product is an aggregate that represents a product
package product

import (
	"errors"
	"time"

	"github.com/MCPTechnology/go_microservices/pkg/validations"
	"github.com/google/uuid"
)

// ErrMissingValues is returned when a product is created without a name or description
var ErrMissingValues = errors.New("Missing values")

// Product is an aggregate that combines item with a price and quantity
type Product struct {
	// ID is the entity identifier
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" validate:"required,min=3,max=20"`
	Desctiption string    `json:"description" validate:"required,min=5,max=150"`
	Price       float64   `json:"price" validate:"required,gte=0"`
	Quantity    int       `json:"quantity" validate:"required,gte=0"`
	UpdatedAt   string    `json:"-"`
	CreatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}

// New product will return an error if name of description is empty
func NewProduct(name string, description string, price float64, quantity int) (Product, error) {
	p := Product{
		ID:          uuid.New(),
		Name:        name,
		Desctiption: description,
		Price:       price,
		Quantity:    quantity,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   "",
	}
	err := validations.Validate(p)
	if err != nil {
		return Product{}, err
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
