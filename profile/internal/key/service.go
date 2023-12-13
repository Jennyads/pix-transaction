package key

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
)

type Service interface {
	CreateKey(key *Key) (*Key, error)
	UpdateKey(key *Key) (*Key, error)
	ListKey(req *ListKeyRequest) ([]*Key, error)
	DeleteKey(id *KeyRequest) error
	FindKey(ctx context.Context, key string, accountId string) (*Key, error)
}

type service struct {
	repo Repository
}

func (s service) CreateKey(key *Key) (*Key, error) {
	key.CreatedAt = time.Now()
	key.UpdatedAt = time.Now()
	key.Id = uuid.New().String()
	return s.repo.CreateKey(key)
}

func (s service) UpdateKey(key *Key) (*Key, error) {
	if key.Id == "" {
		return nil, errors.New("id is required")
	}
	key.UpdatedAt = time.Now()
	return s.repo.UpdateKey(key)
}

func (s service) ListKey(req *ListKeyRequest) ([]*Key, error) {
	if len(req.keyIDs) == 0 {
		return nil, errors.New("key_ids is required")
	}
	return s.repo.ListKey(req.keyIDs)
}
func (s service) DeleteKey(request *KeyRequest) error {
	if request.keyID == "" {
		return errors.New("account_id is required")
	}
	return s.repo.DeleteKey(request.keyID)
}

func (s *service) FindKey(ctx context.Context, key string, accountId string) (*Key, error) {
	keys, err := s.repo.FindKey(ctx, key, accountId)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
