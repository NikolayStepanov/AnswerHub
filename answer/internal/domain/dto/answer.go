package dto

import "github.com/google/uuid"

type AnswerDTO struct {
	BaseDTO
	QuestionID int64     `json:"question_id"`
	UserID     uuid.UUID `json:"user_id"`
	Text       string    `json:"text"`
}

type AnswerResponseDTO struct {
	BaseDTO
	UserID uuid.UUID `json:"user_id"`
	Text   string    `json:"text"`
}

type AnswersResponseDTO []AnswerResponseDTO

func (p AnswersResponseDTO) Len() int {
	return len(p)
}

func (p AnswersResponseDTO) Less(i, j int) bool {
	return p[i].ID < p[j].ID
}

func (p AnswersResponseDTO) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (a AnswerDTO) ToResponse() AnswerResponseDTO {
	return AnswerResponseDTO{
		BaseDTO: a.BaseDTO,
		UserID:  a.UserID,
		Text:    a.Text,
	}
}

func ToAnswerResponseList(answersDTO []AnswerDTO) []AnswerResponseDTO {
	responses := make([]AnswerResponseDTO, len(answersDTO))
	for i, answerDTO := range answersDTO {
		responses[i] = answerDTO.ToResponse()
	}
	return responses
}
