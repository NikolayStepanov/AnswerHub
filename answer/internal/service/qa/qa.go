package qa

import (
	"context"
	"errors"

	"github.com/NikolayStepanov/AnswerHub/internal/domain"
	"github.com/NikolayStepanov/AnswerHub/internal/repository"
	"github.com/NikolayStepanov/AnswerHub/internal/service"
	"github.com/google/uuid"
)

type Service struct {
	rep      repository.QA
	question service.Question
	answer   service.Answer
}

func NewService(rep repository.QA, answer service.Answer, question service.Question) *Service {
	return &Service{rep, question, answer}
}

func (s *Service) GetQuestionWithAnswers(ctx context.Context, questionID int64) (domain.QA, error) {
	var (
		qa  domain.QA
		err error
	)

	qa, err = s.rep.GetQuestionWithAnswers(ctx, questionID)
	if err != nil {
		return domain.QA{}, err
	}

	return qa, nil
}

func (s *Service) CreateAnswer(ctx context.Context, questionID int64, userID uuid.UUID, text string) (domain.Base, error) {
	var (
		base  domain.Base
		exist bool
		err   error
	)

	exist, err = s.question.Exists(ctx, questionID)
	if err != nil {
		return domain.Base{}, err
	}
	if !exist {
		return domain.Base{}, errors.New("question does not exist")
	}

	base, err = s.answer.Create(ctx, questionID, userID, text)
	if err != nil {
		return domain.Base{}, err
	}
	return base, nil
}
