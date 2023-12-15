package profile

import (
	"github.com/shopspring/decimal"
	"time"
	proto "transaction/proto/v1"
)

type Account struct {
	Id        string
	UserID    string
	Balance   decimal.Decimal
	Agency    string
	Bank      string
	BlockedAt *time.Time
}

func ProtoToAccount(account *proto.Account) *Account {
	return &Account{
		UserID:  account.UserId,
		Balance: decimal.NewFromFloat(account.Balance),
		Agency:  account.Agency,
		Bank:    account.Bank,
	}
}

type Status string

const (
	StatusCompleted Status = "COMPLETED"
	StatusFailed    Status = "FAILED"
)

type Webhook struct {
	AccountID  string
	ReceiverID string
	Amount     decimal.Decimal
	Status     Status
}
