package dto

type AnswerDTO struct {
	ID         int64  `json:"id"`
	QuestionID int64  `json:"question_id"`
	UserID     string `json:"user_id"`
	Text       string `json:"text"`
}
