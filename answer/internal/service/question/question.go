package question

import (
	"context"
	"time"

	"github.com/NikolayStepanov/AnswerHub/internal/domain/dto"
)

type Service struct {
}

func (s *Service) Get(ctx context.Context, questionID int64) (dto.QuestionDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Create(ctx context.Context, text string) (dto.BaseDTO, error) {
	baseDTO := dto.BaseDTO{
		ID:        0,
		CreatedAt: time.Now(),
	}
	return baseDTO, nil
}

func (s *Service) List(ctx context.Context) ([]dto.QuestionDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Delete(ctx context.Context, questionID int64) error {
	return nil
}

func NewService() *Service {
	return &Service{}
}
