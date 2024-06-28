package postgres

import (
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
		err := rows.Scan(&u.ID, &u.Name, &u.CreatedAt, &u.UpdatedAt)
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
	err := row.Scan(&u.ID, &u.Name, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepositoryImpl) Create(user *domain.User) (*domain.User, error) {
	sql := "INSERT INTO users (name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id,created_at,updated_at"
	now := time.Now()
	row := db.QueryRow(sql, user.Name, now, now)
	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Update(user *domain.User) error {
	sql := "UPDATE users SET name = $1, updated_at = $2 WHERE id = $3 RETURNING id,name,created_at,updated_at"
	now := time.Now()
	_, err := db.Exec(sql, user.Name, now, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Delete(id []int) error {
	sql := "DELETE FROM users WHERE id in ($1)"
	_, err := db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}
