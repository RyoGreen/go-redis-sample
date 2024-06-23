package domain

type User struct {
	ID        int
	Name      string
	CreatedAt string
	UpdatedAt string
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindById(id int) (User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
	Delete(id int) error
}
