package domain

import (
	"github.com/google/uuid"
)

type Answer struct {
	Base
	QuestionID int64
	UserID     uuid.UUID
	Text       string
}
