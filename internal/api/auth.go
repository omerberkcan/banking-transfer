package api

import (
	"github.com/labstack/echo/v4"
	"github.com/omerberkcan/banking-transfer/dto"
	"github.com/omerberkcan/banking-transfer/internal/service"
)

type authHandler struct {
	authService service.IAuthService
}

type IAuthHandler interface {
	Login(c echo.Context) error
}

var (
	invalidUser     = "Incorrect ID No or Password "
	validationError = "Validate Error"
	invalidJSON     = "Invalid Json"
)

func (a authHandler) Login(c echo.Context) error {
	var loginReq dto.LoginDTO
	var err error

	if err = c.Bind(&loginReq); err != nil {
		return err
	}

	if err = c.Validate(&loginReq); err != nil {
		return err //api.RespondWithError(c, http.StatusBadRequest, global.ValidationError, "Invalid Json Data", err)
	}

	_, err = a.authService.CheckLoginInformation(loginReq.IdNo, loginReq.Password)
	if err != nil {
		return err // api.RespondWithError(c, http.StatusUnauthorized, global.InvalidUser, "Incorrect Username Or Password", err)
	}

	//CreateToken

	return nil
}
