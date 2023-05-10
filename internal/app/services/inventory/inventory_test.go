// Inventory package holds all services that are responsible for managing products

package inventory

import (
	"testing"

	"github.com/MCPTechnology/go_microservices/internal/errs"
	"github.com/MCPTechnology/go_microservices/scripts/seeds"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewInventoryService(t *testing.T) {
	seedProducts := seeds.SeedProducts()
	type args struct {
		cfgs []InventoryConfiguration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "successfully configure an inventory services",
			args: args{
				cfgs: []InventoryConfiguration{
					WithMemoryProductRepository(seedProducts),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is, err := NewInventoryService(tt.args.cfgs...)
			assert.Nil(t, err)
			products, err := is.GetAllProducts()
			assert.Nil(t, err)
			assert.Equal(t, len(seedProducts), len(products))
		})
	}
}

func TestInventoryService_AddProduct(t *testing.T) {
	seedProducts := seeds.SeedProducts()
	is, err := NewInventoryService(
		WithMemoryProductRepository(seedProducts),
	)
	assert.Nil(t, err)

	type args struct {
		name        string
		description string
		price       float64
		quantity    int
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			name: "successfully add a product",
			args: args{
				name:        "Product",
				description: "description",
				price:       1.01,
				quantity:    1,
			},
			expectedErr: nil,
		}, {
			name: "validation error when adding a product",
			args: args{
				name:        "P",
				description: "desc",
				price:       1.01,
				quantity:    -1,
			},
			expectedErr: errs.ErrValidation,
		}, {
			name: "product already exists",
			args: args{
				name:        "Product",
				description: "description",
				price:       1.01,
				quantity:    1,
			},
			expectedErr: errs.ErrBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := is.AddProduct(tt.args.name, tt.args.description, tt.args.price, tt.args.quantity)
			assert.ErrorIs(t, err, tt.expectedErr)
		})
	}
}

func TestInventoryService_UpdateProduct(t *testing.T) {
	seedProducts := seeds.SeedProducts()
	is, err := NewInventoryService(
		WithMemoryProductRepository(seedProducts),
	)
	assert.Nil(t, err)

	type args struct {
		uid         uuid.UUID
		name        string
		description string
		price       float64
		quantity    int
	}
	tests := []struct {
		name        string
		args        args
		want        uuid.UUID
		expectedErr error
	}{
		{
			name: "successfully update a product",
			args: args{
				uid:         seedProducts[0].ID,
				name:        "NewName",
				description: "description",
				price:       10.1,
				quantity:    2,
			},
			expectedErr: nil,
		}, {
			name: "invalid product update",
			args: args{
				uid:         seedProducts[0].ID,
				name:        "NewName",
				description: "description",
				price:       -10.1,
				quantity:    2,
			},
			expectedErr: errs.ErrValidation,
		},{
			name: "unexisting product",
			args: args{
				uid:         uuid.New(),
				name:        "NewName",
				description: "description",
				price:       10.1,
				quantity:    2,
			},
			expectedErr: errs.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := is.UpdateProduct(tt.args.uid, tt.args.name, tt.args.description, tt.args.price, tt.args.quantity)
			assert.ErrorIs(t, err, tt.expectedErr)
		})
	}
}
