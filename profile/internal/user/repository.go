package user

import "profile/platform/dynamo"

type Repository interface {
	CreateUser(account *User) (*User, error)
	FindUserById(id int) (*User, error)
	UpdateUser(account *User) (*User, error)
	ListUser(userIDs []int) ([]*User, error)
	DeleteUser(id int) error
}

type repository struct {
	db dynamo.Client
}

func (r repository) CreateUser(account *User) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) FindUserById(id int) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) UpdateUser(account *User) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) ListUser(ints []int) ([]*User, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) DeleteUser(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db dynamo.Client) Repository {
	return &repository{
		db: db,
	}
}
