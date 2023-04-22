// Inventory package holds all services that are responsible for managing products
package inventory

import (
	productAggregate "github.com/MCPTechnology/go_microservices/internal/app/domains/product"
	productMemoryRepo "github.com/MCPTechnology/go_microservices/internal/app/domains/product/memory"
	dtos "github.com/MCPTechnology/go_microservices/internal/app/dtos"
	"github.com/google/uuid"
)

type InventoryConfiguration func(is *InventoryService) error

type InventoryService struct {
	products productAggregate.Repository
}

func NewInventoryService(cfgs ...InventoryConfiguration) (*InventoryService, error) {
	is := &InventoryService{}
	for _, cfg := range cfgs {
		err := cfg(is)
		if err != nil {
			return nil, err
		}
	}
	return is, nil
}

func WithMemoryProductRepository(productsList []productAggregate.Product) InventoryConfiguration {
	products := productMemoryRepo.New()
	return func(is *InventoryService) error {
		for _, product := range productsList {
			err := products.Add(product)
			if err != nil {
				return err
			}
		}
		is.products = products
		return nil
	}
}

func (is *InventoryService) AddProduct(p *dtos.ProductRequestDto) (uuid.UUID, error) {
	product, err := productAggregate.NewProduct(p.Name, p.Desctiption, p.Price, p.Quantity)
	if err != nil {
		return uuid.Nil, err
	}

	err = is.products.Add(product)
	if err != nil {
		return uuid.Nil, err
	}

	return product.GetID(), nil
}

func (is *InventoryService) GetAllProducts() (dtos.ProductsResponseDto, error) {
	products, err := is.products.GetAll()
	if err != nil {
		return nil, err
	}

	return dtos.FromProducts(products), nil
}

func (is *InventoryService) UpdateProduct(uid uuid.UUID, p *dtos.ProductRequestDto) (uuid.UUID, error) {
	product, err := productAggregate.NewProduct(p.Name, p.Desctiption, p.Price, p.Quantity)
	if err != nil {
		return uuid.Nil, err
	}
	product.ID = uid
	err = is.products.Update(product)
	if err != nil {
		return uuid.Nil, err
	}
	return product.GetID(), nil
}
