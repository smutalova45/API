package models

import "github.com/google/uuid"

type Products struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price int       `json:"price"`
}
