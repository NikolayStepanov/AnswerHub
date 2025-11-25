package qa

import (
	"context"
	"errors"

	"github.com/NikolayStepanov/AnswerHub/internal/domain/dto"
	"github.com/NikolayStepanov/AnswerHub/internal/service"
	"github.com/google/uuid"
)

type Service struct {
	answer   service.Answer
	question service.Question
}

func (s *Service) GetQuestionWithAnswers(ctx context.Context, questionID int64) (dto.QuestionResponseDTO, error) {
	var (
		questionDTO dto.QuestionDTO
		answersDTO  []dto.AnswerDTO
		err         error
	)
	questionDTO, err = s.question.Get(ctx, questionID)
	if err != nil {
		return dto.QuestionResponseDTO{}, err
	}

	answersDTO, err = s.answer.GetAnswersByQuestionID(ctx, questionID)
	if err != nil {
		return dto.QuestionResponseDTO{}, err
	}
	answersResponses := dto.ToAnswerResponseList(answersDTO)
	return dto.QuestionResponseDTO{
		QuestionDTO:        questionDTO,
		AnswersResponseDTO: answersResponses,
	}, nil
}

func (s *Service) CreateAnswer(ctx context.Context, questionID int64, userID uuid.UUID, text string) (dto.BaseDTO, error) {
	var (
		baseDTO dto.BaseDTO
		exist   bool
		err     error
	)

	exist, err = s.question.Exists(ctx, questionID)
	if err != nil {
		return dto.BaseDTO{}, err
	}
	if !exist {
		return dto.BaseDTO{}, errors.New("question does not exist")
	}

	baseDTO, err = s.answer.Create(ctx, questionID, userID, text)
	if err != nil {
		return dto.BaseDTO{}, err
	}
	return baseDTO, nil
}

func NewService(answer service.Answer, question service.Question) *Service {
	return &Service{answer, question}
}
