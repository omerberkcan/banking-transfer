package transfer

import (
	"log"

	"github.com/omerberkcan/banking-transfer/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (tr Repository) Create(t *model.Transfer) error {
	err := tr.db.Create(&t).Error
	return err
}

func (tr *Repository) FindByOriginID(orignID int) ([]model.Transfer, error) {
	var t []model.Transfer
	err := tr.db.Joins("UserDestination").Where("user_origin_id = ? ", orignID).Find(&t).Error
	return t, err
}

func (tr Repository) WithTrx(trxHandle *gorm.DB) Repository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return tr
	}
	tr.db = trxHandle
	return tr
}
