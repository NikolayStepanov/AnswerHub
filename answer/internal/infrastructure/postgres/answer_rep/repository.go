package answer_rep

import (
	"context"

	"github.com/NikolayStepanov/AnswerHub/internal/domain"
	"github.com/NikolayStepanov/AnswerHub/internal/infrastructure/postgres"
	"github.com/NikolayStepanov/AnswerHub/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Get(ctx context.Context, id int64) (domain.Answer, error) {
	var answer postgres.GORMAnswer
	err := r.db.WithContext(ctx).First(&answer, "id = ?", id).Error
	if err != nil {
		logger.Warn("failed to get answer", zap.Error(err))
		return domain.Answer{}, err
	}
	return domain.ToAnswer(answer), nil
}

func (r *Repository) Create(ctx context.Context, questionID int64, userID uuid.UUID, text string) (domain.Base, error) {
	answer := postgres.GORMAnswer{
		QuestionID: questionID,
		UserID:     userID,
		Text:       text,
	}
	err := r.db.WithContext(ctx).Create(&answer).Error
	if err != nil {
		logger.Warn("answer creation failed", zap.Error(err))
		return domain.Base{}, err
	}
	return domain.Base{ID: answer.ID}, nil
}

func (r *Repository) Delete(ctx context.Context, answerID int64) error {
	res := r.db.WithContext(ctx).Delete(&postgres.GORMAnswer{}, answerID)
	if res.Error != nil {
		logger.Warn("answer deletion failed", zap.Error(res.Error))
		return res.Error
	}

	if res.RowsAffected == 0 {
		logger.Warn("answer deletion failed", zap.Error(res.Error))
	}
	return nil
}
