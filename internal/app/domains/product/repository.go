// Package product
// File: repository.go
// It holds the repository implementation and the implementation for a Product Repository
package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound      = errors.New("the product was not found")
	ErrProductAlreadyExists = errors.New("the product already exists")
	ErrNoProductsFound      = errors.New("no products were found")
)

type Repository interface {
	GetAll() ([]Product, error)
	GetByID(id uuid.UUID) (Product, error)
	GetByName(name string) (Product, error)
	Add(product Product) error
	Update(product Product) error
	Delete(id uuid.UUID) error
	ProductsCount() int
}
