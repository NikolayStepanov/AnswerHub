package dto

import "github.com/google/uuid"

type AnswerDTO struct {
	BaseDTO
	QuestionID int64     `json:"question_id"`
	UserID     uuid.UUID `json:"user_id"`
	Text       string    `json:"text"`
}
