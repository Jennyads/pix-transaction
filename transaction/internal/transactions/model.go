package transactions

import (
	"time"
	pb "transaction/proto"
)

type Transaction struct {
	PK          int
	AccountID   int `dynamodbav:"SK"`
	Receiver    int
	Value       float64
	Status      string
	CreatedAt   time.Time
	ProcessedAt time.Time
}

type TransactionRequest struct {
	transactionID int
}

type ListTransactionRequest struct {
	transactionIDs []int64
}

type ListTransaction struct {
	Transactions []*Transaction
}

func ProtoToTransactionRequest(request *pb.TransactionRequest) *TransactionRequest {
	return &TransactionRequest{
		transactionID: int(request.TransactionId),
	}
}

func ProtoToListTransactionRequest(request *pb.ListTransactionRequest) *ListTransactionRequest {
	return &ListTransactionRequest{
		transactionIDs: request.TransactionId,
	}
}

func ProtoToTransaction(transaction *pb.Transaction) *Transaction {
	return &Transaction{
		AccountID: int(transaction.AccountId),
		Receiver:  int(transaction.Receiver),
		Value:     transaction.Value,
		Status:    transaction.Status,
	}
}
