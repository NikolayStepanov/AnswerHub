package qa_rep

import (
	"context"

	"github.com/NikolayStepanov/AnswerHub/internal/domain"
	"github.com/NikolayStepanov/AnswerHub/internal/infrastructure/postgres"
	"github.com/NikolayStepanov/AnswerHub/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r Repository) GetQuestionWithAnswers(ctx context.Context, questionID int64) (domain.QA, error) {
	var question postgres.GORMQuestion
	err := r.db.WithContext(ctx).First(&question, "id = ?", questionID).Error
	if err != nil {
		logger.Warn("question not found", zap.Error(err))
		return domain.QA{}, err
	}

	var answers []postgres.GORMAnswer
	err = r.db.WithContext(ctx).Where("question_id = ?", questionID).Find(&answers).Error
	if err != nil {
		return domain.QA{}, err
	}
	qa := domain.QA{
		Question: domain.ToQuestion(question),
		Answers:  domain.ToAnswers(answers),
	}
	return qa, nil
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
