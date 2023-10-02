package repository

import "github.com/omerberkcan/banking-transfer/model"

type Stores interface {
	Users() UserRepository
	Ping() error
}

type UserRepository interface {
	FindByID(id uint) (*model.User, error)
	FindByIDNo(IDNo string) (*model.User, error)
	UpdateBalance(userID uint, balance float32) error
	Create(user *model.User) error
	Migrate() error
}
