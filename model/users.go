package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	IdNo     string          `gorm:"type:varchar(11);not null;uniqueIndex;comment:TCNo"` //This was made only for Turkish citizens
	Name     string          `gorm:"type:varchar(100);not null"`
	Password string          `gorm:"type:varchar(100);not null"`
	Balance  decimal.Decimal `gorm:"type:decimal(10,2)"`
}
