package key

import (
	"gorm.io/gorm"
	"profile/internal/cfg"
)

type Repository interface {
	CreateKey(key *Key) (*Key, error)
	UpdateKey(key *Key) (*Key, error)
	ListKey(ids []string) ([]*Key, error)
	DeleteKey(id string) error
}

type repository struct {
	db  *gorm.DB
	cfg *cfg.Config
}

func (r repository) CreateKey(key *Key) (*Key, error) {
	err := r.db.Create(&key).Error
	if err != nil {
		return nil, err
	}
	return key, nil
}

func (r repository) UpdateKey(key *Key) (*Key, error) {
	result := r.db.Save(&key)
	if result.Error != nil {
		return nil, result.Error
	}
	return key, nil
}

func (r repository) ListKey(ids []string) ([]*Key, error) {
	var listKey []*Key
	if err := r.db.Where("id IN (?)", ids).Find(&listKey).Error; err != nil {
		return nil, err
	}
	return listKey, nil
}

func (r repository) DeleteKey(id string) error {
	return r.db.Delete(&Key{}, "id = ?", id).Error
}

func NewRepository(db *gorm.DB, config *cfg.Config) Repository {
	return &repository{
		db:  db,
		cfg: config,
	}
}
