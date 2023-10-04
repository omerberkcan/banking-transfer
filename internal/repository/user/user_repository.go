package user

import (
	"log"

	"github.com/omerberkcan/banking-transfer/model"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (ur *Repository) FindByIDNo(IDNo string) (*model.User, error) {
	var user model.User
	err := ur.db.Where("id_no = ?", IDNo).First(&user).Error
	return &user, err
}

func (ur *Repository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := ur.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (ur *Repository) Create(user *model.User) error {
	err := ur.db.Create(&user).Error
	return err
}

func (ur Repository) UpdateBalance(userID uint, balance decimal.Decimal) error {
	err := ur.db.Model(&model.User{}).Where("id = ?", userID).Update("balance", balance).Error
	return err
}

func (ur Repository) WithTrx(trxHandle *gorm.DB) Repository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return ur
	}
	ur.db = trxHandle
	return ur
}
