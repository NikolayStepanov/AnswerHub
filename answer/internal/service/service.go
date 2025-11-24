package service

import (
	"context"

	"github.com/NikolayStepanov/AnswerHub/internal/domain/dto"
	"github.com/google/uuid"
)

type Answer interface {
	Get(ctx context.Context, answerID int64) (dto.AnswerDTO, error)
	Create(ctx context.Context, questionID int64, userID uuid.UUID, text string) (dto.BaseDTO, error)
	Delete(ctx context.Context, answerID int64) error
	GetAnswersByQuestionID(ctx context.Context, questionID int64) ([]dto.AnswerDTO, error)
}

type Question interface {
	Get(ctx context.Context, questionID int64) (dto.QuestionDTO, error)
	Create(ctx context.Context, text string) (dto.BaseDTO, error)
	List(ctx context.Context) ([]dto.QuestionDTO, error)
	Delete(ctx context.Context, questionID int64) error
}
type Services struct {
	Answer   Answer
	Question Question
}

func NewServices(answer Answer, question Question) *Services {
	return &Services{
		answer,
		question,
	}
}
