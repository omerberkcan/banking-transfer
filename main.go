package main

import (
	"log"

	"github.com/omerberkcan/banking-transfer/internal/config"
)

func main() {
	_, err := config.Init()
	if err != nil {
		log.Fatalf("config init error:%s", err.Error())
	}
}
