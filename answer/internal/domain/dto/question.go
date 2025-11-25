package dto

type QuestionDTO struct {
	BaseDTO
	Text string `json:"text"`
}

type QuestionResponseDTO struct {
	QuestionDTO
	AnswersResponseDTO AnswersResponseDTO `json:"answers"`
}
