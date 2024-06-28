package domain

import "time"

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
	Create(job *Job) (*Job, error)
	Update(job *Job) (*Job, error)
	Delete(ids []int) error
}

func NewJob(id int, name, description string) *Job {
	return &Job{
		ID:          id,
		Name:        name,
		Description: description,
	}
}
