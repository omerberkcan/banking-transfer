package main

import (
	"log"

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

	e := api.NewEcho()

	repo := repository.New(db)
	s := service.New(repo)
	handlers := api.NewHandler(s)

	api.SetRouter(e, handlers)

	e.Start(":" + c.System.Port)
}
