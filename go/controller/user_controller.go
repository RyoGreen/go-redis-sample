package controller

import (
	"go-redis/usecase"
	"net/http"
)

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserController() *userHandler {
	return &userHandler{
		userUsecase: usecase.NewUserUsecase(),
	}
}

func (h *userHandler) List(w http.ResponseWriter, r *http.Request) {}
