package postgres

import (
	"time"

	"github.com/google/uuid"
)

type GORMAnswer struct {
	ID         int64     `gorm:"primaryKey"`
	QuestionID int64     `gorm:"not null;index"`
	UserID     uuid.UUID `gorm:"type:uuid;not null"`
	Text       string    `gorm:"type:text;not null"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
type GORMQuestion struct {
	ID        int64        `gorm:"primaryKey"`
	Text      string       `gorm:"type:text;not null"`
	CreatedAt time.Time    `gorm:"default:now()"`
	Answers   []GORMAnswer `gorm:"foreignKey:QuestionID"`
}

func (GORMQuestion) TableName() string { return "questions" }
func (GORMAnswer) TableName() string   { return "answers" }
