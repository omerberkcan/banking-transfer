package service

import (
	"github.com/omerberkcan/banking-transfer/internal/config"
	"github.com/omerberkcan/banking-transfer/internal/repository"
	"github.com/omerberkcan/banking-transfer/internal/session"
)

type Services struct {
	Auth     AuthService
	Account  AccountService
	Transfer TransferService
}

func New(s repository.Stores, redis *session.Redis, cfg *config.SystemConfiguration) *Services {
	return &Services{
		Auth:     authService{store: s, redis: redis, cfg: cfg},
		Account:  accountService{store: s},
		Transfer: transferService{store: s},
	}
}
