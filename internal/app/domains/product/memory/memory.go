// Package memory is an in memory implementation of the ProductRepository interface
package memory

import (
	"sync"

	productAggregate "github.com/MCPTechnology/go_microservices/internal/app/domains/product"
	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]productAggregate.Product
	sync.Mutex
}

// New is a factory function to generate a new repository for products
func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]productAggregate.Product),
	}
}

// GetAll returns all products as a slice
// This will not retuns an error but a database implementation might
func (mpr *MemoryProductRepository) GetAll() (productAggregate.Products, error) {
	// Collects all products from map
	var products productAggregate.Products
	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}

// GetByID searches for a product based on it's ID
func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (productAggregate.Product, error) {
	if product, ok := mpr.products[uuid.UUID(id)]; ok {
		return product, nil
	}
	return productAggregate.Product{}, productAggregate.ErrProductNotFound
}

// Add will Add a new product to the repository
func (mpr *MemoryProductRepository) Add(newProduct productAggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[newProduct.GetID()]; ok {
		return productAggregate.ErrProductAlreadyExists
	}

	mpr.products[newProduct.GetID()] = newProduct
	return nil
}

// Update will change all values for a product based on it's ID
func (mpr *MemoryProductRepository) Update(upprod productAggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[upprod.GetID()]; !ok {
		return productAggregate.ErrProductNotFound
	}

	mpr.products[upprod.GetID()] = upprod
	return nil
}

// Delete remove an product from the repository
func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return productAggregate.ErrProductNotFound
	}
	delete(mpr.products, id)
	return nil
}
