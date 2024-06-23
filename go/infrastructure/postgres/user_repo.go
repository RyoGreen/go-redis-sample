package postgres

import "go-redis/domain"

type UserRepositoryImpl struct {
}

func NewUserPostgresRepository() domain.UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) FindAll() ([]domain.User, error) {
	return nil, nil
}

func (r *UserRepositoryImpl) FindById(id int) (domain.User, error) {
	return domain.User{}, nil
}

func (r *UserRepositoryImpl) Create(user domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (r *UserRepositoryImpl) Update(user domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (r *UserRepositoryImpl) Delete(id int) error {
	return nil
}
