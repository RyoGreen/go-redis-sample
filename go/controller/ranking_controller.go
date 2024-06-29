package controller

import (
	"encoding/json"
	"go-redis/usecase"
	"net/http"
)

type rankingHandler struct {
	rankingUsecase usecase.RankingService
}

func NewRankingController() *rankingHandler {
	return &rankingHandler{
		rankingUsecase: usecase.NewRankingService(),
	}
}

func (h *rankingHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rankings, err := h.rankingUsecase.List(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jobJSON, err := json.Marshal(rankings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jobJSON)

}

func (h *rankingHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := h.rankingUsecase.Update(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
