package qa

import (
	"time"

	"github.com/NikolayStepanov/AnswerHub/internal/delivery/http/handlers/answers"
)

type QuestionResponse struct {
	ID        int64                    `json:"id"`
	Text      string                   `json:"text"`
	CreatedAt time.Time                `json:"created_at"`
	Answers   []answers.AnswerResponse `json:"answers"`
}

type BaseResponse struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
