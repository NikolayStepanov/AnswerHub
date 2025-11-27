package question

import (
	"context"

	"github.com/NikolayStepanov/AnswerHub/internal/domain"
	"github.com/NikolayStepanov/AnswerHub/internal/repository"
)

type Service struct {
	rep repository.Question
}

func NewService(rep repository.Question) *Service {
	return &Service{rep: rep}
}

func (s *Service) Get(ctx context.Context, questionID int64) (domain.Question, error) {
	var (
		question domain.Question
		err      error
	)
	question, err = s.rep.Get(ctx, questionID)
	if err != nil {
		return domain.Question{}, err
	}
	return question, nil
}

func (s *Service) Create(ctx context.Context, text string) (domain.Base, error) {
	var (
		base domain.Base
		err  error
	)
	base, err = s.rep.Create(ctx, text)
	if err != nil {
		return domain.Base{}, err
	}
	return base, nil
}

func (s *Service) List(ctx context.Context) ([]domain.Question, error) {
	var (
		questions []domain.Question
		err       error
	)
	questions, err = s.rep.List(ctx)
	if err != nil {
		return []domain.Question{}, err
	}
	return questions, nil
}

func (s *Service) Exists(ctx context.Context, questionID int64) (bool, error) {
	var (
		exist bool
		err   error
	)
	exist, err = s.rep.Exists(ctx, questionID)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (s *Service) Delete(ctx context.Context, questionID int64) error {
	return s.rep.Delete(ctx, questionID)
}
