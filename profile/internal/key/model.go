package key

import (
	"gorm.io/gorm"
	proto "profile/proto/profile/v1"
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
	Id        string         `gorm:"primaryKey;type:varchar(36);column:id"`
	AccountID int64          `gorm:"foreignKey;type:varchar(36);column:account_id"`
	Name      string         `gorm:"type:varchar(200);column:name"`
	Type      Type           `gorm:"type:varchar(100);column:type"`
	CreatedAt time.Time      `gorm:"type:datetime;column:created_at"`
	UpdatedAt time.Time      `gorm:"type:datetime;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;type:datetime;column:deleted_at"`
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
		Type:      Type(key.Type.String()),
	}
}

func ToProto(key *Key) *proto.KeyResponse {
	return &proto.KeyResponse{
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
