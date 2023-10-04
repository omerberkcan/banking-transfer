package dto

import "github.com/shopspring/decimal"

type TransferDTO struct {
	IDNo        string          `json:"id_no" validate:"required,len=11"`
	Amount      decimal.Decimal `json:"amount" validate:"required"`
	Description string          `json:"description"`
}
