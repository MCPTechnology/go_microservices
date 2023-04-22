package seeds

import (
	"fmt"

	"github.com/MCPTechnology/go_microservices/internal/app/domains/product"
)

func SeedProducts() []product.Product {
	products := make([]product.Product, 10, 10)
	for i := range products {
		prod, _ := product.NewProduct(
			fmt.Sprintf("Product %v", i),
			fmt.Sprintf("Test %v", i),
			0.0 + float64(i),
			i,
		)
		products[i] = prod
	}
	return products
}
