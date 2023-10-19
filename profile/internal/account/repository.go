package account

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"profile/internal/cfg"
	"profile/platform/dynamo"
	"strconv"
)

type Repository interface {
	CreateAccount(account *Account) (*Account, error)
	FindAccountById(id int) (*Account, error)
	UpdateAccount(account *Account) (*Account, error)
	ListAccount(accountIDs []int64) ([]*Account, error)
	DeleteAccount(id int) error
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
		TableName: aws.String(r.cfg.DynamodbConfig.AccountTable),
		Item:      value,
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

func (r repository) FindAccountById(id int) (*Account, error) {
	value, err := r.db.DB().GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(r.cfg.DynamodbConfig.AccountTable),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberN{Value: strconv.Itoa(id)},
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

	//upd := expression.
	//	Set(expression.Name("Status"), expression.Value(elem.Status)).
	//	Set(expression.Name("UpdateAt"), expression.Value(elem.UpdateAt)).
	//	Set(expression.Name("Attempts"), expression.Value(elem.Attempts)).
	//	Set(expression.Name("TransactionId"), expression.Value(elem.TransactionId))
	//
	//exp, err := expression.NewBuilder().WithUpdate(upd).Build()
	//if err != nil {
	//	return nil, errors.New("failed to build expression")
	//}
	//item, err := r.db.DB().UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
	//	TableName: aws.String(r.cfg.DynamodbConfig.AccountTable),
	//	Key: map[string]types.AttributeValue{
	//		"PK": &types.AttributeValueMemberN{Value: strconv.FormatInt(account.Id, 10)},
	//	},
	//	ExpressionAttributeNames:  exp.Names(),
	//	ExpressionAttributeValues: exp.Values(),
	//	UpdateExpression:          exp.Update(),
	//	ReturnValues:              types.ReturnValueAllNew,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = attributevalue.UnmarshalMap(item.Attributes, account)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return account, nil
	return nil, nil
}

func (r repository) ListAccount(ids []int64) ([]*Account, error) {
	keys := make([]map[string]types.AttributeValue, len(ids))
	for i, v := range ids {
		keys[i] = map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: strconv.FormatInt(v, 10)},
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

func (r repository) DeleteAccount(id int) error {
	_, err := r.db.DB().DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String(r.cfg.DynamodbConfig.AccountTable),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberN{Value: strconv.Itoa(id)},
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
