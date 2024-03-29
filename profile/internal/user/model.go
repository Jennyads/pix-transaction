package user

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	proto "profile/proto/profile/v1"
	"time"
)

type User struct {
	Id        string         `gorm:"primaryKey;type:varchar(36);column:id"`
	Name      string         `gorm:"type:varchar(255);column:name"`
	Email     string         `gorm:"type:varchar(255);uniqueIndex;column:email"`
	Address   string         `gorm:"type:varchar(255);column:address"`
	Cpf       string         `gorm:"type:varchar(11);uniqueIndex;column:cpf"`
	Phone     string         `gorm:"type:varchar(20);column:phone"`
	Birthday  time.Time      `gorm:"type:date;column:birthday"`
	CreatedAt time.Time      `gorm:"type:datetime;column:created_at"`
	UpdatedAt time.Time      `gorm:"type:datetime;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;type:datetime;column:deleted_at"`
}

type UserRequest struct {
	UserID string
}

type ListUserRequest struct {
	UserIDs []string
}

func ProtoToUser(user *proto.User) *User {
	return &User{
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Cpf:      user.Cpf,
		Phone:    user.Phone,
		Birthday: user.Birthday.AsTime(),
	}
}
func ProtoToUserResponse(user *proto.UserResponse) *User {
	return &User{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Cpf:      user.Cpf,
		Phone:    user.Phone,
		Birthday: user.Birthday.AsTime(),
	}
}

func ToProtoUserResponse(user *User) *proto.UserResponse {
	return &proto.UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Cpf:      user.Cpf,
		Phone:    user.Phone,
		Birthday: timestamppb.New(user.Birthday),
	}
}

func ToProto(user *User) *proto.User {
	return &proto.User{
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Cpf:      user.Cpf,
		Phone:    user.Phone,
		Birthday: timestamppb.New(user.Birthday),
	}
}

func ProtoToUserRequest(request *proto.UserRequest) *UserRequest {
	return &UserRequest{
		UserID: request.Id,
	}

}

func ProtoToUserListRequest(request *proto.ListUserRequest) *ListUserRequest {
	return &ListUserRequest{
		UserIDs: request.Id,
	}
}
