package domain

import (
	"context"
	"time"
)

type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	FindAll() ([]*User, error)
	FindById(id int) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id []int) error
}

func NewUser(id int, name string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}
