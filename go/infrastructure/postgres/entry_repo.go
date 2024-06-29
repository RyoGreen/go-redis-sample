package postgres

import (
	"context"
	"go-redis/domain"
	"time"
)

func NewEntryPostgresRepository() domain.EntryRepository {
	return &EntryRepositoryImpl{}
}

type EntryRepositoryImpl struct{}

func (r *EntryRepositoryImpl) FindAll() ([]*domain.Entry, error) {
	sql := "SELECT id, user_id, job_id, created_at, updated_at FROM entries"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var entries []*domain.Entry
	for rows.Next() {
		var entry domain.Entry
		err := rows.Scan(&entry.ID, &entry.UserID, &entry.JobID, &entry.CreatedAt, &entry.UpdatedAt)
		if err != nil {
			return nil, err
		}
		entries = append(entries, &entry)
	}
	return entries, nil
}

func (r *EntryRepositoryImpl) Find(id int) (*domain.Entry, error) {
	sql := "SELECT id, user_id, job_id, created_at, updated_at FROM entries WHERE id = $1"
	row := db.QueryRow(sql, id)
	var entry domain.Entry
	err := row.Scan(&entry.ID, &entry.UserID, &entry.JobID, &entry.CreatedAt, &entry.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (r *EntryRepositoryImpl) Create(ctx context.Context, entry *domain.Entry) (*domain.Entry, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	sql := `INSERT INTO entries (name, description,created_at,updated_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	now := time.Now()
	row := tx.QueryRow(sql, entry.UserID, entry.JobID, now, now)
	err = row.Scan(&entry.ID, &entry.CreatedAt, &entry.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *EntryRepositoryImpl) Update(ctx context.Context, entry *domain.Entry) (*domain.Entry, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	sql := `UPDATE entries SET user_id = $1, job_id = $2, updated_at = $3 WHERE id = $4 RETURNING id, user_id , job_id, created_at, updated_at`
	now := time.Now()
	row := tx.QueryRow(sql, entry.UserID, entry.JobID, now, entry.ID)
	var updatedEntry domain.Entry
	err = row.Scan(&updatedEntry.ID, &updatedEntry.UserID, &updatedEntry.JobID, &updatedEntry.CreatedAt, &updatedEntry.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &updatedEntry, nil
}

func (r *EntryRepositoryImpl) Delete(ctx context.Context, ids []int) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	sql := `DELETE FROM entries WHERE id in ($1)`
	_, err = tx.Exec(sql, ids)
	if err != nil {
		return err
	}

	return nil
}
