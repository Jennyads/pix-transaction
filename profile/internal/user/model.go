package user

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "profile/proto/v1"
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

func ProtoToUser(user *pb.User) *User {
	return &User{
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Cpf:      user.Cpf,
		Phone:    user.Phone,
		Birthday: user.Birthday.AsTime(),
	}
}

func ToProto(user *User) *pb.User {
	return &pb.User{
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Cpf:      user.Cpf,
		Phone:    user.Phone,
		Birthday: timestamppb.New(user.Birthday),
	}
}

func ProtoToUserRequest(request *pb.UserRequest) *UserRequest {
	return &UserRequest{
		UserID: request.Id,
	}

}

func ProtoToUserListRequest(request *pb.ListUserRequest) *ListUserRequest {
	return &ListUserRequest{
		UserIDs: request.Id,
	}
}
