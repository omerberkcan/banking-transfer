package api

import (
	"github.com/labstack/echo/v4"
	"github.com/omerberkcan/banking-transfer/dto"
	"github.com/omerberkcan/banking-transfer/helper"
	"github.com/omerberkcan/banking-transfer/internal/service"
)

type authHandler struct {
	authService service.IAuthService
}

type IAuthHandler interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}

var (
	invalidUser     = "Incorrect ID No or Password "
	validationError = "Validate Error"
	invalidJSON     = "Invalid Json"
	invalidIDNo     = "Invalid ID Number"
)

func (ah authHandler) Login(c echo.Context) error {
	var loginReq dto.LoginDTO
	var err error

	if err = c.Bind(&loginReq); err != nil {
		return err
	}

	if err = c.Validate(&loginReq); err != nil {
		return err
	}

	_, err = ah.authService.CheckLoginInformation(loginReq.IdNo, loginReq.Password)
	if err != nil {
		return err
	}

	//CreateToken

	return nil
}

func (ah authHandler) Register(c echo.Context) error {
	var registerReq dto.RegisterDTO
	var err error

	if err = c.Bind(&registerReq); err != nil {
		return err
	}

	if err = c.Validate(&registerReq); err != nil {
		return err
	}

	if !helper.IsNumeric(registerReq.IdNo) {
		return err
	}

	err = ah.authService.CheckAndSaveUser(registerReq)
	if err != nil {
		return err
	}
	return nil
}
