package webhook

import "github.com/shopspring/decimal"

type Status string

const (
	StatusCompleted Status = "COMPLETED"
	StatusFailed    Status = "FAILED"
)

type Webhook struct {
	Sender   Account
	Receiver Account
	Amount   decimal.Decimal
	Status   Status
}

type Account struct {
	Name   string
	Agency string
	Bank   string
	Cpf    string
}
