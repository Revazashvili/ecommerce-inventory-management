package stock

import (
	"context"

	"github.com/google/uuid"
)

type Storage interface {
	GetStock(ctx context.Context, productID uuid.UUID) (Stock, error)
	StockReservationExists(ctx context.Context, orderNumber uuid.UUID) (bool, error)
	ReserveStock(ctx context.Context, s Stock) (Stock, error)
	AddStockReservation(ctx context.Context, sr StockReservation) (StockReservation, error)
}
