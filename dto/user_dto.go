package dto

import (
	"github.com/shopspring/decimal"
)

type UserDTO struct {
	IdNo    string          `json:"id_no"`
	Name    string          `json:"name"`
	Balance decimal.Decimal `json:"balance"`
}
