// Package memory is an in memory implementation of the ProductRepository interface
package memory

import (
	"sync"

	productAggregate "github.com/MCPTechnology/go_microservices/internal/app/domains/product"
	"github.com/MCPTechnology/go_microservices/internal/errs"
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
func (mpr *MemoryProductRepository) GetAll() ([]productAggregate.Product, error) {
	// Collects all products from map
	var products []productAggregate.Product
	if mpr.ProductsCount() == 0 {
		return products, productAggregate.ErrNoProductsFound
	}
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

// GetProduct searches for a product based on it's parameters
func (mpr *MemoryProductRepository) GetByName(name string) (productAggregate.Product, error) {
	for _, product := range mpr.products {
		if product.Name == name {
			return product, nil
		}
	}
	return productAggregate.Product{}, productAggregate.ErrProductNotFound
}

// Add will Add a new product to the repository
func (mpr *MemoryProductRepository) Add(newProduct productAggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, err := mpr.GetByName(newProduct.Name); err == nil {
		return errs.WrapError(errs.ErrBadRequest, productAggregate.ErrProductAlreadyExists)
	}

	mpr.products[newProduct.GetID()] = newProduct
	return nil
}

// Update will change all values for a product based on it's ID
func (mpr *MemoryProductRepository) Update(upprod productAggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	_, err := mpr.GetByID(upprod.GetID())
	if err != nil {
		return errs.WrapError(errs.ErrNotFound, productAggregate.ErrProductNotFound)
	}

	mpr.products[upprod.GetID()] = upprod
	return nil
}

// Delete remove an product from the repository
func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return errs.WrapError(errs.ErrNotFound, productAggregate.ErrProductNotFound)
	}
	delete(mpr.products, id)
	return nil
}

func (mpr *MemoryProductRepository) ProductsCount() int {
	prodCount := len(mpr.products)
	return prodCount
}
