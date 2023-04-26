package product_test

import (
	"errors"
	"testing"

	"github.com/MCPTechnology/go_microservices/internal/app/domains/product"
	"github.com/MCPTechnology/go_microservices/pkg/errs"
)

func TestProduct_NewProduct(t *testing.T) {
	type testCase struct {
		name        string
		productName string
		description string
		price       float64
		quantity    int
	}

	tc := testCase{
		name:        "Product fields are valid",
		productName: "Product 1",
		description: "test product description",
		price:       1.0,
		quantity:    1,
	}

	t.Run(tc.name, func(t *testing.T) {
		_, err := product.NewProduct(tc.productName, tc.description, tc.price, tc.quantity)
		if err != nil {
			t.Fatal("New Product instance shouldn't return any errors")
		}
	})
}

func TestProduct_ValidationError(t *testing.T) {
	type testCase struct {
		name        string
		productName string
		description string
		price       float64
		quantity    int
	}

	tc := testCase{
		name:        "Product fields are invalid",
		productName: "",
		description: " ",
		price:       0.0,
		quantity:    0,
	}

	t.Run(tc.name, func(t *testing.T) {
		_, err := product.NewProduct(tc.productName, tc.description, tc.price, tc.quantity)
		if !errors.Is(err, errs.ErrValidationError) {
			t.Errorf("Expected error %v and got: %v", errs.ErrValidationError, err)
		}
	})
}
