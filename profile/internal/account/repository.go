package account

import (
	"gorm.io/gorm"
	"profile/internal/cfg"
)

type Repository interface {
	CreateAccount(account *Account) (*Account, error)
	FindAccountById(id string) (*Account, error)
	UpdateAccount(account *Account) (*Account, error)
	ListAccount(accountIDs []string) ([]*Account, error)
	DeleteAccount(id string) error
}

type repository struct {
	db  *gorm.DB
	cfg *cfg.Config
}

func (r repository) CreateAccount(account *Account) (*Account, error) {
	err := r.db.Create(&account).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (r repository) FindAccountById(id string) (*Account, error) {
	var account Account
	result := r.db.Where("id = ?", id).First(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

func (r repository) UpdateAccount(account *Account) (*Account, error) {
	result := r.db.Save(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return account, nil
}

func (r repository) ListAccount(ids []string) ([]*Account, error) {
	var listAccount []*Account
	if err := r.db.Where("id IN (?)", ids).Find(&listAccount).Error; err != nil {
		return nil, err
	}
	return listAccount, nil
}

func (r repository) DeleteAccount(id string) error {
	return r.db.Delete(&Account{}, "id = ?", id).Error
}

func NewRepository(db *gorm.DB, config *cfg.Config) Repository {
	return &repository{
		db:  db,
		cfg: config,
	}
}
