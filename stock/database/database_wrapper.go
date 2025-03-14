package database

import "github.com/jackc/pgx/v5/pgxpool"

func NewStockDatabase(p *pgxpool.Pool) *Queries {
	return New(p)
}
