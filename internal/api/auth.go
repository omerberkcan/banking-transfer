package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/omerberkcan/banking-transfer/dto"
	"github.com/omerberkcan/banking-transfer/helper"
	"github.com/omerberkcan/banking-transfer/internal/service"
	"github.com/shopspring/decimal"
)

type authHandler struct {
	authService service.AuthService
}

type AuthHandler interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}

var (
	invalidUser     = "incorrect ID No or Password "
	validationError = "validate Error"
	invalidJSON     = "invalid Json"
	invalidIDNo     = "invalid ID Number"
	unexpectedError = "unexpected Error"
)

func (ah authHandler) Login(c echo.Context) error {
	var loginReq dto.LoginDTO
	var err error

	if err = c.Bind(&loginReq); err != nil {
		return RespondWithError(c, http.StatusBadRequest, http.StatusBadRequest, invalidJSON)
	}

	if err = c.Validate(&loginReq); err != nil {
		return RespondWithError(c, http.StatusBadRequest, http.StatusBadRequest, validationError)
	}

	usr, err := ah.authService.CheckLoginInformation(loginReq.IdNo, loginReq.Password)
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, http.StatusBadRequest, invalidUser)
	}

	token, err := ah.authService.CreateToken(usr)
	if err != nil {
		log.Printf("create token unexpected error: %s", err.Error())
		return RespondWithError(c, http.StatusBadRequest, http.StatusBadRequest, unexpectedError)
	}

	return RespondWithData(c, http.StatusOK, "Success", echo.Map{"access_token": token})
}

func (ah authHandler) Register(c echo.Context) error {
	var registerReq dto.RegisterDTO
	var err error

	if err = c.Bind(&registerReq); err != nil {
		return RespondWithError(c, http.StatusBadRequest, http.StatusBadRequest, invalidJSON)
	}

	if err = c.Validate(&registerReq); err != nil {
		return RespondWithError(c, http.StatusBadRequest, http.StatusBadRequest, validationError)
	}

	if !helper.IsNumeric(registerReq.IdNo) {
		return RespondWithError(c, http.StatusBadRequest, http.StatusBadRequest, invalidIDNo)
	}

	if !registerReq.Balance.IsPositive() {
		registerReq.Balance = decimal.NewFromFloat32(100)
	}

	err = ah.authService.CheckAndSaveUser(registerReq)
	if err != nil {
		return err
	}

	return RespondWithOk(c, http.StatusOK, "Success")
}
