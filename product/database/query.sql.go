// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const delete = `-- name: Delete :exec
delete from products.products where id=$1
`

func (q *Queries) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, delete, id)
	return err
}

const get = `-- name: Get :many
select id, name from products.products 
where name ilike concat('%', $1::text ,'%')
`

func (q *Queries) Get(ctx context.Context, name string) ([]Product, error) {
	rows, err := q.db.Query(ctx, get, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCount = `-- name: GetCount :one
select count(*) from products.products 
where name ilike concat('%', $1::text ,'%')
`

func (q *Queries) GetCount(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRow(ctx, getCount, name)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const insert = `-- name: Insert :exec
insert into products.products (id, name) values ($1, $2)
`

type InsertParams struct {
	ID   uuid.UUID
	Name string
}

func (q *Queries) Insert(ctx context.Context, arg InsertParams) error {
	_, err := q.db.Exec(ctx, insert, arg.ID, arg.Name)
	return err
}

const update = `-- name: Update :exec
update products.products set name=$1 where id=$2
`

type UpdateParams struct {
	Name string
	ID   uuid.UUID
}

func (q *Queries) Update(ctx context.Context, arg UpdateParams) error {
	_, err := q.db.Exec(ctx, update, arg.Name, arg.ID)
	return err
}
