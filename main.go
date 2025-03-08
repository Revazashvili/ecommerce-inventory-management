package main

import (
	"log"
	"net/http"

	"github.com/Revazashvili/ecommerce-inventory-management/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Mount("/api/product", handlers.ProductRoutes())
	// r.Mount("/api/product", ProductRoutes())

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
