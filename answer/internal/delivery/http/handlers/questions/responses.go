package questions

import "github.com/NikolayStepanov/AnswerHub/internal/domain/dto"

type createResponse struct {
	BaseDTO dto.BaseDTO
}

type List struct {
	Questions []dto.QuestionDTO
}
