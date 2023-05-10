package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	productAggregate "github.com/MCPTechnology/go_microservices/internal/app/domains/product"
	"github.com/MCPTechnology/go_microservices/internal/app/handlers/inventory/dtos"
	"github.com/MCPTechnology/go_microservices/internal/errs"
	"github.com/MCPTechnology/go_microservices/internal/utils"
	"github.com/MCPTechnology/go_microservices/scripts/seeds"
	"github.com/stretchr/testify/assert"
)

func TestNewInventoryHandler(t *testing.T) {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	type args struct {
		l *log.Logger
	}
	tests := []struct {
		name string
		args args
		want *InventoryHandler
	}{
		{
			name: "successfully start an inventory handler",
			args: args{
				l: logger,
			},
			want: NewInventoryHandler(logger),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInventoryHandler(tt.args.l)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func setupInventoryService(t *testing.T, products ...productAggregate.Product) *InventoryHandler {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	ih := NewInventoryHandler(logger)

	for _, prod := range products {
		_, err := ih.inventoryService.AddProduct(prod.Name, prod.Description, prod.Price, prod.Quantity)
		assert.Nil(t, err)
	}

	return ih
}

func getValidProductInstance(t *testing.T) productAggregate.Product {
	prod, err := productAggregate.NewProduct(
		"Product",
		"Description",
		1.01,
		1,
	)
	assert.Nil(t, err)
	return prod
}

func getInvalidProductInstance(t *testing.T) productAggregate.Product {
	prod, err := productAggregate.NewProduct(
		"P",
		"Des",
		1,
		-1,
	)
	assert.Nil(t, err)
	return prod
}

func TestInventoryHandler_GetProducts(t *testing.T) {
	products := seeds.SeedProducts()
	ih := setupInventoryService(t, products...)

	handler := http.HandlerFunc(ih.GetProducts)

	rw := httptest.NewRecorder()
	r := &http.Request{}

	handler.ServeHTTP(rw, r)
	assert.Equal(t, http.StatusOK, rw.Code)

	responseProducts := &[]productAggregate.Product{}
	err := utils.FromJson(rw.Body, responseProducts)
	assert.Nil(t, err)
	assert.Equal(t, len(products), len(*responseProducts))
}

func TestInventoryHandler_AddProduct(t *testing.T) {
	products := seeds.SeedProducts()
	ih := setupInventoryService(t, products...)

	handler := http.HandlerFunc(ih.AddProduct)

	rw := httptest.NewRecorder()
	r := &http.Request{}

	existingProduct := getValidProductInstance(t)

	type args struct {
		product dtos.ProductRequestDto
	}
	tests := []struct {
		name                 string
		args                 args
		wantErr              bool
		expectedStatusCode   int
		expectedErrorMessage error
	}{
		{
			name: "successfully add a product",
			args: args{
				product: dtos.NewProductRequestDtoFromProduct(existingProduct),
			},
			wantErr:              false,
			expectedStatusCode:   http.StatusOK,
			expectedErrorMessage: nil,
		}, {
			name: "product already exists",
			args: args{
				product: dtos.NewProductRequestDtoFromProduct(existingProduct),
			},
			wantErr:              true,
			expectedStatusCode:   http.StatusBadRequest,
			expectedErrorMessage: errs.ErrBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := utils.ToJson(rw, tt.args.product)
			assert.Nil(t, err)
			handler.ServeHTTP(rw, r)
			assert.Equal(t, tt.expectedStatusCode, rw.Code)

			// responseProducts := &[]productAggregate.Product{}
			// err := utils.FromJson(rw.Body, responseProducts)
			// assert.Nil(t, err)
			// assert.Equal(t, len(products), len(*responseProducts))
		})
	}
}
