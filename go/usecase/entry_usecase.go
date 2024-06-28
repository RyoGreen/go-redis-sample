package usecase

import (
	"go-redis/domain"
)

type EntryUsecase interface {
}

type EntryUsecaseImpl struct {
	EntryRepo domain.EntryCache
}

func NewEntryUsecase() EntryUsecase {
	return &EntryUsecaseImpl{}
}
