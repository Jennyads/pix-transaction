package user

type Service interface {
	CreateUser(user *User) (*User, error)
	FindUserById(req *UserRequest) (*User, error)
	UpdateUser(user *User) (*User, error)
	ListUsers(req *ListUserRequest) ([]*User, error)
	DeleteUser(req *UserRequest) error
}

type service struct {
	repo Repository
}

func (s service) CreateUser(user *User) (*User, error) {
	return s.repo.CreateUser(user)
}

func (s service) FindUserById(userRequest *UserRequest) (*User, error) {
	return s.repo.FindUserById(userRequest.UserID)
}

func (s service) UpdateUser(user *User) (*User, error) {
	return s.repo.UpdateUser(user)
}

func (s service) ListUsers(listUsers *ListUserRequest) ([]*User, error) {
	return s.repo.ListUsers(listUsers.UserIDs)
}

func (s service) DeleteUser(request *UserRequest) error {
	return s.repo.DeleteUser(request.UserID)
}
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
