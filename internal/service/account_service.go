package service

import (
	"github.com/omerberkcan/banking-transfer/dto"
	"github.com/omerberkcan/banking-transfer/internal/repository"
)

type (
	accountService struct {
		store repository.Stores
	}

	AccountService interface {
		FindAccountByID(id uint) (*dto.UserDTO, error)
	}
)

func (as accountService) FindAccountByID(id uint) (*dto.UserDTO, error) {
	usr, err := as.store.Users().FindByID(id)
	if err != nil {
		return nil, err
	}

	userDTO := &dto.UserDTO{
		IdNo:    usr.IdNo,
		Name:    usr.Name,
		Balance: usr.Balance,
	}

	return userDTO, nil
}
