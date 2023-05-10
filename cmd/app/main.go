package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	inventoryHandler "github.com/MCPTechnology/go_microservices/internal/app/handlers/inventory"
	"github.com/gorilla/mux"
)

const (
	host                   string        = ""
	port                   int32         = 9090
	appName                string        = "products_api_"
	idleTimeout            time.Duration = 500 * time.Second
	readTimeout            time.Duration = 500 * time.Second
	writeTimeout           time.Duration = 500 * time.Second
	gracefulShutdownPeriod time.Duration = 30 * time.Second
)

var addr string = fmt.Sprintf("%v:%v", host, port)

func main() {
	serveMux := mux.NewRouter()
	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	putRouter := serveMux.Methods(http.MethodPut).Subrouter()

	logger := log.New(os.Stdout, appName, log.LstdFlags)
	products := inventoryHandler.NewInventoryHandler(logger)
	// postRouter.Use(products.MiddlewareProductValidation)
	// putRouter.Use(products.MiddlewareProductValidation)

	getRouter.HandleFunc("/inventory/products", products.GetProducts)
	postRouter.HandleFunc("/inventory/product", products.AddProduct)
	putRouter.HandleFunc("/inventory/product/{id:[a-z0-9-]+}", products.UpdateProduct)

	server := &http.Server{
		Addr:         addr,
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	// * Starting server
	go func() {
		logger.Println("Starting server on port ", port)
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// * Trap termination or interrupt signals and gracefully shutdown the server
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, syscall.SIGTERM)

	// * Block until a signal is received
	signal := <-signalChannel
	logger.Printf("Received terminate signal: %v\n  Graceful shutdown...\n", signal)

	// * Gracefully shutdown the server, waiting for running operations to close
	timeoutContext, cancelFunc := context.WithTimeout(context.Background(), gracefulShutdownPeriod)
	defer cancelFunc()
	server.Shutdown(timeoutContext)
}
