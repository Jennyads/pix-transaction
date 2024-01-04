package key

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
	CreateKey(ctx context.Context, key *Key) (*Key, error)
	UpdateKey(ctx context.Context, key *Key) (*Key, error)
	ListKey(ctx context.Context, ids []string) ([]*Key, error)
	DeleteKey(ctx context.Context, id string) error
	FindKey(ctx context.Context, key string) (*Key, error)
}

type repository struct {
	db  dynamo.Client
	cfg *cfg.Config
}

func (r repository) CreateKey(ctx context.Context, key *Key) (*Key, error) {
	value, err := attributevalue.MarshalMap(key)
	if err != nil {
		return nil, err
	}

	var item *dynamodb.PutItemOutput
	item, err = r.db.DB().PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           aws.String(r.cfg.DynamodbConfig.TransactionTable),
		Item:                value,
		ConditionExpression: aws.String("attribute_not_exists(PK)"),
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

func (r repository) UpdateKey(ctx context.Context, key *Key) (*Key, error) {
	upd := expression.Set(expression.Name("Name"), expression.Value(key.Name)).
		Set(expression.Name("Type"), expression.Value(key.Type))

	expr, err := expression.NewBuilder().WithUpdate(upd).Build()
	if err != nil {
		return nil, errors.New("failed to build expression")
	}

	dyInput := &dynamodb.UpdateItemInput{
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: key.Id},
		},
		TableName:                 aws.String(r.cfg.DynamodbConfig.KeysTable),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}

	_, err = r.db.DB().UpdateItem(ctx, dyInput)
	if err != nil {
		return nil, errors.New("failed to update item")
	}

	return key, nil
}

func (r repository) ListKey(ctx context.Context, ids []string) ([]*Key, error) {
	keys := make([]map[string]types.AttributeValue, len(ids))
	for i, v := range ids {
		keys[i] = map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: v},
		}
	}

	value, err := r.db.DB().BatchGetItem(ctx, &dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			r.cfg.DynamodbConfig.KeysTable: {
				Keys: keys,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	listKey := make([]*Key, len(value.Responses[r.cfg.DynamodbConfig.KeysTable]))
	for i := range value.Responses[r.cfg.DynamodbConfig.KeysTable] {
		err = attributevalue.UnmarshalMap(value.Responses[r.cfg.DynamodbConfig.KeysTable][i], &listKey[i])
		if err != nil {
			return nil, err
		}
	}
	return listKey, nil
}

func (r repository) DeleteKey(ctx context.Context, id string) error {
	_, err := r.db.DB().DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(r.cfg.DynamodbConfig.KeysTable),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: id},
		},
	})
	return err
}

func (r repository) FindKey(ctx context.Context, key string) (*Key, error) {
	value, err := r.db.DB().Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(r.cfg.DynamodbConfig.KeysTable),
		KeyConditionExpression: aws.String("#name = :name"),
		ExpressionAttributeNames: map[string]string{
			"#name": "Name",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":name": &types.AttributeValueMemberS{Value: key},
		},
	})
	if err != nil {
		return nil, err
	}

	if len(value.Items) == 0 {
		return nil, nil
	}

	var keyModel Key
	err = attributevalue.UnmarshalMap(value.Items[0], &keyModel)
	if err != nil {
		return nil, err
	}
	return &keyModel, nil
}

func NewRepository(db dynamo.Client, config *cfg.Config) Repository {
	return &repository{
		db:  db,
		cfg: config,
	}
}
