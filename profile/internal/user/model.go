package user

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	proto "profile/proto/v1"
	"time"
)

type User struct {
	Id        string `dynamodbav:"PK"`
	Name      string
	Email     string
	Address   string
	Cpf       string
	Phone     string
	Birthday  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
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
