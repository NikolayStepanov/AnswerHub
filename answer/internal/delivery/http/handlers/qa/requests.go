package qa

import "github.com/google/uuid"

type CreateAnswerRequest struct {
	UserID uuid.UUID `json:"user_id"`
	Text   string    `json:"text"`
}
