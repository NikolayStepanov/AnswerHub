package qa

import "github.com/google/uuid"

type createAnswerRequest struct {
	QuestionID int64     `json:"question_id"`
	UserID     uuid.UUID `json:"user_id"`
	Text       string    `json:"text"`
}
