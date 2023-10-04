package api

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	customMiddle "github.com/omerberkcan/banking-transfer/internal/middleware"
	"github.com/omerberkcan/banking-transfer/internal/service"
)

type Handlers struct {
	Auth     AuthHandler
	Account  AccountHandler
	Transfer TransferHandler
}

func NewHandler(s *service.Services) *Handlers {
	return &Handlers{
		Auth:     authHandler{authService: s.Auth},
		Account:  accountHandler{acntService: s.Account},
		Transfer: transferHandler{transferService: s.Transfer},
	}
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	cv.validator.RegisterValidation("not blank", validators.NotBlank)
	return cv.validator.Struct(i)
}

func NewEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     []string{echo.DELETE, echo.GET, echo.PUT, echo.POST},
		AllowCredentials: true,
	}))
	e.Validator = &CustomValidator{validator: validator.New()}

	return e
}

func SetRouter(e *echo.Echo, h *Handlers, m *customMiddle.Middelwares, jwtSecret string) {
	e.POST("v1/login", h.Auth.Login)
	e.POST("v1/register", h.Auth.Register)

	config := echojwt.Config{
		SigningKey: []byte(jwtSecret),
		ContextKey: "user",
	}

	e.GET("v1/accounts/profile", h.Account.GetAccountInfoByID, echojwt.WithConfig(config), m.JwtMiddleware.AuthControl)
	e.POST("v1/transfers", h.Transfer.TransferMoney, echojwt.WithConfig(config), m.JwtMiddleware.AuthControl)

	e.GET("v1/health", HealthStatus)

}

func HealthStatus(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("Address:%s .%s", c.Echo().ListenerAddr().String(), "Service is healthy!"))
}
