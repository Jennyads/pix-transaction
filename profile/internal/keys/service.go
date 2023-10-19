package keys

import (
	"errors"
	"time"
)

type Service interface {
	CreateKey(key *Key) (*Key, error)
	UpdateKey(key *Key) (*Key, error)
	ListKey(req *ListKeyRequest) ([]*Key, error)
	DeleteKey(id *KeyRequest) error
}

type service struct {
	repo Repository
}

func (s service) CreateKey(key *Key) (*Key, error) {
	key.CreatedAt = time.Now()
	key.UpdatedAt = time.Now()
	return s.repo.CreateKey(key)
}

func (s service) UpdateKey(key *Key) (*Key, error) {
	if key.Id == 0 {
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
	if request.keyID == 0 {
		return errors.New("account_id is required")
	}
	return s.repo.DeleteKey(int(request.keyID))
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
