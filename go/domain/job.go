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
	FindAll() ([]Job, error)
	FindById(id int) (Job, error)
	Create(job Job) (Job, error)
	Update(job Job) (Job, error)
	Delete(id int) error
}
