package user

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

type mockRepo struct {
	Repository
	mock.Mock
}

func (m *mockRepo) CreateUser(user *User) (*User, error) {
	args := m.Called(user)
	return args.Get(0).(*User), args.Error(1)
}

func (m *mockRepo) FindUserById(id string) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

func (m *mockRepo) UpdateUser(user *User) (*User, error) {
	args := m.Called(user)
	return args.Get(0).(*User), args.Error(1)
}

func (m *mockRepo) ListUsers(list []string) ([]*User, error) {
	args := m.Called(list)
	return args.Get(0).([]*User), args.Error(1)
}

func (m *mockRepo) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	cases := []struct {
		name     string
		req      *User
		mockFunc func(repo *mockRepo)
		want     *User
		err      error
	}{
		{
			name: "success",
			req: &User{
				Name:    "Jenny",
				Email:   "jenny@gmail.com",
				Address: "242 - Vista Verde",
				Cpf:     "12223560",
				Phone:   "12981463657",
				//Birthday: time.Time{},
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("CreateUser", &User{
					Name:    "Jenny",
					Email:   "jenny@gmail.com",
					Address: "242 - Vista Verde",
					Cpf:     "12223560",
					Phone:   "12981463657",
					//Birthday: time.Now(),
				}).
					Return(&User{
						Id:      "1",
						Name:    "Jenny",
						Email:   "jenny@gmail.com",
						Address: "242 - Vista Verde",
						Cpf:     "12223560",
						Phone:   "12981463657",
						//Birthday:  time.Now(),
						//CreatedAt: time.Now(),
						//UpdatedAt: time.Now(),
					}, nil)
			},
			want: &User{
				Id:      "1",
				Name:    "Jenny",
				Email:   "jenny@gmail.com",
				Address: "242 - Vista Verde",
				Cpf:     "12223560",
				Phone:   "12981463657",
				//Birthday:  time.Now(),
				//CreatedAt: time.Now(),
				//UpdatedAt: time.Now(),
			},
			err: nil,
		},
		//{
		//	name: "failed because error in create user",
		//	req: &User{
		//		Name:    "Jenny",
		//		Email:   "jenny@gmail.com",
		//		Address: "242 - Vista Verde",
		//		Cpf:     "12223560",
		//		Phone:   "12981463657",
		//		//Birthday: time.Now(),
		//	},
		//	mockFunc: func(repo *mockRepo) {
		//		repo.On("CreateUser", &User{
		//			Name:    "Jenny",
		//			Email:   "jenny@gmail.com",
		//			Address: "242 - Vista Verde",
		//			Cpf:     "12223560",
		//			Phone:   "12981463657",
		//			//Birthday: time.Now(),
		//		}).
		//			Return((*User)(nil), errors.New("MOCK-ERROR"))
		//	},
		//	want: nil,
		//	err:  errors.New("MOCK-ERROR"),
		//},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := new(mockRepo)

			c.mockFunc(repo)

			s := NewService(repo)

			got, err := s.CreateUser(c.req)
			require.NoError(t, err)
			require.NotNil(t, got)

			//assert.Equal(t, c.want.Birthday.Unix(), got.Birthday.Unix())
			//assert.Equal(t, c.want.CreatedAt.Unix(), got.CreatedAt.Unix())
			//assert.Equal(t, c.want.UpdatedAt.Unix(), got.UpdatedAt.Unix())

			//repo.AssertExpectations(t)
		})
	}
}

func TestFindUserById(t *testing.T) {
	cases := []struct {
		name     string
		req      *UserRequest
		mockFunc func(repo *mockRepo)
		want     *User
		err      error
	}{
		{
			name: "success",
			req: &UserRequest{
				UserID: "1",
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("FindUserById", "1").
					Return(&User{
						Id:      "1",
						Name:    "Jenny",
						Email:   "jenny@gmail.com",
						Address: "242 - Vista Verde",
						Cpf:     "12223560",
						Phone:   "12981463657",
						//Birthday:  time.Now(),
						//CreatedAt: time.Now(),
						//UpdatedAt: time.Now(),
					}, nil)
			},
			want: &User{
				Id:      "1",
				Name:    "Jenny",
				Email:   "jenny@gmail.com",
				Address: "242 - Vista Verde",
				Cpf:     "12223560",
				Phone:   "12981463657",
				//Birthday:  time.Now(),
				//CreatedAt: time.Now(),
				//UpdatedAt: time.Now(),
			},
			err: nil,
		},
		{
			name: "failed because error in find user by id",
			req: &UserRequest{
				UserID: "1",
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("FindUserById", "1").
					Return((*User)(nil), errors.New("MOCK-ERROR"))
			},
			want: nil,
			err:  errors.New("MOCK-ERROR"),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := new(mockRepo)
			c.mockFunc(repo)
			s := NewService(repo)
			got, err := s.FindUserById(c.req)
			assert.Equal(t, err, c.err)
			assert.Equal(t, got, c.want)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	cases := []struct {
		name     string
		req      *User
		mockFunc func(repo *mockRepo)
		want     *User
		err      error
	}{
		{
			name: "success",
			req: &User{
				Id:      "1",
				Name:    "Jenny",
				Email:   "jenny@gmail.com",
				Address: "242 - Vista Verde",
				Cpf:     "12223561",
				Phone:   "12981463658",
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("UpdateUser", &User{
					Id:      "1",
					Name:    "Jenny",
					Email:   "jenny@gmail.com",
					Address: "242 - Vista Verde",
					Cpf:     "12223561",
					Phone:   "12981463658",
				}).
					Return(&User{
						Id:      "1",
						Name:    "Jenny",
						Email:   "jenny@gmail.com",
						Address: "242 - Vista Verde",
						Cpf:     "12223561",
						Phone:   "12981463658",
					}, nil)
			},
			want: &User{
				Id:      "1",
				Name:    "Jenny",
				Email:   "jenny@gmail.com",
				Address: "242 - Vista Verde",
				Cpf:     "12223561",
				Phone:   "12981463658",
			},
			err: nil,
		},
		{
			name: "failed because error in update user",
			req: &User{
				Id: "1",
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("UpdateUser", &User{
					Id: "1",
				}).
					Return((*User)(nil), errors.New("MOCK-ERROR"))
			},
			want: nil,
			err:  errors.New("MOCK-ERROR"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := new(mockRepo)
			c.mockFunc(repo)
			s := NewService(repo)
			got, err := s.UpdateUser(c.req)
			assert.Equal(t, err, c.err)
			assert.Equal(t, got, c.want)
		})
	}
}

func TestListUsers(t *testing.T) {
	cases := []struct {
		name     string
		req      *ListUserRequest
		mockFunc func(repo *mockRepo)
		want     []*User
		err      error
	}{
		{
			name: "success",
			req: &ListUserRequest{
				UserIDs: []string{"1", "2"},
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("ListUsers", []string{"1", "2"}).
					Return([]*User{
						{
							Id: "1",
						},
						{
							Id: "2",
						},
					}, nil)
			},
			want: []*User{
				{
					Id: "1",
				},
				{
					Id: "2",
				},
			},
			err: nil,
		},
		{
			name: "failed because error in list users",
			req: &ListUserRequest{
				UserIDs: []string{"1", "2"},
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("ListUsers", []string{"1", "2"}).
					Return(([]*User)(nil), errors.New("MOCK-ERROR"))
			},
			want: nil,
			err:  errors.New("MOCK-ERROR"),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := new(mockRepo)

			c.mockFunc(repo)

			s := NewService(repo)

			got, err := s.ListUsers(c.req)
			assert.Equal(t, err, c.err)
			assert.Equal(t, got, c.want)
		})
	}

}

func TestDeleteUser(t *testing.T) {
	cases := []struct {
		name     string
		req      *UserRequest
		mockFunc func(repo *mockRepo)
		err      error
	}{
		{
			name: "success",
			req: &UserRequest{
				UserID: "1",
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("DeleteUser", "1").
					Return(nil)
			},
			err: nil,
		},
		{
			name: "failed because error in delete user",
			req: &UserRequest{
				UserID: "1",
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("DeleteUser", "1").
					Return(errors.New("MOCK-ERROR"))
			},
			err: errors.New("MOCK-ERROR"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := new(mockRepo)
			c.mockFunc(repo)
			s := NewService(repo)
			err := s.DeleteUser(c.req)
			assert.Equal(t, err, c.err)
		})
	}
}
