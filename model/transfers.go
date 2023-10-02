package model

import "gorm.io/gorm"

type Transfer struct {
	gorm.Model
	UserOriginID      uint
	UserDestinationID uint
}
