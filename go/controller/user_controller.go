package controller

import (
	"encoding/json"
	"go-redis/controller/in"
	"go-redis/usecase"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserController() *userHandler {
	return &userHandler{
		userUsecase: usecase.NewUserUsecase(),
	}
}

func (h *userHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUsecase.List()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	usersJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(usersJSON)
}

func (h *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	user, err := h.userUsecase.Get(idInt)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	userJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req *in.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdUser, err := h.userUsecase.Create(ctx, req)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	userJSON, err := json.Marshal(createdUser)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}

func (h *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req *in.UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	updatedUser, err := h.userUsecase.Update(ctx, req)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	userJSON, err := json.Marshal(updatedUser)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}

func (h *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req *in.DeleteUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = h.userUsecase.Delete(ctx, req)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
