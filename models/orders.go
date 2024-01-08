package models

import (
	"time"

	"github.com/google/uuid"
)

type Orders struct {
	Id        uuid.UUID `json:"id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"createdat"`
	UserId    uuid.UUID `json:"user_id"`
}
