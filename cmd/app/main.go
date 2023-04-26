package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	appHandlers "github.com/MCPTechnology/go_microservices/internal/app/handlers"
	"github.com/gorilla/mux"
)

const (
	ip   string = ""
	port int32  = 9090
)

var addr string = fmt.Sprintf("%v:%v", ip, port)

func main() {
	serveMux := mux.NewRouter()
	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	putRouter := serveMux.Methods(http.MethodPut).Subrouter()

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	products := appHandlers.NewProducts(logger)
	postRouter.Use(products.MiddlewareProductValidation)
	putRouter.Use(products.MiddlewareProductValidation)

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

	//* Start the Server
	go func() {
		logger.Println("Starting server on port ", port)
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// Trap sigterm or interrupt and gracefully shutdown the server
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	// Block until a signal is received
	signal := <-signalChannel
	logger.Printf("Received terminate signal: %v\nGraceful shutdon...\n", signal)

	// Gracefully shutdown the server, waiting a max of 30 seconds for running operations to close
	timeoutContext, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()
	server.Shutdown(timeoutContext)
}
