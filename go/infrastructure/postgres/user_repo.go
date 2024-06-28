package postgres

import (
	"context"
	"go-redis/domain"
	"time"
)

type UserRepositoryImpl struct {
}

func NewUserPostgresRepository() domain.UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) FindAll() ([]*domain.User, error) {
	sql := "SELECT * FROM users"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*domain.User
	for rows.Next() {
		u := &domain.User{}
		err = rows.Scan(&u.ID, &u.Name, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *UserRepositoryImpl) FindById(id int) (*domain.User, error) {
	sql := "SELECT * FROM users WHERE id = $1"
	row := db.QueryRow(sql, id)
	u := &domain.User{}
	if err := row.Scan(&u.ID, &u.Name, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
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

	sql := "INSERT INTO users (name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id,created_at,updated_at"
	now := time.Now()
	row := db.QueryRow(sql, user.Name, now, now)
	err = row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
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

	sql := "UPDATE users SET name = $1, updated_at = $2 WHERE id = $3 RETURNING id, name, created_at, updated_at"
	now := time.Now()

	row := db.QueryRow(sql, user.Name, now, user.ID)
	updatedUser := &domain.User{}
	err = row.Scan(&updatedUser.ID, &updatedUser.Name, &updatedUser.CreatedAt, &updatedUser.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, id []int) error {
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

	sql := "DELETE FROM users WHERE id in ($1)"
	_, err = db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}
