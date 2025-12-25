package domain

import "github.com/NikolayStepanov/AnswerHub/internal/infrastructure/postgres"

func ToAnswer(dbA postgres.GORMAnswer) Answer {
	return Answer{
		Base: Base{
			ID:        dbA.ID,
			CreatedAt: dbA.CreatedAt,
		},
		QuestionID: dbA.QuestionID,
		UserID:     dbA.UserID,
		Text:       dbA.Text,
	}
}
