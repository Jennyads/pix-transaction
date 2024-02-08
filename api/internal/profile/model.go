package profile

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type WebhookStatus string

const (
	Pending   WebhookStatus = "FAILED"
	Confirmed WebhookStatus = "COMPLETED"
)

type Webhook struct {
	Sender   Account       `json:"sender"`
	Receiver Account       `json:"receiver"`
	Amount   float64       `json:"amount"`
	Status   WebhookStatus `json:"status"`
}

type PixTransaction struct {
	UserId      string               `json:"userId"`
	SenderKey   string               `json:"senderKey"`
	ReceiverKey string               `json:"receiverKey"`
	Amount      float64              `json:"amount"`
	Hour        *timestamp.Timestamp `json:"hour"`
	Status      string               `json:"status"`
}

type Account struct {
	Name   string `json:"name"`
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
