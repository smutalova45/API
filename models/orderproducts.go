package models

import "github.com/google/uuid"

type Orderproducts struct {
	Id        uuid.UUID
	OrderId   uuid.UUID
	ProductId uuid.UUID
	Quantity  int
	Price     int
}
