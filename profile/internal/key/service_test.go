package key

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockRepoKey struct {
	Repository
	mock.Mock
}

func (m *mockRepoKey) CreateKey(key *Key) (*Key, error) {
	args := m.Called(key)
	return args.Get(0).(*Key), args.Error(1)
}

func (m *mockRepoKey) UpdateKey(key *Key) (*Key, error) {
	args := m.Called(key)
	return args.Get(0).(*Key), args.Error(1)
}

func (m *mockRepoKey) ListKey(ids []string) ([]*Key, error) {
	args := m.Called(ids)
	return args.Get(0).([]*Key), args.Error(1)
}

func (m *mockRepoKey) DeleteKey(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *mockRepoKey) FindKey(ctx context.Context, key string, accountId string) (*Key, error) {
	args := m.Called(ctx, key, accountId)
	return args.Get(0).(*Key), args.Error(1)
}

func TestCreateKey(t *testing.T) {
	cases := []struct {
		name     string
		req      *Key
		mockFunc func(repo *mockRepoKey)
		want     *Key
		err      error
	}{
		{
			name: "success",
			req: &Key{
				AccountID: "1",
				Name:      "test_key",
				Type:      Cpf,
			},
			mockFunc: func(repo *mockRepoKey) {
				repo.On("CreateKey", &Key{
					AccountID: "1",
					Name:      "test_key",
					Type:      Cpf,
				}).
					Return(&Key{
						Id:        "1",
						AccountID: "1",
						Name:      "test_key",
						Type:      Cpf,
					}, nil)
			},
			want: &Key{
				Id:        "1",
				AccountID: "1",
				Name:      "test_key",
				Type:      Cpf,
			},
			err: nil,
		},
		{
			name: "failed because error in create key",
			req: &Key{
				AccountID: "1",
				Name:      "test_key",
				Type:      Cpf,
			},
			mockFunc: func(repo *mockRepoKey) {
				repo.On("CreateKey", &Key{
					AccountID: "1",
					Name:      "test_key",
					Type:      Cpf,
				}).
					Return((*Key)(nil), errors.New("MOCK-ERROR"))
			},
			want: nil,
			err:  errors.New("MOCK-ERROR"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := new(mockRepoKey)
			c.mockFunc(repo)
			s := NewService(repo)
			got, err := s.CreateKey(c.req)
			assert.Equal(t, err, c.err)
			assert.Equal(t, got, c.want)
		})
	}
}

func TestUpdateKey(t *testing.T) {
	cases := []struct {
		name     string
		req      *Key
		mockFunc func(repo *mockRepoKey)
		want     *Key
		err      error
	}{
		{
			name: "success",
			req: &Key{
				Id:        "1",
				AccountID: "1",
				Name:      "test_key",
				Type:      Cpf,
			},
			mockFunc: func(repo *mockRepoKey) {
				repo.On("UpdateKey", &Key{
					Id:        "1",
					AccountID: "1",
					Name:      "test_key",
					Type:      Cpf,
				}).
					Return(&Key{
						Id:        "1",
						AccountID: "1",
						Name:      "test_key",
						Type:      Cpf,
					}, nil)
			},
			want: &Key{
				Id:        "1",
				AccountID: "1",
				Name:      "test_key",
				Type:      Cpf,
			},
			err: nil,
		},
		{
			name: "failed because error in update key",
			req: &Key{
				Id:        "1",
				AccountID: "1",
				Name:      "test_key",
				Type:      Cpf,
			},
			mockFunc: func(repo *mockRepoKey) {
				repo.On("UpdateKey", &Key{
					Id:        "1",
					AccountID: "1",
					Name:      "test_key",
					Type:      Cpf,
				}).
					Return((*Key)(nil), errors.New("MOCK-ERROR"))
			},
			want: nil,
			err:  errors.New("MOCK-ERROR"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := new(mockRepoKey)
			c.mockFunc(repo)
			s := NewService(repo)
			got, err := s.UpdateKey(c.req)
			assert.Equal(t, err, c.err)
			assert.Equal(t, got, c.want)
		})
	}
}

func TestListKey(t *testing.T) {
	cases := []struct {
		name     string
		req      *ListKeyRequest
		mockFunc func(repo *mockRepoKey)
		want     []*Key
		err      error
	}{
		{
			name: "success",
			req: &ListKeyRequest{
				keyIDs: []string{"1", "2"},
			},
			mockFunc: func(repo *mockRepoKey) {
				repo.On("ListKey", []string{"1", "2"}).
					Return([]*Key{
						{
							Id:        "1",
							AccountID: "1",
							Name:      "key_1",
							Type:      Cpf,
						},
						{
							Id:        "2",
							AccountID: "1",
							Name:      "key_2",
							Type:      Phone,
						},
					}, nil)
			},
			want: []*Key{
				{
					Id:        "1",
					AccountID: "1",
					Name:      "key_1",
					Type:      Cpf,
				},
				{
					Id:        "2",
					AccountID: "1",
					Name:      "key_2",
					Type:      Phone,
				},
			},
			err: nil,
		},
		{
			name: "failed because error in list keys",
			req: &ListKeyRequest{
				keyIDs: []string{"1", "2"},
			},
			mockFunc: func(repo *mockRepoKey) {
				repo.On("ListKey", []string{"1", "2"}).
					Return(([]*Key)(nil), errors.New("MOCK-ERROR"))
			},
			want: nil,
			err:  errors.New("MOCK-ERROR"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := new(mockRepoKey)
			c.mockFunc(repo)
			s := NewService(repo)
			got, err := s.ListKey(c.req)
			assert.Equal(t, err, c.err)
			assert.Equal(t, got, c.want)
		})
	}
}

func TestDeleteKey(t *testing.T) {
	cases := []struct {
		name     string
		req      *KeyRequest
		mockFunc func(repo *mockRepoKey)
		err      error
	}{
		{
			name: "success",
			req: &KeyRequest{
				keyID: "1",
			},
			mockFunc: func(repo *mockRepoKey) {
				repo.On("DeleteKey", "1").
					Return(nil)
			},
			err: nil,
		},
		{
			name: "failed because error in delete key",
			req: &KeyRequest{
				keyID: "2",
			},
			mockFunc: func(repo *mockRepoKey) {
				repo.On("DeleteKey", "2").
					Return(errors.New("MOCK-ERROR"))
			},
			err: errors.New("MOCK-ERROR"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := new(mockRepoKey)
			c.mockFunc(repo)
			s := NewService(repo)
			err := s.DeleteKey(c.req)
			assert.Equal(t, err, c.err)
		})
	}
}

func TestFindKey(t *testing.T) {
	cases := []struct {
		name      string
		key       string
		accountId string
		mockFunc  func(repo *mockRepoKey)
		want      *Key
		err       error
	}{
		{
			name:      "success",
			key:       "some_key",
			accountId: "1",
			mockFunc: func(repo *mockRepoKey) {
				repo.On("FindKey", context.Background(), "some_key", "1").
					Return(&Key{Id: "1", AccountID: "1", Name: "some_key", Type: Cpf}, nil)
			},
			want: &Key{
				Id:        "1",
				AccountID: "1",
				Name:      "some_key",
				Type:      Cpf,
			},
			err: nil,
		},
		{
			name:      "key not found",
			key:       "invalid_key",
			accountId: "1",
			mockFunc: func(repo *mockRepoKey) {
				repo.On("FindKey", context.Background(), "invalid_key", "1").
					Return((*Key)(nil), errors.New("MOCK-ERROR"))
			},
			want: nil,
			err:  errors.New("MOCK-ERROR"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := new(mockRepoKey)
			c.mockFunc(repo)
			s := NewService(repo)
			got, err := s.FindKey(context.Background(), c.key, c.accountId)
			assert.Equal(t, err, c.err)
			assert.Equal(t, got, c.want)
		})
	}
}
