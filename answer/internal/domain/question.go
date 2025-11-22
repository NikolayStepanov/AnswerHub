package domain

import "time"

type Question struct {
	ID        int64
	Text      string
	CreatedAt time.Time
}
