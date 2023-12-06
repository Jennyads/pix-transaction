package account

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
)

type Service interface {
	CreateAccount(ctx context.Context, account *Account) (*Account, error)
	FindAccountById(ctx context.Context, req *AccountRequest) (*Account, error)
	UpdateAccount(ctx context.Context, account *Account) (*Account, error)
	ListAccounts(ctx context.Context, req *ListAccountRequest) ([]*Account, error)
	DeleteAccount(ctx context.Context, req *AccountRequest) error
}

type service struct {
	repo Repository
}

func (s service) CreateAccount(ctx context.Context, account *Account) (*Account, error) {
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()
	account.Id = uuid.New().String()
	return s.repo.CreateAccount(account)
}

func (s service) FindAccountById(ctx context.Context, request *AccountRequest) (*Account, error) {
	return s.repo.FindAccountById(request.AccountID)
}

func (s service) UpdateAccount(ctx context.Context, account *Account) (*Account, error) {
	if account.Id == "" {
		return nil, errors.New("account_id is required")
	}
	account.UpdatedAt = time.Now()
	return s.repo.UpdateAccount(account)
}

func (s service) ListAccounts(ctx context.Context, listAccount *ListAccountRequest) ([]*Account, error) {
	if len(listAccount.AccountIDs) == 0 {
		return nil, errors.New("account_ids is required")
	}
	return s.repo.ListAccount(listAccount.AccountIDs)
}

func (s service) DeleteAccount(ctx context.Context, request *AccountRequest) error {
	return s.repo.DeleteAccount(request.AccountID)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
