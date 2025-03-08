package product

import (
	"context"

	"github.com/google/uuid"
)

type Storage interface {
	Search(ctx context.Context, name string) ([]Product, error)
	Count(ctx context.Context, name string) (int, error)
	Add(ctx context.Context, product Product) (Product, error)
	Update(ctx context.Context, product Product) (Product, error)
	Remove(ctx context.Context, uuid2 uuid.UUID) error
}
