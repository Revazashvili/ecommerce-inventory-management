package pgx_helper

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Stmt struct {
	Ctx  context.Context
	Sql  string
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

func ExecQueryStructs[T any](pool *pgxpool.Pool, stmt Stmt) ([]T, error) {
	conn, err := pool.Acquire(stmt.Ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(stmt.Ctx, stmt.Sql, stmt.Args...)

	if err != nil {
		return nil, err
	}

	result, err := pgx.CollectRows(rows, pgx.RowToStructByName[T])
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ExecQueryOne[T any](pool *pgxpool.Pool, stmt Stmt) (T, error) {
	var result T
	conn, err := pool.Acquire(stmt.Ctx)
	if err != nil {
		return result, err
	}
	defer conn.Release()

	rows, err := conn.Query(stmt.Ctx, stmt.Sql, stmt.Args...)

	if err != nil {
		return result, err
	}

	result, err = pgx.CollectOneRow(rows, pgx.RowTo[T])
	if err != nil {
		return result, err
	}

	return result, nil
}
