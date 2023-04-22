package product_test

import (
	"testing"

	"github.com/MCPTechnology/go_microservices/internal/app/domains/product"
)

func TestProduct_NewProduct(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		description string
		price       float64
		expectedErr error
		quantity    int
	}

	testCases := []testCase{
		{
			test:        "should return error if name is empty",
			name:        "",
			description: " ",
			price:       0.0,
			quantity:    1,
			expectedErr: product.ErrMissingValues,
		}, {
			test:        "values are all Valid",
			name:        "test_name",
			description: "test product description",
			price:       1.0,
			quantity:    1,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := product.NewProduct(tc.name, tc.description, tc.price, tc.quantity)
			if err != tc.expectedErr {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}
