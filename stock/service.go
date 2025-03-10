package stock

import (
	"context"
	"errors"
	"log"
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

func (ss *Service) Unreserve(ctx context.Context, orderNumber uuid.UUID) error {
	sr, err := ss.storage.GetStockReservation(ctx, orderNumber)
	log.Println(*sr.CancelDate)

	if err != nil {
		return err
	}

	sr.Cancel()

	s, err := ss.storage.GetStock(ctx, sr.ProductId)

	if err != nil {
		return err
	}

	s.Unreserve(sr.Quantity)

	_, err = ss.storage.CancelStockReservation(ctx, sr)

	if err != nil {
		return err
	}

	_, err = ss.storage.UpdateStockReserve(ctx, s)

	if err != nil {
		return err
	}

	return nil
}

func (ss *Service) Reserve(ctx context.Context, productID uuid.UUID, quantity int, orderNumber uuid.UUID) error {

	sre, err := ss.storage.StockReservationExists(ctx, orderNumber)

	if err != nil {
		return err
	}

	if sre {
		return errors.New("reservation for order number already exists")
	}

	s, err := ss.storage.GetStock(ctx, productID)

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

	_, err = ss.storage.UpdateStockReserve(ctx, s)
	if err != nil {
		return err
	}

	return nil
}
