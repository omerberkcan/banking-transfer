package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	UserOriginID      uint
	UserOrigin        User `gorm:"foreignKey:UserOriginID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserDestinationID uint
	UserDestination   User            `gorm:"foreignKey:UserDestinationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Amount            decimal.Decimal `gorm:"type:decimal(10,2)"`
	Description       string          `gorm:"type:text"`
	Tax               decimal.Decimal `gorm:"type:decimal(10,2)"`
}
