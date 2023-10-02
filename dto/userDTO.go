package dto

import "github.com/shopspring/decimal"

type LoginDTO struct {
	IdNo     string `json:"id_no" validate:"required,len=11"`
	Password string `json:"password" validate:"required"`
}

type RegisterDTO struct {
	IdNo     string          `json:"id_no" validate:"required,len=11"`
	Name     string          `json:"name" validate:"required"`
	Password string          `json:"password" validate:"required"`
	Balance  decimal.Decimal `json:"balance" validate:"min=0"`
}
