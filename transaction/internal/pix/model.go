package pix

import "github.com/shopspring/decimal"

type PixEvent struct {
	Account struct {
		Name   int64  `json:"name"`
		Cpf    string `json:"cpf"`
		Agency string `json:"agency"`
		Bank   string `json:"bank"`
	} `json:"account"`
	Receiver   string          `json:"receiver"`
	Amount     decimal.Decimal `json:"amount"`
	WebhookUrl string          `json:"webhook_url"`
}
