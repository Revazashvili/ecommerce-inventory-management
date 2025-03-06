package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

const dbURL = "postgres://user:pass@localhost:5432/products"

func main() {

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		fmt.Println(err)
	}

	storage := NewProductStorage(pool)
	p := Product{
		Id:   uuid.New(),
		Name: "test",
	}
	p, err = storage.Add(ctx, p)

	if err != nil {
		fmt.Println(err)
	}

	p.Name = "test updated"
	p, err = storage.Update(ctx, p)

	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second * 5)
	err = storage.Remove(ctx, p.Id)

	if err != nil {
		fmt.Println(err)
	}
}

type Product struct {
	Id   uuid.UUID
	Name string
}

type ProductStorage interface {
	Add(ctx context.Context, product Product) (Product, error)
	Update(ctx context.Context, product Product) (Product, error)
	Remove(ctx context.Context, uuid2 uuid.UUID) error
}

type PostgresProductStorage struct {
	pool *pgxpool.Pool
}

func NewProductStorage(p *pgxpool.Pool) ProductStorage {
	return &PostgresProductStorage{
		pool: p,
	}
}

func (pps *PostgresProductStorage) Add(ctx context.Context, product Product) (Product, error) {
	conn, err := pps.pool.Acquire(ctx)
	if err != nil {
		return Product{}, err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "insert into products.products (id, name) values ($1, $2)", product.Id, product.Name)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (pps *PostgresProductStorage) Update(ctx context.Context, product Product) (Product, error) {
	conn, err := pps.pool.Acquire(ctx)
	if err != nil {
		return Product{}, err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "update products.products set name=$1 where id=$2", product.Name, product.Id)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (pps *PostgresProductStorage) Remove(ctx context.Context, id uuid.UUID) error {
	conn, err := pps.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "delete from products.products where id=$1", id)

	if err != nil {
		return err
	}

	return nil
}
