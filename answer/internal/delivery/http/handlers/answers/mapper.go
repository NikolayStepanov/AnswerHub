package answers

import "github.com/NikolayStepanov/AnswerHub/internal/domain"

func ToAnswerResponse(a domain.Answer) AnswerResponse {
	return AnswerResponse{
		ID:         a.ID,
		QuestionID: a.QuestionID,
		Text:       a.Text,
		CreatedAt:  a.CreatedAt,
		UserID:     a.UserID,
	}
}
