package qa

import (
	"github.com/NikolayStepanov/AnswerHub/internal/delivery/http/handlers/answers"
	"github.com/NikolayStepanov/AnswerHub/internal/domain"
)

func ToQuestionResponse(qa domain.QA) QuestionResponse {
	return QuestionResponse{
		ID:        qa.Question.ID,
		Text:      qa.Question.Text,
		CreatedAt: qa.Question.CreatedAt,
		Answers:   toAnswerResponses(qa.Answers),
	}
}

func toAnswerResponses(items []domain.Answer) []answers.AnswerResponse {
	res := make([]answers.AnswerResponse, len(items))

	for i, a := range items {
		res[i] = answers.AnswerResponse{
			ID:        a.ID,
			Text:      a.Text,
			UserID:    a.UserID,
			CreatedAt: a.CreatedAt,
		}
	}

	return res
}

func ToBaseResponse(base domain.Base) BaseResponse {
	return BaseResponse{
		ID:        base.ID,
		CreatedAt: base.CreatedAt,
	}
}
