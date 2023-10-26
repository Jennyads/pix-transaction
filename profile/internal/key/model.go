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
	Id        string `dynamodbav:"PK"`
	AccountID string `dynamodbav:"SK"`
	Name      string
	Type      Type
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type KeyRequest struct {
	keyID string
}

type ListKeyRequest struct {
	keyIDs []string
}

func ProtoToKey(key *pb.Key) *Key {
	return &Key{
		AccountID: key.AccountId,
		Name:      key.Name,
		Type:      Type(key.Type),
	}
}

func KeyToProto(key *Key) *pb.Key {
	return &pb.Key{
		AccountId: key.AccountID,
		Name:      key.Name,
		Type:      pb.Type(pb.Type_value[string(key.Type)]),
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
