package domain

import (
	"context"
	"time"
)

type Entry struct {
	ID        int
	UserID    int
	JobID     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewEntry(id, userID, jobID int) *Entry {
	return &Entry{
		ID:     id,
		UserID: userID,
		JobID:  jobID,
	}
}

type EntryCache interface {
	GetAll(ctx context.Context) ([]*Entry, error)
	GetByID(ctx context.Context, id int) (*Entry, error)
	Update(ctx context.Context) error
}

type EntryRepository interface {
	FindAll() ([]*Entry, error)
	Find(id int) (*Entry, error)
	Create(ctx context.Context, entry *Entry) (*Entry, error)
	Update(ctx context.Context, entry *Entry) (*Entry, error)
	Delete(ctx context.Context, ids []int) error
}
