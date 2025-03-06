package pgx_helper

import (
	"context"
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

