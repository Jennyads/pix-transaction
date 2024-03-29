package transactions

import (
	"time"
	proto "transaction/proto/v1"
)

type Status string

const (
	StatusPending   Status = "PENDING"
	StatusCompleted Status = "COMPLETED"
	StatusFailed    Status = "FAILED"
)

type Transaction struct {
	ID          string `dynamodbav:"PK"`
	AccountID   int64
	Receiver    string
	Value       float64
	Status      Status
	ErrMessage  string
	CreatedAt   time.Time
	ProcessedAt time.Time
	UpdatedAt   time.Time
}

type TransactionRequest struct {
	TransactionID string
	AccountID     int64
}

type ListTransactionRequest struct {
	transactionIDs []string
}

type ListTransaction struct {
	Transactions []*Transaction
}

func ProtoToTransactionRequest(request *proto.TransactionRequest) *TransactionRequest {
	return &TransactionRequest{
		TransactionID: request.TransactionId,
		AccountID:     request.AccountId,
	}
}

func ProtoToListTransactionRequest(request *proto.ListTransactionRequest) *ListTransactionRequest {
	return &ListTransactionRequest{
		transactionIDs: request.TransactionId,
	}
}

func ProtoToTransaction(transaction *proto.Transaction) *Transaction {
	return &Transaction{
		AccountID: transaction.AccountId,
		Receiver:  transaction.Receiver,
		Value:     transaction.Value,
		Status:    Status(transaction.Status),
	}
}

func ToProto(transaction *Transaction) *proto.Transaction {
	return &proto.Transaction{
		AccountId: transaction.AccountID,
		Receiver:  transaction.Receiver,
		Value:     transaction.Value,
		Status:    string(transaction.Status),
	}
}
