package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/microservices/handlers"
)

// REST means Json over HTTP

func main() {

	ll := log.New(os.Stdout, "product-api", log.LstdFlags)

	// Create the handlers
	// hh := handlers.NewHello(ll)
	// gh := handlers.NewGoodbuy(ll)
	ph := handlers.NewProducts(ll)

	// Create a new server mux and register a handlers
	// sm := http.NewServeMux()
	// sm.Handle("/", ph)
	// sm.Handle("/goodbuy", gh)
	// sm.Handle("/products", ph)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	// sm.Handle("/products", ph)

	// Create a new server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start the server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			ll.Fatal(err)
		}
	}()

	// trap sigterms or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// signal.Notify(c, os.Kill)

	sig := <-c
	ll.Println("Recieved Terminate, graceful shutdown", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.Shutdown(ctx)
	// http.ListenAndServe(":9090", sm)

}
