package dto

import "github.com/shopspring/decimal"

type LoginDTO struct {
	IdNo     string `json:"id_no" validate:"required,len=11,not blank"`
	Password string `json:"password" validate:"required,not blank"`
}

type RegisterDTO struct {
	IdNo     string          `json:"id_no" validate:"required,len=11,not blank"`
	Name     string          `json:"name" validate:"required,not blank"`
	Password string          `json:"password" validate:"required,not blank"`
	Balance  decimal.Decimal `json:"balance"`
}
