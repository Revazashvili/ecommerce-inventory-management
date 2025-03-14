package stock

import (
	"context"
	"errors"
	"time"

	sd "github.com/Revazashvili/ecommerce-inventory-management/stock/database"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service struct {
	q *sd.Queries
}

func NewService(q *sd.Queries) *Service {
	return &Service{
		q: q,
	}
}

func (ss *Service) GetStocks(ctx context.Context, productID *uuid.UUID, from *time.Time, to *time.Time) ([]sd.Stock, error) {

	pID := pgtype.UUID{Valid: false}
	if productID != nil {
		pID = pgtype.UUID{Valid: true, Bytes: *productID}
	}

	return ss.q.GetStocks(ctx, sd.GetStocksParams{
		ProductID: pID,
		From:      from,
		To:        to,
	})
}

func (ss *Service) Unreserve(ctx context.Context, orderNumber uuid.UUID) error {
	dsr, err := ss.q.GetStockReservation(ctx, orderNumber)

	if err != nil {
		return err
	}

	sr := dsr.StockReservation

	s, err := ss.q.GetStock(ctx, sr.ProductID)

	if err != nil {
		return err
	}

	n := time.Now()
	err = ss.q.CancelStockReservation(ctx, sd.CancelStockReservationParams{ID: sr.ID, Canceldate: &n})

	if err != nil {
		return err
	}

	err = ss.q.UpdateStockReserve(ctx, sd.UpdateStockReserveParams{
		Reservedquantity: s.Stock.ReservedQuantity - sr.Quantity,
		ID:               s.Stock.ID,
		Version:          s.Stock.Version,
	})

	if err != nil {
		return err
	}

	return nil
}

func (ss *Service) Reserve(ctx context.Context, productID uuid.UUID, quantity int, orderNumber uuid.UUID) error {

	sre, err := ss.q.StockReservationExists(ctx, orderNumber)

	if err != nil {
		return err
	}

	if sre {
		return errors.New("reservation for order number already exists")
	}

	s, err := ss.q.GetStock(ctx, productID)

	if err != nil {
		return err
	}

	aq := int(s.Stock.Quantity - s.Stock.ReservedQuantity)

	if aq < quantity {
		return errors.New("not enought available quantity")
	}

	err = ss.q.AddStockReservation(ctx, sd.AddStockReservationParams{
		ID:          uuid.New(),
		ProductID:   productID,
		OrderNumber: orderNumber,
		Quantity:    int32(quantity),
		CreateDate:  time.Now(),
	})

	if err != nil {
		return err
	}

	err = ss.q.UpdateStockReserve(ctx, sd.UpdateStockReserveParams{
		ID:               s.Stock.ID,
		Version:          s.Stock.Version,
		Reservedquantity: int32(s.Stock.ReservedQuantity + int32(quantity)),
	})

	if err != nil {
		return err
	}

	return nil
}
