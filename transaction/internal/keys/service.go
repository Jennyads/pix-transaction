package key

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
)

type Service interface {
	CreateKey(ctx context.Context, key *Key) (*Key, error)
	UpdateKey(ctx context.Context, key *Key) (*Key, error)
	ListKey(ctx context.Context, req *ListKeyRequest) ([]*Key, error)
	DeleteKey(ctx context.Context, id *KeyRequest) error
	FindKey(ctx context.Context, key *KeyRequest) (*Key, error)
}

type service struct {
	repo Repository
}

func (s *service) CreateKey(ctx context.Context, key *Key) (*Key, error) {
	findKey, err := s.repo.FindKey(ctx, key.Name)
	if err != nil {
		return nil, err
	}

	if findKey != nil {
		return nil, errors.New("key already exists")
	}

	key.CreatedAt = time.Now()
	key.UpdatedAt = time.Now()
	key.Id = uuid.New().String()
	return s.repo.CreateKey(ctx, key)
}

func (s *service) UpdateKey(ctx context.Context, key *Key) (*Key, error) {
	if key.Id == "" {
		return nil, errors.New("id is required")
	}
	key.UpdatedAt = time.Now()
	return s.repo.UpdateKey(ctx, key)
}

func (s *service) ListKey(ctx context.Context, req *ListKeyRequest) ([]*Key, error) {
	return s.repo.ListKey(ctx, req.keyIDs)
}
func (s *service) DeleteKey(ctx context.Context, request *KeyRequest) error {
	return s.repo.DeleteKey(ctx, request.keyID)
}

func (s *service) FindKey(ctx context.Context, key *KeyRequest) (*Key, error) {
	keys, err := s.repo.FindKey(ctx, key.keyID)
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
