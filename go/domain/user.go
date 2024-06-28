package domain

import "time"

type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	FindAll() ([]*User, error)
	FindById(id int) (*User, error)
	Create(user *User) (*User, error)
	Update(user *User) error
	Delete(id []int) error
}

func NewUser(id int, name string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}
