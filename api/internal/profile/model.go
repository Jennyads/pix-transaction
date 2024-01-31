package profile

type WebhookStatus string

const (
	Pending   WebhookStatus = "FAILED"
	Confirmed WebhookStatus = "COMPLETED"
)

type Webhook struct {
	Sender   Account
	Receiver Account
	Amount   float64
	Status   WebhookStatus
}

type Account struct {
	Name   string
	Agency string
	Bank   string
}

type User struct {
	Id      string
	Name    string
	Email   string `validate:"isValidEmail"`
	Address string
	Cpf     string `validate:"isValidCPF"`
	Phone   string `validate:"isValidPhoneNumber"`
}
