package stock

import (
	"time"

	"github.com/google/uuid"
)

type Stock struct {
	Id               uuid.UUID
	ProductId        uuid.UUID
	Quantity         int
	ReservedQuantity int
	Version          int
	CreateDate       time.Time
	LastUpdateDate   time.Time
}

func (s *Stock) GetAvailableQuantity() int {
	return s.Quantity - s.ReservedQuantity
}

func (s *Stock) Reserve(quantity int) {
	s.ReservedQuantity += quantity
	s.LastUpdateDate = time.Now()
}

func (s *Stock) Unreserve(quantity int) {
	s.ReservedQuantity -= quantity
	s.LastUpdateDate = time.Now()
}

type StockReservation struct {
	Id          uuid.UUID
	ProductId   uuid.UUID
	OrderNumber uuid.UUID
	Quantity    int
	CreateDate  time.Time
	CancelDate  time.Time
}
