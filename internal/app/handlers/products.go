package handlers

import (
	"log"
	"net/http"

	"github.com/MCPTechnology/go_microservices/internal/app/services/inventory"
	"github.com/MCPTechnology/go_microservices/scripts/seeds"
)

type Products struct {
	logger           *log.Logger
	inventoryService *inventory.InventoryService
}

func NewProducts(l *log.Logger) *Products {
	is, err := inventory.NewInventoryService(
		inventory.WithMemoryProductRepository(seeds.SeedProducts()),
	)
	if err != nil {
		l.Printf("Error when configuring the Inventory Service: %v\n", err)
	}
	return &Products{l, is}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}

	// Catch all else
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp, err := p.inventoryService.GetAllProducts()
	if err != nil {
		http.Error(rw, "Unable to query all Products", http.StatusInternalServerError)
	}
	err = lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

