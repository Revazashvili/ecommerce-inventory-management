package stock

import (
	"context"

	"github.com/google/uuid"
)

type Storage interface {
	GetStock(ctx context.Context, productID uuid.UUID) (Stock, error)
	StockReservationExists(ctx context.Context, orderNumber uuid.UUID) (bool, error)
	GetStockReservation(ctx context.Context, orderNumber uuid.UUID) (StockReservation, error)
	CancelStockReservation(ctx context.Context, sr StockReservation) (StockReservation, error)
	UpdateStockReserve(ctx context.Context, s Stock) (Stock, error)
	AddStockReservation(ctx context.Context, sr StockReservation) (StockReservation, error)
}
