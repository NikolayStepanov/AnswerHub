package dto

import (
	"time"

	"github.com/google/uuid"
)

type BaseDTO struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
