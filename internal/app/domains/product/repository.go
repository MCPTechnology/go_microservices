// Package product
// File: repository.go
// It holds the repository implementation and the implementation for a Product Repository
package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound      = errors.New("The product was not found")
	ErrProductAlreadyExists = errors.New("The product already exists")
)

type Repository interface {
	GetAll() ([]Product, error)
	GetByID(id uuid.UUID) (Product, error)
	Add(product Product) error
	Update(product Product) error
	Delete(id uuid.UUID) error
}
