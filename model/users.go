package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TCNo    string          `gorm:"type:varchar(11);not null;uniqueIndex"`
	Name    string          `gorm:"type:varchar(100);not null"`
	Balance decimal.Decimal `gorm:"type:decimal(10,2)"`
}
