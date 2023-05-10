package product

import (
	"testing"

	"github.com/MCPTechnology/go_microservices/internal/errs"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	type args struct {
		name        string
		description string
		price       float64
		quantity    int
	}
	tests := []struct {
		name        string
		args        args
		want        Product
		wantErr     bool
		expectedErr error
	}{
		{
			name: "Successfully create a new Protuct aggregate",
			args: args{
				name:        "Product1",
				description: "Description 1",
				price:       10.01,
				quantity:    1,
			},
			want: Product{
				Name:        "Product1",
				Description: "Description 1",
				Price:       10.01,
				Quantity:    1,
			},
			wantErr:     false,
			expectedErr: nil,
		},
		{
			name: "Validation error when creating a product aggregate",
			args: args{
				name:        "Product1",
				description: "",
				price:       10.01,
				quantity:    -1,
			},
			want:        Product{},
			wantErr:     true,
			expectedErr: errs.ErrValidation,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProduct(tt.args.name, tt.args.description, tt.args.price, tt.args.quantity)
			assert.Equalf(t, err != nil, tt.wantErr, "NewProduct() error = %v, wantErr %v", err, tt.wantErr)
			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Truef(t, got.DeepEqual(tt.want), "NewProduct() = %v, want %v", got, tt.want)
		})
	}
}
