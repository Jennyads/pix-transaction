package transactions

import (
	"time"
	pb "transaction/proto"
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

func ProtoToTransactionRequest(request *pb.TransactionRequest) *TransactionRequest {
	return &TransactionRequest{
		TransactionID: request.TransactionId,
		AccountID:     request.AccountId,
	}
}

func ProtoToListTransactionRequest(request *pb.ListTransactionRequest) *ListTransactionRequest {
	return &ListTransactionRequest{
		transactionIDs: request.TransactionId,
	}
}

func ProtoToTransaction(transaction *pb.Transaction) *Transaction {
	return &Transaction{
		AccountID: transaction.AccountId,
		Receiver:  transaction.Receiver,
		Value:     transaction.Value,
		Status:    transaction.Status,
	}
}

func ToProto(transaction *Transaction) *pb.Transaction {
	return &pb.Transaction{
		AccountId: transaction.AccountID,
		Receiver:  transaction.Receiver,
		Value:     transaction.Value,
		Status:    transaction.Status,
	}
}
