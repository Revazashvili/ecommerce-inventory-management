package handlers

import (
	"encoding/json"
	"net/http"

	pd "github.com/Revazashvili/ecommerce-inventory-management/product/database"
	"github.com/go-chi/chi/v5"
)

func ProductRoutes(q *pd.Queries) chi.Router {
	r := chi.NewRouter()
	r.Get("/", getHandler(q))
	r.Get("/count", getCountHandler(q))
	return r
}

func getHandler(q pd.Querier) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		name := r.URL.Query().Get("name")
		products, err := q.Get(r.Context(), name)

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

func getCountHandler(q pd.Querier) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		count, err := q.GetCount(r.Context(), name)

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
