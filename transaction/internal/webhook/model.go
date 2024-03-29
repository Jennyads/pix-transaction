package webhook

import "github.com/shopspring/decimal"

type Status string

const (
	StatusCompleted Status = "COMPLETED"
	StatusFailed    Status = "FAILED"
)

type Webhook struct {
	TransactionId string          `json:"transaction_id"`
	Sender        Account         `json:"sender"`
	Receiver      Account         `json:"receiver"`
	Amount        decimal.Decimal `json:"amount"`
	Status        Status          `json:"status"`
}

type Account struct {
	Name   int64
	Agency string
	Bank   string
	Cpf    string
}
