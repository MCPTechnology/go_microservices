package dtos

import (
	"github.com/MCPTechnology/go_microservices/internal/app/domains/product"
	"github.com/google/uuid"
)

type ProductRequestDto struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

func NewProductRequestDtoFromProduct(product product.Product) ProductRequestDto {
	return ProductRequestDto{
		Name: product.Name,
		Description: product.Description,
		Price: product.Price,
		Quantity: product.Quantity,
	}
}

type ProductIdResponseDto struct {
	ID uuid.UUID `json:"id"`
}
