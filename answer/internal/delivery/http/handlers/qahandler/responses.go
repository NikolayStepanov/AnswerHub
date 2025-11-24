package qahandler

import "github.com/NikolayStepanov/AnswerHub/internal/domain/dto"

type getQuestionResponse struct {
	Question dto.QuestionDTO        `json:"question"`
	Answers  dto.AnswersResponseDTO `json:"answers"`
}
