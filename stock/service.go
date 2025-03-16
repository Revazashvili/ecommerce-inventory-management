package stock

import (
	"context"
	"errors"
	"time"

	"github.com/Revazashvili/ecommerce-inventory-management/internal"
	sd "github.com/Revazashvili/ecommerce-inventory-management/stock/database"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Service struct {
	s *sd.Storage
}

func NewService(s *sd.Storage) *Service {
	return &Service{
		s: s,
	}
}

func (ss *Service) AddStock(ctx context.Context, productID uuid.UUID, quantity int) error {
	return ss.s.ExecWithTx(ctx, func(q sd.Querier) error {
		s, err := q.GetStock(ctx, productID)

		if err != nil && err != pgx.ErrNoRows {
			return err
		}

		if err == pgx.ErrNoRows {
			now := time.Now()

			sip := sd.InsertParams{
				ID:             uuid.New(),
				ProductID:      productID,
				Quantity:       int32(quantity),
				Version:        int32(1),
				CreateDate:     now,
				LastUpdateDate: now,
			}

			return q.Insert(ctx, sip)
		} else {
			q.UpdateStockQuantity(ctx, sd.UpdateStockQuantityParams{
				ID:       s.Stock.ID,
				Version:  s.Stock.Version,
				Quantity: s.Stock.Quantity + int32(quantity),
			})
		}

		return nil
	})
}

func (ss *Service) GetStocks(ctx context.Context, productID *uuid.UUID, from *time.Time, to *time.Time) ([]sd.Stock, error) {
	return ss.s.Querier.GetStocks(ctx, sd.GetStocksParams{
		ProductID: internal.ToPgTypeUUID(productID),
		From:      from,
		To:        to,
	})
}

func (ss *Service) Unreserve(ctx context.Context, orderNumber uuid.UUID) error {
	return ss.s.ExecWithTx(ctx, func(q sd.Querier) error {
		dsr, err := q.GetStockReservation(ctx, orderNumber)

		if err != nil {
			if err == pgx.ErrNoRows {
				return errors.New("stock reserverion not found")
			}

			return err
		}

		sr := dsr.StockReservation

		s, err := q.GetStock(ctx, sr.ProductID)

		if err != nil {
			return err
		}

		n := time.Now()
		err = q.CancelStockReservation(ctx, sd.CancelStockReservationParams{ID: sr.ID, Canceldate: &n})

		if err != nil {
			return err
		}

		err = q.UpdateStockReserve(ctx, sd.UpdateStockReserveParams{
			Reservedquantity: s.Stock.ReservedQuantity - sr.Quantity,
			ID:               s.Stock.ID,
			Version:          s.Stock.Version,
		})

		return err
	})
}

func (ss *Service) Reserve(ctx context.Context, productID uuid.UUID, quantity int, orderNumber uuid.UUID) error {
	return ss.s.ExecWithTx(ctx, func(q sd.Querier) error {
		sre, err := q.StockReservationExists(ctx, orderNumber)

		if err != nil {
			return err
		}

		if sre {
			return errors.New("reservation for order number already exists")
		}

		s, err := q.GetStock(ctx, productID)

		if err != nil {
			return err
		}

		aq := int(s.Stock.Quantity - s.Stock.ReservedQuantity)

		if aq < quantity {
			return errors.New("not enought available quantity")
		}

		err = q.AddStockReservation(ctx, sd.AddStockReservationParams{
			ID:          uuid.New(),
			ProductID:   productID,
			OrderNumber: orderNumber,
			Quantity:    int32(quantity),
			CreateDate:  time.Now(),
		})

		if err != nil {
			return err
		}

		err = q.UpdateStockReserve(ctx, sd.UpdateStockReserveParams{
			ID:               s.Stock.ID,
			Version:          s.Stock.Version,
			Reservedquantity: int32(s.Stock.ReservedQuantity + int32(quantity)),
		})

		return err
	})
}
