package account

import (
	pb "profile/proto/v1"
	"time"
)

type Account struct {
	Id        int64 `dynamodbav:"PK"`
	UserID    int64 `dynamodbav:"SK"`
	Balance   float64
	Agency    string
	Bank      string
	Key       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type AccountRequest struct {
	AccountID int64
}

type ListAccountRequest struct {
	AccountIDs []int64
}

func ProtoToAccount(account *pb.Account) *Account {
	return &Account{
		UserID:  account.UserId,
		Balance: account.Balance,
		Agency:  account.Agency,
		Bank:    account.Bank,
		Key:     account.Key,
	}
}

func ProtoToAccountRequest(request *pb.AccountRequest) *AccountRequest {
	return &AccountRequest{
		AccountID: request.AccountId,
	}
}

func ProtoToAccountListRequest(request *pb.ListAccountRequest) *ListAccountRequest {
	return &ListAccountRequest{
		AccountIDs: request.AccountId,
	}
}
