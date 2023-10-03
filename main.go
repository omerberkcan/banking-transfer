package main

import (
	"log"

	"github.com/omerberkcan/banking-transfer/internal/api"
	"github.com/omerberkcan/banking-transfer/internal/config"
	"github.com/omerberkcan/banking-transfer/internal/middleware"
	"github.com/omerberkcan/banking-transfer/internal/repository"
	"github.com/omerberkcan/banking-transfer/internal/service"
	"github.com/omerberkcan/banking-transfer/internal/session"
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

	log.Printf("config %v", c)
	redis, err := session.RedisConnect(&c.Redis)
	if err != nil {
		log.Fatalf("cannot connect redis: %s", err.Error())
	}
	redisRepo := session.New(redis)

	e := api.NewEcho()

	repo := repository.New(db)
	s := service.New(repo, redisRepo, &c.System)
	handlers := api.NewHandler(s)
	middlewares := middleware.New(redisRepo)

	api.SetRouter(e, handlers, middlewares, c.System.TokenSecretKey)

	e.Start(":" + c.System.Port)
}
