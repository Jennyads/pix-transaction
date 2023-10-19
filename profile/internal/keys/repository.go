package keys

import "profile/platform/dynamo"

type Repository interface {
	CreateKey(key *Key) (*Key, error)
	UpdateKey(key *Key) (*Key, error)
	ListKey(keyIDs []int64) ([]*Key, error)
	DeleteKey(id int) error
}

type repository struct {
	db dynamo.Client
}

func (r repository) CreateKey(key *Key) (*Key, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) UpdateKey(key *Key) (*Key, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) ListKey(keyIDs []int64) ([]*Key, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) DeleteKey(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db dynamo.Client) Repository {
	return &repository{
		db: db,
	}
}
