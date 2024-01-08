package models

import (
	"time"

	"github.com/google/uuid"
)

type Orders struct {
	Id        uuid.UUID
	Amount    int
	CreatedAt time.Time
	UserId    uuid.UUID
}
