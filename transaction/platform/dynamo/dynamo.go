package dynamo

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"transaction/internal/cfg"
)

type Client interface {
	Connect(dynamodbConfig cfg.DynamodbConfig) Client
	DB() *dynamodb.Client
}

type client struct {
	db *dynamodb.Client
}

func (c *client) DB() *dynamodb.Client {
	return c.db
}

func (c *client) Connect(dynamodbConfig cfg.DynamodbConfig) Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-west-2"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: fmt.Sprintf("http://%s:%s", dynamodbConfig.Host, dynamodbConfig.Port)}, nil
			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "local", SecretAccessKey: "local", SessionToken: "",
				Source: "Mock credentials used above for local instance",
			},
		}))
	if err != nil {
		panic(err)
	}
	c.db = dynamodb.NewFromConfig(cfg)
	return c
}

func NewClient() Client {
	return &client{}
}
