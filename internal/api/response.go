package api

import "github.com/labstack/echo/v4"

type ResponseDTO struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func RespondWithError(c echo.Context, code int, spccode int, message string) error {
	return RespondWithJSON(c, code, ResponseDTO{Code: spccode, Status: "Error", Message: message})
}

func RespondWithOk(c echo.Context, code int, message string) error {
	return RespondWithJSON(c, code, ResponseDTO{Code: code, Status: "Success", Message: message})
}

func RespondWithJSON(c echo.Context, code int, payload ResponseDTO) error {
	c.Request().Header.Set("Content-Type", "application/json")
	return c.JSON(code, payload)
}

func RespondWithData(c echo.Context, code int, message string, Data interface{}) error {
	return RespondWithJSON(c, code, ResponseDTO{Code: code, Status: "Success", Message: message, Data: Data})
}
