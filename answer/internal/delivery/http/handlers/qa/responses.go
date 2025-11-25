package qa

import "github.com/NikolayStepanov/AnswerHub/internal/domain/dto"

type getQuestionResponse struct {
	dto.QuestionResponseDTO
}

type createAnswerResponse struct {
	BaseDTO dto.BaseDTO
}
