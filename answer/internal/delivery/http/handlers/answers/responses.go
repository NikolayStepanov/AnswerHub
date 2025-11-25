package answers

import "github.com/NikolayStepanov/AnswerHub/internal/domain/dto"

type getAnswerResponse struct {
	Answers dto.AnswerDTO `json:"answer"`
}
