package domain

import "time"

type Entry struct {
	ID        int
	UserID    int
	JobID     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EntryRepository interface {
	GetAll() ([]Entry, error)
	GetByID(id int) (Entry, error)
	Create(e *Entry) error
	Update(e *Entry) error
	Delete(id int) error
}

type EntryCache interface {
	GetAll() ([]Entry, error)
	GetByID(id int) (Entry, error)
	Create(e *Entry) error
	Update(e *Entry) error
	Delete(id int) error
}
