package pix

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"profile/internal/cfg"
	"profile/platform/dynamo"
)

type Repository interface {
	CreatePixTransaction(transaction *PixTransaction) (*PixTransaction, error)
	ListPixTransactions(request *ListPixTransactionsRequest) ([]*PixTransaction, error)
}

type repository struct {
	db  dynamo.Client
	cfg *cfg.Config
}

func (r repository) CreatePixTransaction(pixTransaction *PixTransaction) (*PixTransaction, error) {
	value, err := attributevalue.MarshalMap(pixTransaction)
	if err != nil {
		return nil, err
	}
	item, err := r.db.DB().PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName:           aws.String(r.cfg.DynamodbConfig.PixTable),
		Item:                value,
		ConditionExpression: aws.String("attribute_not_exists(PK) and attribute_not_exists(SK)"),
	})
	if err != nil {
		return nil, err
	}
	err = attributevalue.UnmarshalMap(item.Attributes, pixTransaction)
	if err != nil {
		return nil, err
	}
	return pixTransaction, nil

}
func (r repository) ListPixTransactions(request *ListPixTransactionsRequest) ([]*PixTransaction, error) {
	pixTransactions := make([]map[string]types.AttributeValue, len(request.PixTransactionIDs))
	for i, v := range request.PixTransactionIDs {
		pixTransactions[i] = map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: v},
		}
	}
	value, err := r.db.DB().BatchGetItem(context.Background(), &dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			r.cfg.DynamodbConfig.PixTable: {
				Keys: pixTransactions,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	listPixTransaction := make([]*PixTransaction, len(value.Responses[r.cfg.DynamodbConfig.PixTable]))
	for i := range value.Responses[r.cfg.DynamodbConfig.PixTable] {
		err = attributevalue.UnmarshalMap(value.Responses[r.cfg.DynamodbConfig.PixTable][i], &listPixTransaction[i])
		if err != nil {
			return nil, err
		}
	}
	return listPixTransaction, nil
}
func NewRepository(db dynamo.Client, config *cfg.Config) Repository {
	return &repository{
		db:  db,
		cfg: config,
	}
}
