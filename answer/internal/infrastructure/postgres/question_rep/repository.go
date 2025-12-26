package question_rep

import (
	"context"
	"errors"
	"fmt"

	"github.com/NikolayStepanov/AnswerHub/internal/domain"
	"github.com/NikolayStepanov/AnswerHub/internal/infrastructure/postgres"
	"github.com/NikolayStepanov/AnswerHub/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) Get(ctx context.Context, questionID int64) (domain.Question, error) {
	var question postgres.GORMQuestion
	err := r.db.WithContext(ctx).First(&question, "id = ?", questionID).Error
	if err != nil {
		logger.Warn("question not found", zap.Error(err))
		return domain.Question{}, err
	}
	return domain.ToQuestion(question), nil
}

func (r *Repository) Exists(ctx context.Context, questionID int64) (bool, error) {
	var question postgres.GORMQuestion

	err := r.db.WithContext(ctx).Select("id").Where("id = ?", questionID).First(&question).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Warn("question does not exist", zap.Int64("id", questionID))
		return false, nil
	}

	if err != nil {
		logger.Warn("question does not exist", zap.Error(err))
		return false, err
	}
	return true, nil
}

func (r *Repository) Create(ctx context.Context, text string) (domain.Base, error) {
	question := postgres.GORMQuestion{
		Text: text,
	}

	if err := r.db.WithContext(ctx).Create(&question).Error; err != nil {
		logger.Warn("question create error", zap.Error(err))
		return domain.Base{}, err
	}

	return domain.Base{
		ID: question.ID,
	}, nil
}

func (r *Repository) List(ctx context.Context) ([]domain.Question, error) {
	var questions []postgres.GORMQuestion

	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&questions).Error
	if err != nil {
		logger.Warn("questions not found", zap.Error(err))
		return nil, err
	}
	return domain.ToQuestions(questions), nil
}

func (r *Repository) Delete(ctx context.Context, questionID int64) error {
	res := r.db.WithContext(ctx).Delete(&postgres.GORMQuestion{}, questionID)
	if res.Error != nil {
		logger.Warn("question delete error", zap.Error(res.Error))
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("question with id %d not found", questionID)
	}
	return nil
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
