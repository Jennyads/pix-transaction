package account

import "profile/platform/dynamo"

type Repository interface {
	CreateAccount(account *Account) (*Account, error)
	FindAccountById(id int) (*Account, error)
	UpdateAccount(account *Account) (*Account, error)
	ListAccount(accountIDs []int) ([]*Account, error)
	DeleteAccount(id int) error
}

type repository struct {
	db dynamo.Client
}

func (r repository) CreateAccount(account *Account) (*Account, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) FindAccountById(id int) (*Account, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) UpdateAccount(account *Account) (*Account, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) ListAccount(ints []int) ([]*Account, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) DeleteAccount(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db dynamo.Client) Repository {
	return &repository{
		db: db,
	}
}
