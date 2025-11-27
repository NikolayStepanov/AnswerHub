package answer_rep

import (
	"context"

	"github.com/NikolayStepanov/AnswerHub/internal/domain"
	"github.com/google/uuid"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Get(ctx context.Context, id int64) (domain.Answer, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Create(ctx context.Context, questionID int64, userID uuid.UUID, text string) (domain.Base, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Delete(ctx context.Context, answerID int64) error {
	//TODO implement me
	panic("implement me")
}
