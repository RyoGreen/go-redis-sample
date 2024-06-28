package usecase

import (
	"go-redis/controller/out"
	"go-redis/domain"
	"go-redis/infrastructure/postgres"
)

type EntryUsecase interface {
	ListEntries() ([]*out.EntryResponse, error)
	GetEntry(id int) (*out.EntryResponse, error)
	UpdateAllEntries() error
}

type EntryUsecaseImpl struct {
	EntryCacheRepo domain.EntryCache
	UserRepo       domain.UserRepository
	JobRepo        domain.JobRepository
}

func NewEntryUsecase() EntryUsecase {
	return &EntryUsecaseImpl{
		EntryCacheRepo: domain.NewEntryCache(),
		UserRepo:       postgres.NewUserPostgresRepository(),
		JobRepo:        postgres.NewJobPostgresRepository(),
	}
}

func (u *EntryUsecaseImpl) ListEntries() ([]*out.EntryResponse, error) {
	entries, err := u.EntryCacheRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var entryResponses = make([]*out.EntryResponse, 0, len(entries))
	for _, entry := range entries {
		entryResponses = append(entryResponses, &out.EntryResponse{
			ID:        entry.ID,
			UserID:    entry.UserID,
			JobID:     entry.JobID,
			CreatedAt: entry.CreatedAt,
			UpdatedAt: entry.UpdatedAt,
		})
	}
	return nil, nil
}

func (u *EntryUsecaseImpl) GetEntry(id int) (*out.EntryResponse, error) {
	entry, err := u.EntryCacheRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	entryResponse := &out.EntryResponse{
		ID:        entry.ID,
		UserID:    entry.UserID,
		JobID:     entry.JobID,
		CreatedAt: entry.CreatedAt,
		UpdatedAt: entry.UpdatedAt,
	}
	return entryResponse, nil
}

func (u *EntryUsecaseImpl) UpdateAllEntries() error {
	return u.EntryCacheRepo.Update()
}
