package service

import "github.com/omerberkcan/banking-transfer/internal/repository"

type Services struct {
	Auth IAuthService
}

func New(s repository.Stores) *Services {
	return &Services{
		Auth: authService{store: s},
	}
}
