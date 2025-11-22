package answer

import (
	"context"

	"github.com/NikolayStepanov/AnswerHub/internal/domain/dto"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service struct {
	log *zap.Logger
}

func NewService(log *zap.Logger) *Service {
	return &Service{log}
}

func (s Service) Get(ctx context.Context, answerID int64) (dto.AnswerDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Create(ctx context.Context, questionID int64, userID uuid.UUID, text string) (dto.BaseDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Delete(ctx context.Context, answerID int64) error {
	//TODO implement me
	panic("implement me")
}
