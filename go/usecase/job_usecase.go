package usecase

import "go-redis/domain"

type JobUsecase interface {
}

type JobUsecaseImpl struct {
	JobRepo domain.JobRepository
}

func NewJobUsecase() JobUsecase {
	return &JobUsecaseImpl{}
}
