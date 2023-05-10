// The inventory handler deals with product management inside an inventory
package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/MCPTechnology/go_microservices/internal/app/domains/product"
	dtos "github.com/MCPTechnology/go_microservices/internal/app/handlers/inventory/dtos"
	"github.com/MCPTechnology/go_microservices/internal/app/services/inventory"
	"github.com/MCPTechnology/go_microservices/internal/errs"
	"github.com/MCPTechnology/go_microservices/internal/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type InventoryHandler struct {
	logger           *log.Logger
	inventoryService *inventory.InventoryService
}

func NewInventoryHandler(l *log.Logger, products ...product.Product) *InventoryHandler {
	is, err := inventory.NewInventoryService(
		inventory.WithMemoryProductRepository(products),
	)
	if err != nil {
		// TODO Handle inventory service startup error
		l.Printf("Error when configuring the Inventory Service: %v\n", err)
	}
	return &InventoryHandler{l, is}
}

func (p *InventoryHandler) GetProducts(rw http.ResponseWriter, r *http.Request) {
	productsList, err := p.inventoryService.GetAllProducts()
	if err != nil {
		http.Error(rw, "Unable to query all Products", http.StatusInternalServerError)
	}

	p.logger.Printf("Products: %v", productsList)
	utils.ToHTTPResponse(rw, http.StatusOK, productsList)
}

func (p *InventoryHandler) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := &dtos.ProductRequestDto{}
	err := utils.FromJson(r.Body, prod)
	if err != nil {
		wErr := errs.WrapError(errs.ErrBadRequest, err)
		utils.ToHTTPErrorResponse(rw, wErr)
		return
	}
	productId, err := p.inventoryService.AddProduct(prod.Name, prod.Description, prod.Price, prod.Quantity)
	if err != nil {
		utils.ToHTTPErrorResponse(rw, err)
		return
	}
	responseDto := dtos.ProductIdResponseDto{ID: productId}
	utils.ToHTTPResponse(rw, http.StatusOK, responseDto)
}

func (p *InventoryHandler) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	uid, err := uuid.Parse(idString)
	if err != nil {
		utils.ToHTTPErrorResponse(rw, errs.WrapError(errs.ErrBadRequest, err))
		return
	}
	prod := &dtos.ProductRequestDto{}
	err = utils.FromJson(r.Body, prod)
	if err != nil {
		wErr := errs.WrapError(errs.ErrBadRequest, err)
		utils.ToHTTPErrorResponse(rw, wErr)
		return
	}
	productId, err := p.inventoryService.UpdateProduct(uid, prod.Name, prod.Description, prod.Price, prod.Quantity)
	if err != nil {
		utils.ToHTTPErrorResponse(rw, err)
		return
	}
	responseDto := dtos.ProductIdResponseDto{ID: productId}
	utils.ToHTTPResponse(rw, http.StatusOK, responseDto)
}

type KeyProduct struct{}

func (p *InventoryHandler) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := dtos.ProductRequestDto{}
		err := utils.FromJson(r.Body, prod)
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
