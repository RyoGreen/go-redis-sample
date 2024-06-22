package controller

import "go-redis/usecase"

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserController() *userHandler {
	return &userHandler{
		userUsecase: usecase.NewUserUsecase(),
	}
}
