package transactions

import (
	"time"
	proto "transaction/proto/v1"
)

type Transaction struct {
	ID          string `dynamodbav:"PK"`
	AccountID   string `dynamodbav:"SK"`
	Receiver    int64
	Value       float64
	Status      string
	CreatedAt   time.Time
	ProcessedAt time.Time
}

type TransactionRequest struct {
	TransactionID string
	AccountID     string
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
		Status:    transaction.Status,
	}
}

func ToProto(transaction *Transaction) *proto.Transaction {
	return &proto.Transaction{
		AccountId: transaction.AccountID,
		Receiver:  transaction.Receiver,
		Value:     transaction.Value,
		Status:    transaction.Status,
	}
}
