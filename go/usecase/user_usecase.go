package usecase

import "go-redis/repository"

type UserUsecase interface {
}

type UserUsecaseImpl struct {
	userRepo repository.UserRepository
}

func NewUserUsecase() UserUsecase {
	return &UserUsecaseImpl{
		userRepo: repository.NewUserRepository(),
	}
}
