package question

import (
	"context"

	"github.com/NikolayStepanov/AnswerHub/internal/domain/dto"
	"go.uber.org/zap"
)

type Service struct {
	log *zap.Logger
}

func (s Service) Get(ctx context.Context, questionID int64) (dto.QuestionDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Create(ctx context.Context, text string) (dto.BaseDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) List(ctx context.Context) ([]dto.QuestionDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Delete(ctx context.Context, questionID int64) error {
	//TODO implement me
	panic("implement me")
}

func NewService(log *zap.Logger) *Service {
	return &Service{log}
}
