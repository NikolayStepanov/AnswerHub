package questions

import "github.com/NikolayStepanov/AnswerHub/internal/domain"

func ToQuestionResponse(q domain.Question) QuestionResponse {
	return QuestionResponse{
		ID:        q.ID,
		Text:      q.Text,
		CreatedAt: q.CreatedAt,
	}
}

func ToQuestionList(questions []domain.Question) []QuestionResponse {
	resp := make([]QuestionResponse, 0, len(questions))
	for _, q := range questions {
		resp = append(resp, ToQuestionResponse(q))
	}
	return resp
}

func ToBaseResponse(base domain.Base) BaseResponse {
	return BaseResponse{
		ID:        base.ID,
		CreatedAt: base.CreatedAt,
	}
}
