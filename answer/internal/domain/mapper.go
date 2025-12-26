package domain

import (
	"github.com/NikolayStepanov/AnswerHub/internal/infrastructure/postgres"
)

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

func ToQuestion(dbQ postgres.GORMQuestion) Question {
	return Question{
		Base: Base{
			ID:        dbQ.ID,
			CreatedAt: dbQ.CreatedAt,
		},
		Text: dbQ.Text,
	}
}

func ToQuestions(dbQuestions []postgres.GORMQuestion) []Question {
	questions := make([]Question, 0, len(dbQuestions))
	for _, question := range dbQuestions {
		questions = append(questions, Question{
			Base: Base{
				ID:        question.ID,
				CreatedAt: question.CreatedAt,
			},
			Text: question.Text,
		})
	}
	return questions
}
