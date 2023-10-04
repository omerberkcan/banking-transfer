package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TokenDetails struct {
	UserID    int
	Token     string
	Uuid      uuid.UUID
	AtExpires time.Time
}
