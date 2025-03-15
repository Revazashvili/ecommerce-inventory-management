package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Revazashvili/ecommerce-inventory-management/stock"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func StockRoutes(service *stock.Service) chi.Router {
	r := chi.NewRouter()

	r.Get("/", stocksHandler(service))
	r.Post("/add", addHandler(service))
	r.Post("/reserve", reserveHandler(service))
	r.Post("/unreserve", unreserveHandler(service))

	return r
}

type AddStockRequest struct {
	ProductID uuid.UUID
	Quantity  int
}

func addHandler(service *stock.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var asr AddStockRequest
		err := json.NewDecoder(r.Body).Decode(&asr)

		if err != nil {
			http.Error(w, "Can't unmarshal request", http.StatusInternalServerError)
			return
		}

		err = service.AddStock(r.Context(), asr.ProductID, asr.Quantity)

		if err != nil {
			http.Error(w, "Can't add stock", http.StatusInternalServerError)
			return
		}
	}
}

type GetStocksRequest struct {
	ProductID *uuid.UUID
	From      *time.Time
	To        *time.Time
}

func stocksHandler(service *stock.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var gtr GetStocksRequest
		err := json.NewDecoder(r.Body).Decode(&gtr)

		if err != nil {
			http.Error(w, "Can't unmarshal request", http.StatusInternalServerError)
			return
		}

		stocks, err := service.GetStocks(r.Context(), gtr.ProductID, gtr.From, gtr.To)

		if err != nil {
			http.Error(w, "can't retrieve rows", http.StatusInternalServerError)
			return
		}

		if len(stocks) > 0 {
			err = json.NewEncoder(w).Encode(stocks)

			if err != nil {
				http.Error(w, "Can't marshal response", http.StatusInternalServerError)
				return
			}
		}
	}
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

func unreserveHandler(service *stock.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var urr UnreserveRequest
		err := json.NewDecoder(r.Body).Decode(&urr)

		if err != nil {
			http.Error(w, "Can't unmarshal request", http.StatusInternalServerError)
			return
		}
		err = service.Unreserve(r.Context(), urr.OrderNumber)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type UnreserveRequest struct {
	OrderNumber uuid.UUID
}
