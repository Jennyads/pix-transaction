package keys

import (
	pb "profile/proto/v1"
	"time"
)

type Type string

const (
	Cpf    Type = "cpf"
	Phone  Type = "phone"
	Email  Type = "email"
	Random Type = "random"
)

type Key struct {
	Id        int64 `dynamodbav:"PK"`
	AccountID int64 `dynamodbav:"SK"`
	Name      string
	Type      Type
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type KeyRequest struct {
	keyID int64
}

type ListKeyRequest struct {
	keyIDs []int64
}

func ProtoToKey(key *pb.Key) *Key {
	return &Key{
		AccountID: key.AccountId,
		Name:      key.Name,
		Type:      Type(key.Type),
	}
}

func ProtoToKeyListRequest(request *pb.ListKeyRequest) *ListKeyRequest {
	return &ListKeyRequest{
		keyIDs: request.KeyId,
	}
}
func ProtoToKeyRequest(request *pb.KeyRequest) *KeyRequest {
	return &KeyRequest{
		keyID: request.KeyId,
	}
}
