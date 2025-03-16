// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	Delete(ctx context.Context, id uuid.UUID) error
	Get(ctx context.Context, name string) ([]Product, error)
	GetCount(ctx context.Context, name string) (int64, error)
	Insert(ctx context.Context, arg InsertParams) error
	Update(ctx context.Context, arg UpdateParams) error
}

var _ Querier = (*Queries)(nil)
