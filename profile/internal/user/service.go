package user

import (
	"errors"
	"time"
)

type Service interface {
	CreateUser(user *User) (*User, error)
	FindUserById(req *UserRequest) (*User, error)
	UpdateUser(user *User) (*User, error)
	ListUsers(req *ListUserRequest) ([]*User, error)
	DeleteUser(req *UserRequest) error
}

type service struct {
	repo Repository
}

func (s service) CreateUser(user *User) (*User, error) {

	// TODO validar cpf duplicado
	return nil, &errors.ErrDuplicated{Msg: "cpf already exists"}

	// TODO validar email duplicado
	return nil, &ErrDuplicated{Msg: "email already exists"}

	// TODO validar phone duplicado
	return nil, &ErrDuplicated{Msg: "phone already exists"}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return s.repo.CreateUser(user)
}

func (s service) FindUserById(userRequest *UserRequest) (*User, error) {
	return s.repo.FindUserById(int(userRequest.UserID))
}

func (s service) UpdateUser(user *User) (*User, error) {
	user.UpdatedAt = time.Now()
	return s.repo.UpdateUser(user)
}

func (s service) ListUsers(listUsers *ListUserRequest) ([]*User, error) {
	return s.repo.ListUsers(listUsers.UserIDs)
}

func (s service) DeleteUser(request *UserRequest) error {
	return s.repo.DeleteUser(int(request.UserID))
}
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
