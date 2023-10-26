package transactions

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"transaction/internal/cfg"
	"transaction/platform/dynamo"
)

type Repository interface {
	CreateTransaction(transaction *Transaction) error
	FindTransaction(id string) (*Transaction, error)
	ListTransactions(ids []string) ([]*Transaction, error)
}

type repository struct {
	db  dynamo.Client
	cfg *cfg.Config
}

func (r repository) CreateTransaction(transaction *Transaction) error {
	value, err := attributevalue.MarshalMap(transaction)
	if err != nil {
		return err
	}

	item, err := r.db.DB().PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(r.cfg.DynamodbConfig.TransactionTable),
		Item:      value,
	})
	if err != nil {
		return err
	}

	err = attributevalue.UnmarshalMap(item.Attributes, transaction)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) FindTransaction(id string) (*Transaction, error) {
	value, err := r.db.DB().GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(r.cfg.DynamodbConfig.TransactionTable),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberN{Value: id},
		},
	})
	if err != nil {
		return nil, err
	}

	var transaction Transaction
	err = attributevalue.UnmarshalMap(value.Item, &transaction)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r repository) ListTransactions(ids []string) ([]*Transaction, error) {
	keys := make([]map[string]types.AttributeValue, len(ids))
	for i, v := range ids {
		keys[i] = map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: v},
		}
	}

	value, err := r.db.DB().BatchGetItem(context.Background(), &dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			r.cfg.DynamodbConfig.TransactionTable: {
				Keys: keys,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	listTransaction := make([]*Transaction, len(value.Responses[r.cfg.DynamodbConfig.TransactionTable]))
	for i := range value.Responses[r.cfg.DynamodbConfig.TransactionTable] {
		err = attributevalue.UnmarshalMap(value.Responses[r.cfg.DynamodbConfig.TransactionTable][i], &listTransaction[i])
		if err != nil {
			return nil, err
		}
	}
	return listTransaction, nil
}

func NewRepository(db dynamo.Client) Repository {
	return &repository{
		db: db,
	}
}
