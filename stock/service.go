package stock

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (ss *Service) Reserve(ctx context.Context, productID uuid.UUID, quantity int, orderNumber uuid.UUID) error {
	s, err := ss.storage.Get(ctx, productID)

	if err != nil {
		return err
	}

	if s.GetAvailableQuantity() < quantity {
		return errors.New("not enought available quantity")
	}

	sr := StockReservation{
		Id:          uuid.New(),
		ProductId:   productID,
		OrderNumber: orderNumber,
		Quantity:    quantity,
		CreateDate:  time.Now(),
	}

	s.Reserve(quantity)

	_, err = ss.storage.AddStockReservation(ctx, sr)
	if err != nil {
		return err
	}

	_, err = ss.storage.ReserveStock(ctx, s)
	if err != nil {
		return err
	}

	return nil
}
