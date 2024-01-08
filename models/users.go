package models

import "github.com/google/uuid"

type Users struct {
	Id        uuid.UUID
	Firstname string
	Email     string
	Phone     string
}
