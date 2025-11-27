package qa_rep

import (
	"context"

	"github.com/NikolayStepanov/AnswerHub/internal/domain"
)

type Repository struct{}

func (r Repository) GetQuestionWithAnswers(ctx context.Context, questionID int64) (domain.QA, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository() *Repository {
	return &Repository{}
}
