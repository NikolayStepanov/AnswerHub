package questions

import (
	"time"
)

type BaseResponse struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
type List struct {
	Questions []QuestionResponse `json:"questions"`
}
type QuestionResponse struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
