package seeds

import "github.com/MCPTechnology/go_microservices/internal/app/domains/product"

func SeedProducts() []product.Product {
	products := make([]product.Product, 10, 10)
	for i := range products {
		prod, _ := product.NewProduct(
			"Product 1",
			"Test 1",
			0.01,
		)
		products[i] = prod
	}
	return products
}
