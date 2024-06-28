package domain

import (
	"context"
	"time"
)

type Job struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type JobRepository interface {
	FindAll() ([]*Job, error)
	FindById(id int) (*Job, error)
	Create(ctx context.Context, job *Job) (*Job, error)
	Update(ctx context.Context, job *Job) (*Job, error)
	Delete(ctx context.Context, ids []int) error
}

func NewJob(id int, name, description string) *Job {
	return &Job{
		ID:          id,
		Name:        name,
		Description: description,
	}
}
