package transactions

import (
	"transaction/platform/dynamo"
)

type Repository interface {
	CreateTransaction(transaction *Transaction) error
	FindTransaction(id *TransactionRequest) (*Transaction, error)
	ListTransactions(transactionIDs *ListTransactionRequest) (*ListTransaction, error)
}

type repository struct {
	db dynamo.Client
}

func (r repository) CreateTransaction(transaction *Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) FindTransaction(id *TransactionRequest) (*Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) ListTransactions(transactionIDs *ListTransactionRequest) (*ListTransaction, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db dynamo.Client) Repository {
	return &repository{
		db: db,
	}
}
