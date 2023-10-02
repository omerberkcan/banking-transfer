package service

import "github.com/omerberkcan/banking-transfer/internal/repository"

type (
	authService struct {
		store repository.Stores
	}

	IAuthService interface {
		CheckLoginInformation()
		CreateToken()
	}
)

func (s authService) CheckLoginInformation() {

}

func (s authService) CreateToken() {

}
