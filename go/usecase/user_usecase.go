package usecase

import (
	"go-redis/controller/in"
	"go-redis/controller/out"
	"go-redis/domain"
	"go-redis/infrastructure/postgres"
)

type UserUsecase interface {
	List() ([]*out.UserResponse, error)
	Get(id int) (*out.UserResponse, error)
	Create(user *in.CreateUserRequest) (*out.UserResponse, error)
	Update(user *in.UpdateUserRequest) (*out.UserResponse, error)
	Delete(ids *in.DeleteUserRequest) error
}

type UserUsecaseImpl struct {
	userRepo domain.UserRepository
}

func NewUserUsecase() UserUsecase {
	return &UserUsecaseImpl{
		userRepo: postgres.NewUserPostgresRepository(),
	}
}

func (u *UserUsecaseImpl) List() ([]*out.UserResponse, error) {
	users, err := u.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var usersResponse = make([]*out.UserResponse, 0, len(users))
	for _, user := range users {
		usersResponse = append(usersResponse, &out.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdateAt:  user.UpdatedAt,
		})
	}
	return usersResponse, nil
}

func (u *UserUsecaseImpl) Get(id int) (*out.UserResponse, error) {
	user, err := u.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	userResponse := &out.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdatedAt,
	}

	return userResponse, nil
}

func (u *UserUsecaseImpl) Create(user *in.CreateUserRequest) (*out.UserResponse, error) {
	userDomain := domain.NewUser(0, user.Name)
	userCreated, err := u.userRepo.Create(userDomain)
	if err != nil {
		return nil, err
	}

	userResponse := &out.UserResponse{
		ID:        userCreated.ID,
		Name:      userCreated.Name,
		CreatedAt: userCreated.CreatedAt,
		UpdateAt:  userCreated.UpdatedAt,
	}
	return userResponse, nil
}

func (u *UserUsecaseImpl) Update(user *in.UpdateUserRequest) (*out.UserResponse, error) {
	userDomain := domain.NewUser(user.ID, user.Name)
	err := u.userRepo.Update(userDomain)
	if err != nil {
		return nil, err
	}
	getUser, err := u.userRepo.FindById(user.ID)
	if err != nil {
		return nil, err
	}
	var userResponse = &out.UserResponse{
		ID:        getUser.ID,
		Name:      getUser.Name,
		CreatedAt: getUser.CreatedAt,
		UpdateAt:  getUser.UpdatedAt,
	}
	return userResponse, nil
}

func (u *UserUsecaseImpl) Delete(ids *in.DeleteUserRequest) error {
	return u.userRepo.Delete(ids.ID)
}
