package service

import (
	"github.com/omerberkcan/banking-transfer/internal/config"
	"github.com/omerberkcan/banking-transfer/internal/repository"
	"github.com/omerberkcan/banking-transfer/internal/session"
)

type Services struct {
	Auth AuthService
}

func New(s repository.Stores, redis session.Session, cfg *config.SystemConfiguration) *Services {
	return &Services{
		Auth: authService{store: s, redis: redis, cfg: cfg},
	}
}
