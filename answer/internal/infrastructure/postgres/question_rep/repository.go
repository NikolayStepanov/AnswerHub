package question_rep

import (
	"context"

	"github.com/NikolayStepanov/AnswerHub/internal/domain"
)

type Repository struct {
}

func (r Repository) Get(ctx context.Context, questionID int64) (domain.Question, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Exists(ctx context.Context, questionID int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Create(ctx context.Context, text string) (domain.Base, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) List(ctx context.Context) ([]domain.Question, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Delete(ctx context.Context, questionID int64) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository() *Repository {
	return &Repository{}
}
