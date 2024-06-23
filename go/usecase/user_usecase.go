package usecase

import (
	"go-redis/domain"
	"go-redis/infrastructure/postgres"
)

type UserUsecase interface {
}

type UserUsecaseImpl struct {
	userRepo domain.UserRepository
}

func NewUserUsecase() UserUsecase {
	return &UserUsecaseImpl{
		userRepo: postgres.NewUserPostgresRepository(),
	}
}
