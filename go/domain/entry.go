package domain

import "time"

type Entry struct {
	ID        int
	UserID    int
	JobID     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
