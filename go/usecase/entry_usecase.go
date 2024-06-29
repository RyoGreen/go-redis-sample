package usecase

import (
	"context"
	"go-redis/controller/in"
	"go-redis/controller/out"
	"go-redis/domain"
	"go-redis/infrastructure/postgres"
)

type EntryUsecase interface {
	List() ([]*out.EntryResponse, error)
	Get(id int) (*out.EntryResponse, error)
	Create(ctx context.Context, entry *in.EntryCreateRequest) (*out.EntryResponse, error)
	Update(ctx context.Context, entry *in.EntryUpdateRequest) (*out.EntryResponse, error)
	Delete(ctx context.Context, ids *in.EntryDeleteRequest) error
}

type EntryUsecaseImpl struct {
	EntryRepo domain.EntryRepository
}

func NewEntryUsecase() EntryUsecase {
	return &EntryUsecaseImpl{
		EntryRepo: postgres.NewEntryPostgresRepository(),
	}
}

func (u *EntryUsecaseImpl) List() ([]*out.EntryResponse, error) {
	entries, err := u.EntryRepo.FindAll()
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
	return entryResponses, nil
}

func (u *EntryUsecaseImpl) Get(id int) (*out.EntryResponse, error) {
	entry, err := u.EntryRepo.Find(id)
	if err != nil {
		return nil, err
	}
	if entry != nil {
		return &out.EntryResponse{
			ID:        entry.ID,
			UserID:    entry.UserID,
			JobID:     entry.JobID,
			CreatedAt: entry.CreatedAt,
			UpdatedAt: entry.UpdatedAt,
		}, nil
	}
	return nil, nil
}

func (u *EntryUsecaseImpl) Create(ctx context.Context, entry *in.EntryCreateRequest) (*out.EntryResponse, error) {
	newEntry := domain.NewEntry(0, entry.UserID, entry.JobID)
	createdEntry, err := u.EntryRepo.Create(ctx, newEntry)
	if err != nil {
		return nil, err
	}
	return &out.EntryResponse{
		ID:        createdEntry.ID,
		UserID:    createdEntry.UserID,
		JobID:     createdEntry.JobID,
		CreatedAt: createdEntry.CreatedAt,
		UpdatedAt: createdEntry.UpdatedAt,
	}, nil
}

func (u *EntryUsecaseImpl) Update(ctx context.Context, entry *in.EntryUpdateRequest) (*out.EntryResponse, error) {
	newEntry := domain.NewEntry(entry.ID, entry.UserID, entry.JobID)
	updatedEntry, err := u.EntryRepo.Update(ctx, newEntry)
	if err != nil {
		return nil, err
	}
	return &out.EntryResponse{
		ID:        updatedEntry.ID,
		UserID:    updatedEntry.UserID,
		JobID:     updatedEntry.JobID,
		CreatedAt: updatedEntry.CreatedAt,
		UpdatedAt: updatedEntry.UpdatedAt,
	}, nil
}

func (u *EntryUsecaseImpl) Delete(ctx context.Context, ids *in.EntryDeleteRequest) error {
	err := u.EntryRepo.Delete(ctx, ids.IDs)
	if err != nil {
		return err
	}
	return nil
}
