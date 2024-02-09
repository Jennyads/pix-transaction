package pix

import "github.com/shopspring/decimal"

type PixEvent struct {
	AccountName   string          `json:"account_name"`
	AccountCpf    string          `json:"account_cpf"`
	AccountAgency string          `json:"account_agency"`
	AccountBank   string          `json:"account_bank"`
	Receiver      string          `json:"receiver"`
	Amount        decimal.Decimal `json:"amount"`
	Status        string          `json:"status"`
	WebhookUrl    string          `json:"webhook_url"`
}
