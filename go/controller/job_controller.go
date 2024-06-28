package controller

import (
	"go-redis/usecase"
	"net/http"
)

type jobHandler struct {
	jobUsecase usecase.JobUsecase
}

func NewJobController() *jobHandler {
	return &jobHandler{
		jobUsecase: usecase.NewJobUsecase(),
	}
}

func (h *jobHandler) List(w http.ResponseWriter, r *http.Request)   {}
func (h *jobHandler) Get(w http.ResponseWriter, r *http.Request)    {}
func (h *jobHandler) Create(w http.ResponseWriter, r *http.Request) {}
func (h *jobHandler) Update(w http.ResponseWriter, r *http.Request) {}
func (h *jobHandler) Delete(w http.ResponseWriter, r *http.Request) {}
