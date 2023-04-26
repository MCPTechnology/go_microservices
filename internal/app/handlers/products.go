package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	dtos "github.com/MCPTechnology/go_microservices/internal/app/dtos"
	"github.com/MCPTechnology/go_microservices/internal/app/services/inventory"
	"github.com/MCPTechnology/go_microservices/scripts/seeds"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	logger           *log.Logger
	inventoryService *inventory.InventoryService
}

func NewProducts(l *log.Logger) *ProductHandler {
	is, err := inventory.NewInventoryService(
		inventory.WithMemoryProductRepository(seeds.SeedProducts()),
	)
	if err != nil {
		l.Printf("Error when configuring the Inventory Service: %v\n", err)
	}
	return &ProductHandler{l, is}
}

func (p *ProductHandler) GetProducts(rw http.ResponseWriter, r *http.Request) {
	productsList, err := p.inventoryService.GetAllProducts()
	if err != nil {
		http.Error(rw, "Unable to query all Products", http.StatusInternalServerError)
	}

	err = productsList.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *ProductHandler) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*dtos.ProductRequestDto)
	productId, err := p.inventoryService.AddProduct(prod)
	if err != nil {
		http.Error(rw, "Unable to add product", http.StatusInternalServerError)
		return
	}
	responseDto := dtos.ProductIdResponseDto{ID: productId}
	responseDto.ToJson(rw)
}

func (p *ProductHandler) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]

	uid, err := uuid.Parse(idString)
	if err != nil {
		http.Error(rw, "Invalid Product ID", http.StatusBadRequest)
		return
	}
	prod := r.Context().Value(KeyProduct{}).(*dtos.ProductRequestDto)
	productId, err := p.inventoryService.UpdateProduct(uid, prod)
	if err != nil {
		errorStr := fmt.Sprintf("Unable to perform update due to error: %v", err)
		http.Error(rw, errorStr, http.StatusBadRequest)
		return
	}
	responseDto := dtos.ProductIdResponseDto{ID: productId}
	responseDto.ToJson(rw)
}

type KeyProduct struct{}

func (p *ProductHandler) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &dtos.ProductRequestDto{}
		err := prod.FromJon(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
