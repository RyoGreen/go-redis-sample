package controller

import (
	"go-redis/usecase"
	"net/http"
)

type entryHandler struct {
	entryUsecase usecase.EntryUsecase
}

func NewEntryController() *entryHandler {
	return &entryHandler{
		entryUsecase: usecase.NewEntryUsecase(),
	}
}

func (h *entryHandler) ListEntries(w http.ResponseWriter, r *http.Request) {

}
