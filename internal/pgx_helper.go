package pgx_helper

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Stmt struct {
	Ctx context.Context
	Sql string
	Args []any
}

func ExecStmt(pool *pgxpool.Pool, stmt Stmt) error {
	conn, err := pool.Acquire(stmt.Ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(stmt.Ctx, stmt.Sql, stmt.Args...)

	if err != nil {
		return err
	}

	return nil
}

type QueryStmt[T any] struct {
	Ctx context.Context
	Sql string
	Args []any
	Fn pgx.RowToFunc[T]
}

func ExecQuery[T any](pool *pgxpool.Pool, stmt QueryStmt[T]) ([]T, error) {
	conn, err := pool.Acquire(stmt.Ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(stmt.Ctx, stmt.Sql, stmt.Args...)

	if err != nil {
		return nil, err
	}

	products, err := pgx.CollectRows(rows, stmt.Fn)
	if err != nil {
		return nil, err
	}

	return products, nil
}

