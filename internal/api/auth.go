package api

import (
	"github.com/labstack/echo/v4"
	"github.com/omerberkcan/banking-transfer/internal/service"
)

type authHandler struct {
	a service.IAuthService
}

type IAuthHandler interface {
	Login(c echo.Context) error
}

func (a authHandler) Login(c echo.Context) error {
	return nil
}
