package user

import (
	pb "profile/proto/v1"
	"time"
)

type User struct {
	Id        string
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
