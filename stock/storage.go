package stock

import (
	"context"

	"github.com/google/uuid"
)

type Storage interface {
	Get(ctx context.Context, productID uuid.UUID) (Stock, error)
	ReserveStock(ctx context.Context, s Stock) (Stock, error)
	AddStockReservation(ctx context.Context, sr StockReservation) (StockReservation, error)
}
