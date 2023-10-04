package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/omerberkcan/banking-transfer/internal/service"
)

var (
	invalidUserID = "invalid user id"
	userNotFound  = "user not found"
)

type accountHandler struct {
	acntService service.AccountService
}

type AccountHandler interface {
	GetAccountInfoByID(c echo.Context) error
}

func (a accountHandler) GetAccountInfoByID(c echo.Context) error {
	userid, ok := c.Get("userid").(int)
	if !ok {
		return RespondWithError(c, http.StatusBadRequest, http.StatusBadRequest, invalidUserID)
	}

	usr, err := a.acntService.FindAccountByID(uint(userid))
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, http.StatusBadRequest, userNotFound)
	}

	return RespondWithData(c, http.StatusOK, "success", usr)
}
