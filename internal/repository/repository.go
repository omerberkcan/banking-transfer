package repository

import "github.com/omerberkcan/banking-transfer/model"

type Stores interface {
	Users() UserRepository
	Ping() error
}

type UserRepository interface {
	FindByID(id uint) (*model.User, error)
	FindByTCNo(TCNo string) (*model.User, error)
	UpdateBalance(userID uint, balance float32) error
	Create(login *model.User) (*model.User, error)
	Migrate() error
}
