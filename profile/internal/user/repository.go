package user

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
	CreateUser(user *User) (*User, error)
	FindUserById(id string) (*User, error)
	UpdateUser(user *User) (*User, error)
	ListUsers(userIDs []string) ([]*User, error)
	DeleteUser(id string) error
}

type repository struct {
	db  dynamo.Client
	cfg *cfg.Config
}

func (r repository) CreateUser(user *User) (*User, error) {
	value, err := attributevalue.MarshalMap(user)
	if err != nil {
		return nil, err
	}

	_, err = r.db.DB().PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName:           aws.String(r.cfg.DynamodbConfig.UserTable),
		Item:                value,
		ConditionExpression: aws.String("attribute_not_exists(PK)"),
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r repository) ExistWithCpf(cpf string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) FindUserById(id string) (*User, error) {
	value, err := r.db.DB().GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(r.cfg.DynamodbConfig.UserTable),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, err
	}

	var user User
	err = attributevalue.UnmarshalMap(value.Item, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r repository) UpdateUser(user *User) (*User, error) {
	upd := expression.
		Set(expression.Name("Name"), expression.Value(user.Name)).
		Set(expression.Name("Email"), expression.Value(user.Email)).
		Set(expression.Name("Adress"), expression.Value(user.Address)).
		Set(expression.Name("Phone"), expression.Value(user.Phone)).
		Set(expression.Name("Birthday"), expression.Value(user.Birthday))

	exp, err := expression.NewBuilder().WithUpdate(upd).Build()
	if err != nil {
		return nil, errors.New("failed to build expression")
	}
	item, err := r.db.DB().UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName: aws.String(r.cfg.DynamodbConfig.UserTable),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: user.Id},
		},
		ExpressionAttributeNames:  exp.Names(),
		ExpressionAttributeValues: exp.Values(),
		UpdateExpression:          exp.Update(),
		ReturnValues:              types.ReturnValueAllNew,
	})
	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalMap(item.Attributes, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r repository) ListUsers(userIds []string) ([]*User, error) {
	keys := make([]map[string]types.AttributeValue, len(userIds))
	for i, v := range userIds {
		keys[i] = map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: v},
		}
	}

	value, err := r.db.DB().BatchGetItem(context.Background(), &dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			r.cfg.DynamodbConfig.UserTable: {
				Keys: keys,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	listUser := make([]*User, len(value.Responses[r.cfg.DynamodbConfig.UserTable]))
	for i := range value.Responses[r.cfg.DynamodbConfig.UserTable] {
		err = attributevalue.UnmarshalMap(value.Responses[r.cfg.DynamodbConfig.UserTable][i], &listUser[i])
		if err != nil {
			return nil, err
		}
	}
	return listUser, nil
}

func (r repository) DeleteUser(id string) error {
	_, err := r.db.DB().DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String(r.cfg.DynamodbConfig.UserTable),
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
