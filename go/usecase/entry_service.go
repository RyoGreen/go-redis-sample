package usecase

import (
	"context"
	"go-redis/controller/out"
	"go-redis/domain"
	"go-redis/infrastructure/postgres"
	"go-redis/infrastructure/redis"
)

type EntryService interface {
	ListEntries(ctx context.Context) ([]*out.EntryResponse, error)
	GetEntry(ctx context.Context, id int) (*out.EntryResponse, error)
	UpdateAllEntries(ctx context.Context) error
}

type EntryServiceImpl struct {
	EntryCacheRepo domain.EntryCache
	UserRepo       domain.UserRepository
	JobRepo        domain.JobRepository
}

func NewEntryService() EntryService {
	return &EntryServiceImpl{
		EntryCacheRepo: redis.NewRedisRepository(),
		UserRepo:       postgres.NewUserPostgresRepository(),
		JobRepo:        postgres.NewJobPostgresRepository(),
	}
}

func (u *EntryServiceImpl) ListEntries(ctx context.Context) ([]*out.EntryResponse, error) {
	entries, err := u.EntryCacheRepo.GetAll(ctx)
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

func (u *EntryServiceImpl) GetEntry(ctx context.Context, id int) (*out.EntryResponse, error) {
	entry, err := u.EntryCacheRepo.GetByID(ctx, id)
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

func (u *EntryServiceImpl) UpdateAllEntries(ctx context.Context) error {
	return u.EntryCacheRepo.Update(ctx)
}
