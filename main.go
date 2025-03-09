package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Revazashvili/ecommerce-inventory-management/consumers"
	"github.com/Revazashvili/ecommerce-inventory-management/handlers"
	"github.com/Revazashvili/ecommerce-inventory-management/product"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

const dbURL = "postgres://user:pass@localhost:5432/products"

func main() {

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbURL)

	if err != nil {
		log.Println(err)
	}

	storage := product.NewStorage(pool)

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Mount("/api/product", handlers.ProductRoutes(storage))

	consumers.ListenToProductEvents(ctx, storage)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
