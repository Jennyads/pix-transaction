package profile

import (
	"time"
	proto "transaction/proto/v1"
)

type Account struct {
	Id        string
	UserID    string
	Balance   float64
	Agency    string
	Bank      string
	BlockedAt *time.Time
}

func ProtoToAccount(account *proto.Account) *Account {
	return &Account{
		UserID:  account.UserId,
		Balance: account.Balance,
		Agency:  account.Agency,
		Bank:    account.Bank,
	}
}
