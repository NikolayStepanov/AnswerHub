package repository

import (
	"context"

	"github.com/NikolayStepanov/AnswerHub/internal/domain"
	"github.com/google/uuid"
)

type Answer interface {
	Get(ctx context.Context, id int64) (domain.Answer, error)
	Create(ctx context.Context, questionID int64, userID uuid.UUID, text string) (domain.Base, error)
	Delete(ctx context.Context, answerID int64) error
}

type Question interface {
	Get(ctx context.Context, questionID int64) (domain.Question, error)
	Exists(ctx context.Context, questionID int64) (bool, error)
	Create(ctx context.Context, text string) (domain.Base, error)
	List(ctx context.Context) ([]domain.Question, error)
	Delete(ctx context.Context, questionID int64) error
}

type QA interface {
	GetQuestionWithAnswers(ctx context.Context, questionID int64) (domain.QA, error)
}
type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}
