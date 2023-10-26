package keys

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
	CreateKey(key *Key) (*Key, error)
	UpdateKey(key *Key) (*Key, error)
	ListKey(keyIDs []string) ([]*Key, error)
	DeleteKey(id string) error
}

type repository struct {
	db  dynamo.Client
	cfg *cfg.Config
}

func (r repository) CreateKey(key *Key) (*Key, error) {
	value, err := attributevalue.MarshalMap(key)
	if err != nil {
		return nil, err
	}

	item, err := r.db.DB().PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName:           aws.String(r.cfg.DynamodbConfig.KeyTable),
		Item:                value,
		ConditionExpression: aws.String("attribute_not_exists(PK) and attribute_not_exists(SK)"),
	})
	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalMap(item.Attributes, key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func (r repository) UpdateKey(key *Key) (*Key, error) {
	upd := expression.
		Set(expression.Name("Name"), expression.Value(key.Name)).
		Set(expression.Name("Type"), expression.Value(key.Type))

	exp, err := expression.NewBuilder().WithUpdate(upd).Build()
	if err != nil {
		return nil, errors.New("failed to build expression")
	}
	item, err := r.db.DB().UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName: aws.String(r.cfg.DynamodbConfig.KeyTable),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberN{Value: key.Id},
		},
		ExpressionAttributeNames:  exp.Names(),
		ExpressionAttributeValues: exp.Values(),
		UpdateExpression:          exp.Update(),
		ReturnValues:              types.ReturnValueAllNew,
	})
	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalMap(item.Attributes, key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func (r repository) ListKey(keyIDs []string) ([]*Key, error) {
	keys := make([]map[string]types.AttributeValue, len(keyIDs))
	for i, v := range keyIDs {
		keys[i] = map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: v},
		}
	}

	value, err := r.db.DB().BatchGetItem(context.Background(), &dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			r.cfg.DynamodbConfig.KeyTable: {
				Keys: keys,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	listKey := make([]*Key, len(value.Responses[r.cfg.DynamodbConfig.KeyTable]))
	for i := range value.Responses[r.cfg.DynamodbConfig.KeyTable] {
		err = attributevalue.UnmarshalMap(value.Responses[r.cfg.DynamodbConfig.KeyTable][i], &listKey[i])
		if err != nil {
			return nil, err
		}
	}
	return listKey, nil
}

func (r repository) DeleteKey(id string) error {
	_, err := r.db.DB().DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String(r.cfg.DynamodbConfig.KeyTable),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberN{Value: id},
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
