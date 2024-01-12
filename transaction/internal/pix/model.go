package pix

import "github.com/shopspring/decimal"

type Pix struct {
	AccountName   string
	AccountCpf    string
	AccountAgency string
	AccountBank   string
	Receiver      string
	Amount        decimal.Decimal
	Status        string
	WebhookUrl    string
}

type PixEvent struct {
	PixData *Pix
}
