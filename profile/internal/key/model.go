package key

import (
	"gorm.io/gorm"
	proto "profile/proto/v1"
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
	Id        string `gorm:"primarykey;type:varchar(36)"`
	AccountID string `gorm:"foreignKey;type:varchar(36)"`
	Name      string
	Type      Type
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type KeyRequest struct {
	keyID string
}

type ListKeyRequest struct {
	keyIDs []string
}

func ProtoToKey(key *proto.Key) *Key {
	return &Key{
		AccountID: key.AccountId,
		Name:      key.Name,
		Type:      Type(key.Type),
	}
}

func ToProto(key *Key) *proto.Key {
	return &proto.Key{
		AccountId: key.AccountID,
		Name:      key.Name,
		Type:      proto.Type(proto.Type_value[string(key.Type)]),
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
