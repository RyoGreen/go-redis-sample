package usecase

import "go-redis/repository"

type JobUsecase interface {
}

type JobUsecaseImpl struct {
	JobRepo repository.JobRepository
}

func NewJobUsecase() JobUsecase {
	return &JobUsecaseImpl{
		JobRepo: repository.NewJobRepository(),
	}
}
