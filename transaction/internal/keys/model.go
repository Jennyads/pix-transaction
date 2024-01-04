package key

import (
	"time"
	proto "transaction/proto/v1"
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
	Agency    string
	Bank      string
	Account   string
	Name      string
	Type      Type
	CreatedAt time.Time
	UpdatedAt time.Time
}

type KeyRequest struct {
	keyID string
}

type ListKeyRequest struct {
	keyIDs []string
}

func ProtoToKey(key *proto.Key) *Key {
	return &Key{
		Account: key.Account,
		Name:    key.Name,
		Agency:  key.Agency,
		Bank:    key.Bank,
		Type:    Type(key.Type.String()),
	}
}

func ToProto(key *Key) *proto.KeyResponse {
	return &proto.KeyResponse{
		Id:      key.Id,
		Account: key.Account,
		Name:    key.Name,
		Type:    proto.Type(proto.Type_value[string(key.Type)]),
		Agency:  key.Agency,
		Bank:    key.Bank,
	}
}

func ProtoToKeyListRequest(request *proto.ListKeyRequest) *ListKeyRequest {
	return &ListKeyRequest{
		keyIDs: request.KeyId,
	}
}
func ProtoToKeyRequest(request *proto.KeyRequest) *KeyRequest {
	return &KeyRequest{
		keyID: request.KeyId,
	}
}
