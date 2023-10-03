package service

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/gommon/log"
	"github.com/omerberkcan/banking-transfer/dto"
	"github.com/omerberkcan/banking-transfer/internal/config"
	"github.com/omerberkcan/banking-transfer/internal/repository"
	"github.com/omerberkcan/banking-transfer/internal/session"
	"github.com/omerberkcan/banking-transfer/model"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	authService struct {
		store repository.Stores
		redis *session.Redis
		cfg   *config.SystemConfiguration
	}

	AuthService interface {
		CheckLoginInformation(idNo, password string) (*model.User, error)
		CheckAndSaveUser(r dto.RegisterDTO) error
		CreateToken(usr *model.User) (string, error)
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

func (s authService) CreateToken(usr *model.User) (string, error) {
	var err error
	accessSecret := s.cfg.TokenSecretKey
	accessTokenExpireDuration := resolveTokenExpireDuration(s.cfg.AccessTokenExpireTime)

	accessUuid := uuid.NewV4()
	// td.RefreshUuid = uuid.NewV4()
	atExpires := time.Now().Add(accessTokenExpireDuration)
	// td.RtExpires = time.Now().Add(refreshTokenExpireDuration)

	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = usr.ID
	atClaims["uuid"] = accessUuid.String()
	atClaims["exp"] = time.Now().Add(accessTokenExpireDuration).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	acccessToken, err := at.SignedString([]byte(accessSecret))
	if err != nil {
		return "", err
	}

	s.redis.DeleteTokenByUserID(int(usr.ID))
	s.redis.SetToken(model.TokenDetails{UserID: int(usr.ID), Uuid: accessUuid, AtExpires: atExpires}, accessTokenExpireDuration)

	return acccessToken, err
}

func resolveTokenExpireDuration(config string) time.Duration {
	duration, _ := strconv.ParseInt(config[0:len(config)-1], 10, 64)
	timeFormat := config[len(config)-1:]

	switch timeFormat {
	case "m":
		return time.Duration(time.Minute.Nanoseconds() * duration)
	case "h":
		return time.Duration(time.Hour.Nanoseconds() * duration)
	case "d":
		return time.Duration(time.Hour.Nanoseconds() * 24 * duration)
	default:
		return time.Duration(time.Minute.Nanoseconds() * 30)
	}
}
