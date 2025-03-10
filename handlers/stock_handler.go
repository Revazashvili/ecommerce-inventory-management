package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Revazashvili/ecommerce-inventory-management/stock"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func StockRoutes(service *stock.Service) chi.Router {
	r := chi.NewRouter()

	r.Post("/reserve", reserveHandler(service))

	return r
}

func reserveHandler(service *stock.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rr ReserveRequest
		err := json.NewDecoder(r.Body).Decode(&rr)

		if err != nil {
			http.Error(w, "Can't unmarshal request", http.StatusInternalServerError)
			return
		}
		err = service.Reserve(r.Context(), rr.ProductId, rr.Quantity, rr.OrderNumber)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type ReserveRequest struct {
	ProductId   uuid.UUID
	OrderNumber uuid.UUID
	Quantity    int
}
