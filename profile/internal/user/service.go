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
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return s.repo.CreateUser(user)
}

func (s service) FindUserById(userRequest *UserRequest) (*User, error) {
	return s.repo.FindUserById(int(userRequest.UserID))
}

func (s service) UpdateUser(user *User) (*User, error) {
	if user.PK == 0 {
		return nil, errors.New("id is required")
	}
	user.UpdatedAt = time.Now()
	return s.repo.UpdateUser(user)
}

func (s service) ListUsers(listUsers *ListUserRequest) ([]*User, error) {
	if len(listUsers.UserIDs) == 0 {
		return nil, errors.New("user_ids is required")
	}
	return s.repo.ListUsers(listUsers.UserIDs)
}

func (s service) DeleteUser(request *UserRequest) error {
	if request.UserID == 0 {
		return errors.New("account_id is required")
	}
	return s.repo.DeleteUser(int(request.UserID))
}
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
