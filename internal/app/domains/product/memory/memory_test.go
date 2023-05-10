package memory

import (
	"fmt"
	"testing"

	productAggregate "github.com/MCPTechnology/go_microservices/internal/app/domains/product"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

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

func addProductsToRepository(t *testing.T, mpr *MemoryProductRepository, n int) {
	for i := 0; i < n; i++ {
		product, err := productAggregate.NewProduct(fmt.Sprintf("Product%v", i), "Description", 1.99, 1)
		assert.Nil(t, err)
		err = mpr.Add(product)
		assert.Nil(t, err)
	}
}

func TestMemoryProductRepository_GetAll(t *testing.T) {
	type args struct {
		expectedProductCount int
	}

	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			name:        "retrieves all products",
			args:        args{expectedProductCount: 10},
			expectedErr: nil,
		}, {
			name:        "no products found",
			args:        args{expectedProductCount: 0},
			expectedErr: productAggregate.ErrNoProductsFound,
		},
	}
	for _, tt := range tests {
		mpr := New()
		addProductsToRepository(t, mpr, tt.args.expectedProductCount)
		t.Run(tt.name, func(t *testing.T) {
			got, err := mpr.GetAll()
			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equalf(t, tt.args.expectedProductCount, len(got), "Expected to get %v products and got %v", tt.args.expectedProductCount, len(got))
		})
	}
}

func TestMemoryProductRepository_GetByID(t *testing.T) {
	mpr := New()
	existingProduct := getValidProductInstance(t)
	mpr.Add(existingProduct)

	tests := []struct {
		name        string
		id          uuid.UUID
		want        productAggregate.Product
		expectedErr error
	}{
		{
			name:        "retrieves a product",
			id:          existingProduct.GetID(),
			want:        existingProduct,
			expectedErr: nil,
		}, {
			name:        "product does not exist",
			id:          uuid.New(),
			want:        productAggregate.Product{},
			expectedErr: productAggregate.ErrProductNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mpr.GetByID(tt.id)
			assert.ErrorIs(t, err, tt.expectedErr)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestMemoryProductRepository_GetByName(t *testing.T) {
	mpr := New()
	existingProd := getValidProductInstance(t)
	err := mpr.Add(existingProd)
	assert.Nil(t, err)

	type args struct {
		name string
	}
	tests := []struct {
		name        string
		args        args
		want        productAggregate.Product
		expectedErr error
	}{
		{
			name:        "retrieve product by name",
			args:        args{name: existingProd.Name},
			want:        existingProd,
			expectedErr: nil,
		}, {
			name:        "retrieve product by name",
			args:        args{name: "any"},
			want:        productAggregate.Product{},
			expectedErr: productAggregate.ErrProductNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mpr.GetByName(tt.args.name)
			assert.ErrorIs(t, err, tt.expectedErr)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestMemoryProductRepository_Add(t *testing.T) {
	mpr := New()
	type args struct {
		newProduct productAggregate.Product
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			name: "add a product to the repository",
			args: args{
				newProduct: getValidProductInstance(t),
			},
			expectedErr: nil,
		}, {
			name: "product already exists",
			args: args{
				newProduct: getValidProductInstance(t),
			},
			expectedErr: productAggregate.ErrProductAlreadyExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mpr.Add(tt.args.newProduct)
			assert.ErrorIs(t, err, tt.expectedErr)
		})
	}
}

func TestMemoryProductRepository_Update(t *testing.T) {
	mpr := New()
	existingProd := getValidProductInstance(t)
	err := mpr.Add(existingProd)
	assert.Nil(t, err)

	type args struct {
		upprod productAggregate.Product
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			name:        "update a product",
			args:        args{upprod: existingProd},
			expectedErr: nil,
		}, {
			name:        "product does not exist",
			args:        args{upprod: getValidProductInstance(t)},
			expectedErr: productAggregate.ErrProductNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mpr.Update(tt.args.upprod)
			assert.ErrorIs(t, err, tt.expectedErr)
		})
	}
}

func TestMemoryProductRepository_Delete(t *testing.T) {
	mpr := New()
	existingProd := getValidProductInstance(t)
	err := mpr.Add(existingProd)
	assert.Nil(t, err)

	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			name:        "delete a product",
			args:        args{id: existingProd.ID},
			expectedErr: nil,
		}, {
			name:        "product does not exist",
			args:        args{id: getValidProductInstance(t).ID},
			expectedErr: productAggregate.ErrProductNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mpr.Delete(tt.args.id)
			assert.ErrorIs(t, err, tt.expectedErr)
		})
	}
}
