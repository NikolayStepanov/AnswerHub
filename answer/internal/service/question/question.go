package question

import (
	"context"
	"time"

	"github.com/NikolayStepanov/AnswerHub/internal/domain/dto"
)

type Service struct {
}

func (s *Service) Exists(ctx context.Context, questionID int64) (bool, error) {
	return true, nil
}

func (s *Service) Get(ctx context.Context, questionID int64) (dto.QuestionDTO, error) {
	return dto.QuestionDTO{
		BaseDTO: dto.BaseDTO{
			ID:        questionID,
			CreatedAt: time.Now(),
		},
		Text: "question test",
	}, nil
}

func (s *Service) Create(ctx context.Context, text string) (dto.BaseDTO, error) {
	baseDTO := dto.BaseDTO{
		ID:        0,
		CreatedAt: time.Now(),
	}
	return baseDTO, nil
}

func (s *Service) List(ctx context.Context) ([]dto.QuestionDTO, error) {
	return []dto.QuestionDTO{
		{
			BaseDTO: dto.BaseDTO{ID: 1, CreatedAt: time.Now()},
			Text:    "question test1",
		},
		{
			BaseDTO: dto.BaseDTO{ID: 2, CreatedAt: time.Now()},
			Text:    "question test2",
		},
	}, nil
}

func (s *Service) Delete(ctx context.Context, questionID int64) error {
	return nil
}

func NewService() *Service {
	return &Service{}
}
