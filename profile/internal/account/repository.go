package account

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"profile/internal/cfg"
	"profile/platform/dynamo"
)

type Repository interface {
	CreateAccount(account *Account) (*Account, error)
	FindAccountById(id string) (*Account, error)
	UpdateAccount(account *Account) (*Account, error)
	ListAccount(accountIDs []string) ([]*Account, error)
	DeleteAccount(id string) error
}

type repository struct {
	db  dynamo.Client
	cfg *cfg.Config
}

func (r repository) CreateAccount(account *Account) (*Account, error) {
	value, err := attributevalue.MarshalMap(account)
	if err != nil {
		return nil, err
	}

	item, err := r.db.DB().PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName:           aws.String(r.cfg.DynamodbConfig.AccountTable),
		Item:                value,
		ConditionExpression: aws.String("attribute_not_exists(PK) and attribute_not_exists(SK)"),
	})
	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalMap(item.Attributes, account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (r repository) FindAccountById(id string) (*Account, error) {
	value, err := r.db.DB().GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(r.cfg.DynamodbConfig.AccountTable),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, err
	}

	var account Account
	err = attributevalue.UnmarshalMap(value.Item, &account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r repository) UpdateAccount(account *Account) (*Account, error) {
	upd := expression.
		Set(expression.Name("Agency"), expression.Value(account.Agency)).
		Set(expression.Name("Type"), expression.Value(account.Bank))

	exp, err := expression.NewBuilder().WithUpdate(upd).Build()
	if err != nil {
		return nil, errors.New("failed to build expression")
	}
	item, err := r.db.DB().UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName: aws.String(r.cfg.DynamodbConfig.KeyTable),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: account.Id},
		},
		ExpressionAttributeNames:  exp.Names(),
		ExpressionAttributeValues: exp.Values(),
		UpdateExpression:          exp.Update(),
		ReturnValues:              types.ReturnValueAllNew,
	})
	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalMap(item.Attributes, account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (r repository) ListAccount(ids []string) ([]*Account, error) {
	keys := make([]map[string]types.AttributeValue, len(ids))
	for i, v := range ids {
		keys[i] = map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: v},
		}
	}

	value, err := r.db.DB().BatchGetItem(context.Background(), &dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			r.cfg.DynamodbConfig.AccountTable: {
				Keys: keys,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	listAccount := make([]*Account, len(value.Responses[r.cfg.DynamodbConfig.AccountTable]))
	for i := range value.Responses[r.cfg.DynamodbConfig.AccountTable] {
		err = attributevalue.UnmarshalMap(value.Responses[r.cfg.DynamodbConfig.AccountTable][i], &listAccount[i])
		if err != nil {
			return nil, err
		}
	}
	return listAccount, nil
}

func (r repository) DeleteAccount(id string) error {
	_, err := r.db.DB().DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String(r.cfg.DynamodbConfig.AccountTable),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: id},
		},
	})
	return err
}

func NewRepository(db dynamo.Client, config *cfg.Config) Repository {
	return &repository{
		db:  db,
		cfg: config,
	}
}
