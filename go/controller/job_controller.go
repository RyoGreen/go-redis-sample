package controller

import (
	"encoding/json"
	"go-redis/controller/in"
	"go-redis/usecase"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type jobHandler struct {
	jobUsecase usecase.JobUsecase
}

func NewJobController() *jobHandler {
	return &jobHandler{
		jobUsecase: usecase.NewJobUsecase(),
	}
}

func (h *jobHandler) List(w http.ResponseWriter, r *http.Request) {
	jobs, err := h.jobUsecase.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jobJSON, err := json.Marshal(jobs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jobJSON)

}

func (h *jobHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	job, err := h.jobUsecase.Get(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jobJSON, err := json.Marshal(job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jobJSON)
}

func (h *jobHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &in.JobCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	job, err := h.jobUsecase.Create(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jobJSON, err := json.Marshal(job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jobJSON)
}

func (h *jobHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &in.JobUpdateRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	job, err := h.jobUsecase.Update(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jobJSON, err := json.Marshal(job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jobJSON)
}

func (h *jobHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &in.JobDeleteRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.jobUsecase.Delete(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"success"}`))
}
