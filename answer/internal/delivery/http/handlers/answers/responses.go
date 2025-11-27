package answers

import (
	"time"

	"github.com/google/uuid"
)

type AnswerResponse struct {
	ID         int64     `json:"id"`
	QuestionID int64     `json:"question_id"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`
	UserID     uuid.UUID `json:"user_id"`
}
