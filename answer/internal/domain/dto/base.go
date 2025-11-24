package dto

import (
	"time"
)

type BaseDTO struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
