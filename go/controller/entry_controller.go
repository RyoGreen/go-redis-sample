package controller

import (
	"encoding/json"
	"go-redis/controller/in"
	"go-redis/usecase"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type entryHandler struct {
	entryUsecase usecase.EntryUsecase
}

func NewEntryController() *entryHandler {
	return &entryHandler{
		entryUsecase: usecase.NewEntryUsecase(),
	}
}

func (h *entryHandler) List(w http.ResponseWriter, r *http.Request) {
	entries, err := h.entryUsecase.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	entryJSON, err := json.Marshal(entries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(entryJSON)
}

func (h *entryHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	entry, err := h.entryUsecase.Get(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	entryJSON, err := json.Marshal(entry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(entryJSON)
}

func (h *entryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var entryRequest *in.EntryCreateRequest
	err := json.NewDecoder(r.Body).Decode(&entryRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	entry, err := h.entryUsecase.Create(r.Context(), entryRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	entryJSON, err := json.Marshal(entry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(entryJSON)
}

func (h *entryHandler) Update(w http.ResponseWriter, r *http.Request) {
	var entryRequest *in.EntryUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&entryRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	entry, err := h.entryUsecase.Update(r.Context(), entryRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	entryJSON, err := json.Marshal(entry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(entryJSON)
}

func (h *entryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var entryRequest *in.EntryDeleteRequest
	err := json.NewDecoder(r.Body).Decode(&entryRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.entryUsecase.Delete(r.Context(), entryRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "success"}`))
}
