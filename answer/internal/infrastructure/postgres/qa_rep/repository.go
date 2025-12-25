package qa_rep

import (
	"context"

	"github.com/NikolayStepanov/AnswerHub/internal/domain"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r Repository) GetQuestionWithAnswers(ctx context.Context, questionID int64) (domain.QA, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
