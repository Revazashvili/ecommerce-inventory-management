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
		rss, err := q.GetStockReservations(ctx, orderNumber)

		if err != nil {
			if err == pgx.ErrNoRows {
				return errors.New("stock reserverion not found")
			}

			return err
		}

		for _, r := range rss {
			s, err := q.GetStock(ctx, r.ProductID)

			if err != nil {
				return err
			}

			n := time.Now()
			err = q.CancelStockReservation(ctx, sd.CancelStockReservationParams{ID: r.ID, Canceldate: &n})

			if err != nil {
				return err
			}

			err = q.UpdateStockReserve(ctx, sd.UpdateStockReserveParams{
				Reservedquantity: s.Stock.ReservedQuantity - r.Quantity,
				ID:               s.Stock.ID,
				Version:          s.Stock.Version,
			})

			if err != nil {
				return err
			}
		}

		return nil
	})
}

type ProductToReserve struct {
	ProductId uuid.UUID
	Quantity  int
}

func (ss *Service) Reserve(ctx context.Context, products []ProductToReserve, orderNumber uuid.UUID) error {
	return ss.s.ExecWithTx(ctx, func(q sd.Querier) error {
		sre, err := q.StockReservationExists(ctx, orderNumber)

		if err != nil && err != pgx.ErrNoRows {
			return err
		}

		if sre {
			return errors.New("reservation for order number already exists")
		}

		for _, v := range products {
			s, err := q.GetStock(ctx, v.ProductId)

			if err != nil && err != pgx.ErrNoRows {
				return err
			}

			quantity := int32(v.Quantity)

			if (s.Stock.Quantity - s.Stock.ReservedQuantity) < quantity {
				return errors.New("not enough available quantity")
			}

			err = q.AddStockReservation(ctx, sd.AddStockReservationParams{
				ID:          uuid.New(),
				ProductID:   v.ProductId,
				OrderNumber: orderNumber,
				Quantity:    quantity,
				CreateDate:  time.Now(),
			})

			if err != nil {
				return err
			}

			err = q.UpdateStockReserve(ctx, sd.UpdateStockReserveParams{
				ID:               s.Stock.ID,
				Version:          s.Stock.Version,
				Reservedquantity: s.Stock.ReservedQuantity + quantity,
			})

			if err != nil {
				return err
			}
		}

		return nil
	})
}
