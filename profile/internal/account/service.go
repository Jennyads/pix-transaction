package account

import (
	"errors"
	"time"
)

type Service interface {
	CreateAccount(account *Account) (*Account, error)
	FindAccountById(req *AccountRequest) (*Account, error)
	UpdateAccount(account *Account) (*Account, error)
	ListAccount(req *ListAccountRequest) ([]*Account, error)
	DeleteAccount(req *AccountRequest) error
}

type service struct {
	repo Repository
}

func (s service) CreateAccount(account *Account) (*Account, error) {
	if account.UserID == 0 {
		return nil, errors.New("user_id is required")
	}
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()
	return s.repo.CreateAccount(account)
}

func (s service) FindAccountById(request *AccountRequest) (*Account, error) {
	return s.repo.FindAccountById(request.AccountID)
}

func (s service) UpdateAccount(account *Account) (*Account, error) {
	if account.Id == 0 {
		return nil, errors.New("id is required")
	}
	account.UpdatedAt = time.Now()
	return s.repo.UpdateAccount(account)
}

func (s service) ListAccount(listAccount *ListAccountRequest) ([]*Account, error) {
	if len(listAccount.AccountIDs) == 0 {
		return nil, errors.New("account_ids is required")
	}
	return s.repo.ListAccount(listAccount.AccountIDs)
}

func (s service) DeleteAccount(request *AccountRequest) error {
	if request.AccountID == 0 {
		return errors.New("account_id is required")
	}
	return s.repo.DeleteAccount(request.AccountID)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
