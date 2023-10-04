package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/omerberkcan/banking-transfer/dto"
	"github.com/omerberkcan/banking-transfer/helper"
	"github.com/omerberkcan/banking-transfer/internal/service"
)

var (
	invalidAmount = "amount must be bigger than zero"
)

type transferHandler struct {
	transferService service.TransferService
}

type TransferHandler interface {
	TransferMoney(c echo.Context) error
}

func (th transferHandler) TransferMoney(c echo.Context) error {
	var tReq dto.TransferDTO
	var err error

	userid, ok := c.Get("userid").(int)
	if !ok {
		return RespondWithError(c, http.StatusBadRequest, http.StatusBadRequest, invalidUserID)
	}

	if err = c.Bind(&tReq); err != nil {
		return RespondWithError(c, http.StatusBadRequest, http.StatusBadRequest, invalidJSON)
	}

	if err = c.Validate(&tReq); err != nil {
		return RespondWithError(c, http.StatusBadRequest, http.StatusBadRequest, validationError)
	}

	if !helper.IsNumeric(tReq.IDNo) {
		return RespondWithError(c, http.StatusBadRequest, http.StatusBadRequest, invalidIDNo)
	}

	if !tReq.Amount.IsPositive() {
		return RespondWithError(c, http.StatusBadRequest, http.StatusBadRequest, invalidAmount)
	}

	err = th.transferService.MoneyTransfer(uint(userid), tReq)
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, http.StatusBadRequest, err.Error())
	}

	return RespondWithOk(c, http.StatusOK, "Success")
}
