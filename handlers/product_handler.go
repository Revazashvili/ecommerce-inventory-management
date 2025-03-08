package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Revazashvili/ecommerce-inventory-management/product"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const dbURL = "postgres://user:pass@localhost:5432/products"

func ProductRoutes() chi.Router {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbURL)

	if err != nil {
		log.Println(err)
	}

	storage := product.NewStorage(pool)

	r := chi.NewRouter()
	r.Get("/", getHandler(storage))
	r.Get("/count", getCountHandler(storage))
	return r
}

func getHandler(s product.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		products, err := s.Search(r.Context(), name)

		if err != nil {
			http.Error(w, "Internal error happend", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(products)

		if err != nil {
			http.Error(w, "Internal error happend", http.StatusInternalServerError)
			return
		}
	}
}

func getCountHandler(s product.Storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		count, err := s.Count(r.Context(), name)

		if err != nil {
			http.Error(w, "Internal error happend", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(count)

		if err != nil {
			http.Error(w, "Internal error happend", http.StatusInternalServerError)
			return
		}
	}
}
