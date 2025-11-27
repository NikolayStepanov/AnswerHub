package qa

import "github.com/google/uuid"

type CreateAnswerRequest struct {
	QuestionID int64     `json:"question_id"`
	UserID     uuid.UUID `json:"user_id"`
	Text       string    `json:"text"`
}
