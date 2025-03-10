package stock

import (
	"context"

	pgxh "github.com/Revazashvili/ecommerce-inventory-management/internal"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresStorage struct {
	pool *pgxpool.Pool
}

func NewStockStorage(p *pgxpool.Pool) Storage {
	return &postgresStorage{
		pool: p,
	}
}

func (ps *postgresStorage) Get(ctx context.Context, productID uuid.UUID) (Stock, error) {
	var s Stock
	s, err := pgxh.ExecQueryOneStruct[Stock](ps.pool, pgxh.Stmt{
		Ctx:  ctx,
		Sql:  "select id, product_id, quantity, reserved_quantity, version, create_date, last_update_date from products.stocks where product_id = $1",
		Args: []any{productID},
	})

	if err != nil {
		return s, err
	}

	return s, nil
}

func (ps *postgresStorage) ReserveStock(ctx context.Context, s Stock) (Stock, error) {
	err := pgxh.ExecStmt(ps.pool, pgxh.Stmt{
		Ctx:  ctx,
		Sql:  "update products.stocks set reserved_quantity=$1, version = version+1 where id=$2 and version=$3",
		Args: []any{s.ReservedQuantity, s.Id, s.Version},
	})

	if err != nil {
		return s, err
	}

	return s, nil
}

func (ps *postgresStorage) AddStockReservation(ctx context.Context, sr StockReservation) (StockReservation, error) {
	err := pgxh.ExecStmt(ps.pool, pgxh.Stmt{
		Ctx:  ctx,
		Sql:  "insert into products.stock_reservations (id, product_id, order_number, quantity, create_date, cancel_date) values ($1, $2, $3, $4, $5, $6)",
		Args: []any{sr.Id, sr.ProductId, sr.OrderNumber, sr.Quantity, sr.CreateDate, sr.CancelDate},
	})

	if err != nil {
		return sr, err
	}

	return sr, nil
}
