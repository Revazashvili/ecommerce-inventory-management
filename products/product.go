package products

import "github.com/google/uuid"

type Product struct {
	Id   uuid.UUID
	Name string
}