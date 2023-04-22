// Package aggregate
// File: product.go
// Product is an aggregate that represents a product
package product

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

// ErrMissingValues is returned when a product is created without a name or description
var ErrMissingValues = errors.New("Missing values")

// Product is an aggregate that combines item with a price and quantity
type Product struct {
	// ID is the entity identifier
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Desctiption string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	UpdatedAt   string    `json:"-"`
	CreatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}

// New product will return an error if name of description is empty
func NewProduct(name string, description string, price float64, quantity int) (Product, error) {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(description) == "" {
		return Product{}, ErrMissingValues
	}
	return Product{
		ID:          uuid.New(),
		Name:        name,
		Desctiption: description,
		Price:       price,
		Quantity:    quantity,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   "",
	}, nil
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
