package controller

import "go-redis/usecase"

type jobHandler struct {
	jobUsecase usecase.JobUsecase
}

func NewJobController() *jobHandler {
	return &jobHandler{
		jobUsecase: usecase.NewJobUsecase(),
	}
}
