package seeds

import (
	"fmt"

	"github.com/MCPTechnology/go_microservices/internal/app/domains/product"
)

func SeedProducts() []product.Product {
	products := make([]product.Product, 0)
	for i := 1; i < 10; i++ {
		prod, err := product.NewProduct(
			fmt.Sprintf("Product %v", i),
			fmt.Sprintf("Test %v", i),
			0.0+float64(i),
			i,
		)
		if err == nil {
			products = append(products, prod)
		}
	}
	return products
}
