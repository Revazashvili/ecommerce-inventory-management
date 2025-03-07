package main

import (
	"context"
	"fmt"
	"github.com/Revazashvili/ecommerce-inventory-management/products"
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

	storage := products.NewProductStorage(pool)
	p := products.Product{
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