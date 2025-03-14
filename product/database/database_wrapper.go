package database

import "github.com/jackc/pgx/v5/pgxpool"

func NewProductsDatabase(p *pgxpool.Pool) *Queries {
	return New(p)
}
