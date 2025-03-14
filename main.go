package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Revazashvili/ecommerce-inventory-management/consumers"
	"github.com/Revazashvili/ecommerce-inventory-management/handlers"
	pd "github.com/Revazashvili/ecommerce-inventory-management/product/database"
	"github.com/Revazashvili/ecommerce-inventory-management/stock"
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

	productStorage := pd.NewProductsDatabase(pool)
	stockStorage := stock.NewStockStorage(pool)
	stockService := stock.NewService(stockStorage)

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Mount("/api/product", handlers.ProductRoutes(productStorage))
	r.Mount("/api/stock", handlers.StockRoutes(stockService))

	consumers.ListenToProductEvents(ctx, productStorage)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
