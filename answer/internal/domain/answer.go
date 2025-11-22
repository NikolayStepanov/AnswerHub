package domain

import (
	"time"

	"github.com/google/uuid"
)

type Answer struct {
	ID         int64
	QuestionID int64
	UserID     uuid.UUID
	Text       string
	CreatedAt  time.Time
}
