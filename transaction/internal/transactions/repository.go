package transactions

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"transaction/internal/cfg"
	"transaction/platform/dynamo"
)

type Repository interface {
	CreateTransaction(transaction *Transaction) (*Transaction, error)
	FindTransactionById(id string) (*Transaction, error)
	ListTransactions(ids []string) ([]*Transaction, error)
	UpdateTransactionStatus(transaction *Transaction) error
}

type repository struct {
	db  dynamo.Client
	cfg *cfg.Config
}

func (r *repository) CreateTransaction(transaction *Transaction) (*Transaction, error) {
	value, err := attributevalue.MarshalMap(transaction)
	if err != nil {
		return nil, err
	}

	var item *dynamodb.PutItemOutput
	item, err = r.db.DB().PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName:           aws.String(r.cfg.DynamodbConfig.TransactionTable),
		Item:                value,
		ConditionExpression: aws.String("attribute_not_exists(PK)"),
	})
	if err != nil {
		return nil, err
	}
	err = attributevalue.UnmarshalMap(item.Attributes, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *repository) FindTransactionById(id string) (*Transaction, error) {
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

func (r *repository) ListTransactions(ids []string) ([]*Transaction, error) {
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

func (r *repository) UpdateTransactionStatus(transaction *Transaction) error {
	upd := expression.Set(expression.Name("Status"), expression.Value(transaction.Status)).
		Set(expression.Name("UpdateAt"), expression.Value(transaction.UpdatedAt)).
		Set(expression.Name("ProcessedAt"), expression.Value(transaction.ProcessedAt)).
		Set(expression.Name("ErrMessage"), expression.Value(transaction.ErrMessage))

	expr, err := expression.NewBuilder().WithUpdate(upd).Build()
	if err != nil {
		return errors.New("failed to build expression")
	}

	dyInput := &dynamodb.UpdateItemInput{
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: transaction.ID},
		},
		TableName:                 aws.String(r.cfg.DynamodbConfig.TransactionTable),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}

	_, err = r.db.DB().UpdateItem(context.Background(), dyInput)
	if err != nil {
		return errors.New("failed to update item")
	}

	return nil
}

func NewRepository(db dynamo.Client, config *cfg.Config) Repository {
	return &repository{
		db:  db,
		cfg: config,
	}
}
