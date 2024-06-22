package usecase

import (
	"go-redis/repository"
)

type EntryUsecase interface {
}

type EntryUsecaseImpl struct {
	EntryRepo repository.EntryRepository
}

func NewEntryUsecase() EntryUsecase {
	return &EntryUsecaseImpl{
		EntryRepo: repository.NewEntryRepository(),
	}
}
