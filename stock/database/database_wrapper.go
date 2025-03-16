package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewStockStorage(p *pgxpool.Pool) *Storage {
	return &Storage{
		Querier: New(p),
		pool:    p,
	}
}

type Storage struct {
	Querier Querier
	pool    *pgxpool.Pool
}

func (s *Storage) ExecWithTx(ctx context.Context, fn func(Querier) error) error {
	q := s.Querier.(*Queries)

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	qtx := q.WithTx(tx)

	err = fn(qtx)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
