package profile

import "github.com/shopspring/decimal"

type WebhookStatus string

const (
	Pending   WebhookStatus = "FAILED"
	Confirmed WebhookStatus = "COMPLETED"
)

type Webhook struct {
	Sender   Account         `json:"sender"`
	Receiver Account         `json:"receiver"`
	Amount   decimal.Decimal `json:"amount"`
	Status   WebhookStatus   `json:"status"`
}

type PixTransaction struct {
	UserId      string  `json:"user_id"`
	AccountId   string  `json:"account_id"`
	SenderKey   string  `json:"sender_key"`
	ReceiverKey string  `json:"receiver_key"`
	Amount      float64 `json:"amount"`
}

type Account struct {
	Name   string `json:"name"`
	Cpf    string `json:"cpf"`
	Agency string `json:"agency"`
	Bank   string `json:"bank"`
}

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email" validate:"validateData"`
	Address string `json:"address"`
	Cpf     string `json:"cpf" validate:"validateData"`
	Phone   string `json:"phone" validate:"validateData"`
}

type Type string

const (
	TypeCpf    Type = "cpf"
	TypePhone  Type = "phone"
	TypeEmail  Type = "email"
	TypeRandom Type = "random"
)

type Key struct {
	AccountId string `json:"account_id"`
	Name      string `json:"name"`
	Type      Type   `json:"type"`
}
