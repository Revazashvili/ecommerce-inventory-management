package products

import (
	"context"
	"github.com/google/uuid"
)

type ProductStorage interface {
	Add(ctx context.Context, product Product) (Product, error)
	Update(ctx context.Context, product Product) (Product, error)
	Remove(ctx context.Context, uuid2 uuid.UUID) error
}