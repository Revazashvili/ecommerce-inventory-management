package product

import (
    "context"
	pgxh "github.com/Revazashvili/ecommerce-inventory-management/internal"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresProductStorage struct {
	pool *pgxpool.Pool
}

func NewProductStorage(p *pgxpool.Pool) ProductStorage {
	return &PostgresProductStorage{
		pool: p,
	}
}

func (pps *PostgresProductStorage) Search(ctx context.Context, name string) ([]Product, error) {
	products, err := pgxh.ExecQueryStructs[Product](pps.pool, pgxh.Stmt{
		Ctx: ctx,
		Sql: "select id, name from products.products where name ilike $1",
		Args: []any{ "%" + name + "%" },
	})

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (pps *PostgresProductStorage) Add(ctx context.Context, product Product) (Product, error) {
	err := pgxh.ExecStmt(pps.pool, pgxh.Stmt{
		Ctx:  ctx,
		Sql:  "insert into products.products (id, name) values ($1, $2)",
		Args: []any{product.Id, product.Name},
	})

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (pps *PostgresProductStorage) Update(ctx context.Context, product Product) (Product, error) {
	err := pgxh.ExecStmt(pps.pool, pgxh.Stmt{
		Ctx:  ctx,
		Sql:  "update products.products set name=$1 where id=$2",
		Args: []any{product.Name, product.Id},
	})
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (pps *PostgresProductStorage) Remove(ctx context.Context, id uuid.UUID) error {
	err := pgxh.ExecStmt(pps.pool, pgxh.Stmt{
		Ctx:  ctx,
		Sql:  "delete from products.products where id=$1",
		Args: []any{id},
	})

	if err != nil {
		return err
	}

	return nil
}