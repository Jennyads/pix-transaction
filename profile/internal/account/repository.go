package account

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"profile/internal/cfg"
	"sync"
	"time"
)

type Repository interface {
	CreateAccount(account *Account) (*Account, error)
	FindAccountById(id int64) (*Account, error)
	UpdateAccount(account *Account) (*Account, error)
	ListAccount(accountIDs []int64) ([]*Account, error)
	DeleteAccount(id int64) error
	FindByKey(key string) (*Account, error)
	IsAccountActive(ctx context.Context, id int64) (bool, error)
	UpdateAccounts(accounts []*Account) error
}

type repository struct {
	db        *gorm.DB
	cfg       *cfg.Config
	idCounter int64
	mu        sync.Mutex
}

func (r repository) CreateAccount(account *Account) (*Account, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()

	r.idCounter++
	account.Id = r.idCounter

	return account, nil
}

func (r repository) FindAccountById(id int64) (*Account, error) {
	var account Account
	result := r.db.Where("id = ?", id).First(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

func (r repository) UpdateAccount(account *Account) (*Account, error) {
	account.UpdatedAt = time.Now()
	result := r.db.Save(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return account, nil
}

func (r repository) UpdateAccounts(accounts []*Account) error {
	tx := r.db.Begin()
	now := time.Now()
	for _, account := range accounts {
		account.UpdatedAt = now
		result := tx.Save(&account)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}
	tx.Commit()
	return nil
}

func (r repository) ListAccount(ids []int64) ([]*Account, error) {
	var listAccount []*Account
	if err := r.db.Where("id IN (?)", ids).Find(&listAccount).Error; err != nil {
		return nil, err
	}
	return listAccount, nil
}

func (r repository) DeleteAccount(id int64) error {
	return r.db.Delete(&Account{}, "id = ?", id).Error
}

func (r repository) FindByKey(key string) (*Account, error) {
	var account Account
	err := r.db.Raw(`SELECT accounts.* FROM accounts
						INNER JOIN keys on accounts.id = keys.account_id
						WHERE keys.name = @key`,
		sql.Named("key", key)).Scan(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, err
}

func (r repository) IsAccountActive(ctx context.Context, id int64) (bool, error) {
	var account Account
	result := r.db.Select("deleted_at").Where("id = ?", id).First(&account)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	return account.DeletedAt == gorm.DeletedAt{}, nil
}

func NewRepository(db *gorm.DB, config *cfg.Config) Repository {
	return &repository{
		db:  db,
		cfg: config,
	}
}
