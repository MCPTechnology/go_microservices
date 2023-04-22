package dtos

import (
	"encoding/json"
	"io"

	"github.com/MCPTechnology/go_microservices/internal/app/domains/product"
	"github.com/google/uuid"
)

type ProductRequestDto struct {
	Name        string  `json:"name"`
	Desctiption string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

func FromProduct(product product.Product) ProductRequestDto {
	return ProductRequestDto{
		Name:        product.Name,
		Desctiption: product.Desctiption,
		Price:       product.Price,
		Quantity:    product.Quantity,
	}
}

func (p *ProductRequestDto) FromJon(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

type ProductsResponseDto []ProductRequestDto

func FromProducts(products []product.Product) ProductsResponseDto {
	productsDto := ProductsResponseDto{}
	for _, product := range products {
		productsDto = append(productsDto, FromProduct(product))
	}
	return productsDto
}

func (ps *ProductsResponseDto) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ps)
}

type ProductIdResponseDto struct {
	ID uuid.UUID `json:"id"`
}

func (ps *ProductIdResponseDto) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ps)
}
