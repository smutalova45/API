package models

import "github.com/google/uuid"

type Orderproducts struct {
	Id        uuid.UUID `json:"id"`
	OrderId   uuid.UUID `json:"orderid"`
	ProductId uuid.UUID `json:"productid"`
	Quantity  int       `json:"quantity"`
	Price     int       `json:"price"`
}
