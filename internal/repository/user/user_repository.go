package user

import (
	"github.com/omerberkcan/banking-transfer/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (ur *Repository) FindByTCNo(TCNo string) (*model.User, error) {
	var user model.User
	err := ur.db.Where("tcno = ?", TCNo).First(&user).Error
	if err != nil {
		return &user, err
	}
	return &user, err
}

func (ur *Repository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := ur.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return &user, err
	}

	return &user, err
}

func (ur *Repository) Create(user *model.User) (*model.User, error) {
	err := ur.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *Repository) UpdateBalance(userID uint, balance float32) error {
	err := ur.db.Model(&model.User{}).Where("id = ?", userID).Update("balance", balance).Error
	if err != nil {
		return err
	}

	return nil
}

// Migrate ...
func (ur *Repository) Migrate() error {
	return ur.db.AutoMigrate(&model.User{})
}
