package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Revazashvili/ecommerce-inventory-management/consumers"
	"github.com/Revazashvili/ecommerce-inventory-management/handlers"
	pd "github.com/Revazashvili/ecommerce-inventory-management/product/database"
	"github.com/Revazashvili/ecommerce-inventory-management/stock"
	sd "github.com/Revazashvili/ecommerce-inventory-management/stock/database"
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

	defer pool.Close()

	pd := pd.NewProductsDatabase(pool)
	sd := sd.NewStockDatabase(pool)
	ss := stock.NewService(sd)

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Mount("/api/product", handlers.ProductRoutes(pd))
	r.Mount("/api/stock", handlers.StockRoutes(ss))

	consumers.ListenToProductEvents(ctx, pd)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
