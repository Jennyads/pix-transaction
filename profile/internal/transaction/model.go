package transaction

import (
	"fmt"
	"github.com/shopspring/decimal"
	pb "profile/proto/profile/v1"
)

type Pix struct {
	UserID     string          `json:"user_id"`
	AccountID  string          `json:"account_id"`
	Key        string          `json:"key"`
	Receiver   string          `json:"receiver"`
	Amount     decimal.Decimal `json:"amount"`
	Status     string          `json:"status"`
	WebhookUrl string          `json:"webhook_url"`
}

func ProtoToPix(pix *pb.PixTransaction) *Pix {
	amountStr := fmt.Sprintf("%.2f", pix.Amount)
	amount, err := decimal.NewFromString(amountStr)
	if err != nil {
		return nil
	}

	return &Pix{
		AccountID: pix.Id,
		UserID:    pix.UserId,
		Key:       pix.SenderKey,
		Receiver:  pix.ReceiverKey,
		Amount:    amount,
		Status:    pix.Status,
	}
}

type PixEvent struct {
	Account struct {
		Name   string `json:"name"`
		Cpf    string `json:"cpf"`
		Agency string `json:"agency"`
		Bank   string `json:"bank"`
	} `json:"account"`
	Receiver   string          `json:"receiver"`
	Amount     decimal.Decimal `json:"amount"`
	WebhookUrl string          `json:"webhook_url"`
}

type Status string

const (
	StatusCompleted Status = "COMPLETED"
	StatusFailed    Status = "FAILED"
)

func ProtoToWebhook(webhook *pb.Webhook) *Webhook {
	amount := decimal.NewFromFloat(webhook.Amount)
	return &Webhook{
		Sender: Account{
			Name:   webhook.Sender.Name,
			Agency: webhook.Sender.Agency,
			Bank:   webhook.Sender.Bank,
		},
		Receiver: Account{
			Name:   webhook.Receiver.Name,
			Agency: webhook.Receiver.Agency,
			Bank:   webhook.Receiver.Bank,
		},
		Amount: amount,
		Status: Status(webhook.Status.String()),
	}
}

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
}

type Type string

const (
	Cpf    Type = "cpf"
	Phone  Type = "phone"
	Email  Type = "email"
	Random Type = "random"
)

type Key struct {
	Agency  string
	Bank    string
	Cpf     string
	Account string
	Name    string
	Type    Type
}
