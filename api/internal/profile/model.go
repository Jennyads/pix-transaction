package profile

type WebhookStatus string

const (
	Pending   WebhookStatus = "pending"
	Confirmed WebhookStatus = "confirmed"
)

type Webhook struct {
	AccountId  string        `json:"account_id"`
	ReceiverId string        `json:"receiver_id"`
	Amount     float64       `json:"amount"`
	Status     WebhookStatus `json:"status"`
}
