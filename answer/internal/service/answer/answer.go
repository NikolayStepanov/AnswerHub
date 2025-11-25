package answer

import (
	"context"
	"time"

	"github.com/NikolayStepanov/AnswerHub/internal/domain/dto"
	"github.com/google/uuid"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Get(ctx context.Context, answerID int64) (dto.AnswerDTO, error) {
	return dto.AnswerDTO{
		BaseDTO:    dto.BaseDTO{0, time.Now()},
		QuestionID: 1,
		UserID:     uuid.New(),
		Text:       "test",
	}, nil
}

func (s *Service) Create(ctx context.Context, questionID int64, userID uuid.UUID, text string) (dto.BaseDTO, error) {
	return dto.BaseDTO{ID: 5, CreatedAt: time.Now()}, nil
}

func (s *Service) Delete(ctx context.Context, answerID int64) error {
	return nil
}

func (s *Service) GetAnswersByQuestionID(ctx context.Context, questionID int64) ([]dto.AnswerDTO, error) {
	return []dto.AnswerDTO{
		{
			BaseDTO:    dto.BaseDTO{0, time.Now()},
			QuestionID: 1,
			UserID:     uuid.New(),
			Text:       "test",
		},
	}, nil
}

func (s *Service) DeleteAnswersByQuestionID(ctx context.Context, questionID int64) error {
	return nil
}
