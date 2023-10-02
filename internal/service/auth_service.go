package service

import (
	"github.com/labstack/gommon/log"
	"github.com/omerberkcan/banking-transfer/dto"
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
		CheckAndSaveUser(r dto.RegisterDTO) error
		CreateToken()
	}
)

func (as authService) CheckLoginInformation(idNo, password string) (*model.User, error) {
	user, err := as.store.Users().FindByIDNo(idNo)
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

func (as authService) CheckAndSaveUser(r dto.RegisterDTO) error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := model.User{Name: r.Name,
		IdNo:     r.IdNo,
		Balance:  r.Balance,
		Password: string(hashPass),
	}

	err = as.store.Users().Create(&user)

	return err

}

func (s authService) CreateToken() {

}
