package repository

import (
	"github.com/omerberkcan/banking-transfer/internal/repository/transfer"
	"github.com/omerberkcan/banking-transfer/internal/repository/user"
	"github.com/omerberkcan/banking-transfer/model"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Stores interface {
	Users() UserRepository
	Transfer() TransferRepository
	TxBegin() *gorm.DB
	Ping() error
}

type UserRepository interface {
	FindByID(id uint) (*model.User, error)
	FindByIDNo(IDNo string) (*model.User, error)
	UpdateBalance(userID uint, balance decimal.Decimal) error
	Create(user *model.User) error
	WithTrx(trxHandle *gorm.DB) user.Repository
}

type TransferRepository interface {
	FindByOriginID(orignID int) ([]model.Transfer, error)
	Create(t *model.Transfer) error
	WithTrx(trxHandle *gorm.DB) transfer.Repository
}
