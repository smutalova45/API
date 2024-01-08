package models

import "github.com/google/uuid"

type Users struct {
	Id        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
}
