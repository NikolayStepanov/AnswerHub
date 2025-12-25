package answer_rep

import (
	"context"

	"github.com/NikolayStepanov/AnswerHub/internal/domain"
	"github.com/NikolayStepanov/AnswerHub/internal/infrastructure/postgres"
	"github.com/google/uuid"
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
	err := r.db.First(&answer, "id = ?", id).Error
	if err != nil {
		return domain.Answer{}, err
	}
	return domain.ToAnswer(answer), nil
}

func (r *Repository) Create(ctx context.Context, questionID int64, userID uuid.UUID, text string) (domain.Base, error) {
	return domain.Base{}, nil
}

func (r *Repository) Delete(ctx context.Context, answerID int64) error {
	//TODO implement me
	panic("implement me")
}
