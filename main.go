package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/omerberkcan/banking-transfer/internal/api"
	"github.com/omerberkcan/banking-transfer/internal/config"
	"github.com/omerberkcan/banking-transfer/internal/repository"
	"github.com/omerberkcan/banking-transfer/internal/service"
)

func main() {
	c, err := config.Init()
	if err != nil {
		log.Fatalf("config init error:%s", err.Error())
	}

	db, err := repository.ConnectMysqlServer(&c.MySQL)
	if err != nil {
		log.Fatalf("cannot connect mysql server: %s", err.Error())
	}
	repo := repository.New(db)

	s := service.New(repo)

	handlers := api.NewHandler(s)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     []string{echo.DELETE, echo.GET, echo.PUT, echo.POST},
		AllowCredentials: true,
	}))

	e.POST("login", handlers.Auth.Login)

}
