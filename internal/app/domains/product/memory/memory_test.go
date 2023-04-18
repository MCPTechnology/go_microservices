package memory

import (
	"testing"

	"github.com/MCPTechnology/go_microservices/internal/app/domains/product"
	"github.com/google/uuid"
)

func TestMemoryProductRepository_Add(t *testing.T) {
	repo := New()
	product, err := product.NewProduct("Beer", "Good for your social life", 1.99)
	if err != nil {
		t.Error(err)
	}

	repo.Add(product)
	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.products))
	}
}

func TestMemoryProductRepository_Get(t *testing.T) {
	repo := New()
	existingProduct, err := product.NewProduct("Beer", "Good for your social life", 1.99)
	if err != nil {
		t.Error(err)
	}

	repo.Add(existingProduct)
	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.products))
	}

	type testCase struct {
		name        string
		expectedErr error
		id          uuid.UUID
	}

	testCases := []testCase{
		{
			name:        "Successfully retrieves a Product",
			expectedErr: nil,
			id:          existingProduct.GetID(),
		}, {
			name:        "Non existing product",
			expectedErr: product.ErrProductNotFound,
			id:          uuid.New(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.GetByID(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemoryProductRepository_Delete(t *testing.T) {
	repo := New()
	existingProd, err := product.NewProduct("Beer", "Good for your health", 1.99)
	if err != nil {
		t.Error(err)
	}

	repo.Add(existingProd)
	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.products))
	}

	err = repo.Delete(existingProd.GetID())
	if err != nil {
		t.Error(err)
	}
	if len(repo.products) != 0 {
		t.Errorf("Expected 0 products, got %d", len(repo.products))
	}
}
