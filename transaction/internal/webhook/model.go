package webhook

type Webhook struct {
	Sender   Account
	Receiver Account
	Amount   string
	Status   string
}

type Account struct {
	Name   string
	Agency string
	Bank   string
}
