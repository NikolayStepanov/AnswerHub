package answer

import (
	"context"

	"github.com/NikolayStepanov/AnswerHub/internal/domain"
	"github.com/NikolayStepanov/AnswerHub/internal/repository"
	"github.com/google/uuid"
)

type Service struct {
	rep repository.Answer
}

func NewService(rep repository.Answer) *Service {
	return &Service{rep: rep}
}

func (s *Service) Get(ctx context.Context, answerID int64) (domain.Answer, error) {
	var (
		answer domain.Answer
		err    error
	)
	answer, err = s.rep.Get(ctx, answerID)
	if err != nil {
		return domain.Answer{}, err
	}
	return answer, nil
}

func (s *Service) Create(ctx context.Context, questionID int64, userID uuid.UUID, text string) (domain.Base, error) {
	var (
		base domain.Base
		err  error
	)
	base, err = s.rep.Create(ctx, questionID, userID, text)
	if err != nil {
		return domain.Base{}, err
	}
	return base, nil
}

func (s *Service) Delete(ctx context.Context, answerID int64) error {
	return s.rep.Delete(ctx, answerID)
}
