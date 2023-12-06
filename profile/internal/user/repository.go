package user

import (
	"gorm.io/gorm"
	"profile/internal/cfg"
)

type Repository interface {
	CreateUser(user *User) (*User, error)
	FindUserById(id string) (*User, error)
	UpdateUser(user *User) (*User, error)
	ListUsers(ids []string) ([]*User, error)
	DeleteUser(id string) error
}

type repository struct {
	db  *gorm.DB
	cfg *cfg.Config
}

func (r repository) CreateUser(user *User) (*User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r repository) ExistWithCpf(cpf string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) FindUserById(id string) (*User, error) {
	var user User
	result := r.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r repository) UpdateUser(user *User) (*User, error) {
	result := r.db.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r repository) ListUsers(ids []string) ([]*User, error) {
	var listUser []*User
	if err := r.db.Where("id IN (?)", ids).Find(&listUser).Error; err != nil {
		return nil, err
	}
	return listUser, nil
}

func (r repository) DeleteUser(id string) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}

func NewRepository(db *gorm.DB, config *cfg.Config) Repository {
	return &repository{
		db:  db,
		cfg: config,
	}
}
