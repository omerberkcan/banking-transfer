package service

import (
	"github.com/labstack/gommon/log"
	"github.com/omerberkcan/banking-transfer/internal/repository"
	"github.com/omerberkcan/banking-transfer/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	authService struct {
		store repository.Stores
	}

	IAuthService interface {
		CheckLoginInformation(idNo, password string) (*model.User, error)
		CreateToken()
	}
)

func (s authService) CheckLoginInformation(idNo, password string) (*model.User, error) {
	user, err := s.store.Users().FindByIDNo(idNo)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		log.Errorf("When get user in db, error occured : %s", err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s authService) CreateToken() {

}
