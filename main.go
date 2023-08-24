package main

import (
	"log"
	"net/http"

	"github.com/BaseMax/RabbitMQOrderGo/conf"
	"github.com/BaseMax/RabbitMQOrderGo/models"
)

func main() {
	if err := conf.Init(); err != nil {
		log.Fatal("conf:", err)
	}

	if err := models.InitDB(); err != nil {
		log.Fatal("db:", err)
	}

	r := initRoutes()
	addr := conf.GetHttpServerAddr()
	log.Fatal("http:", http.ListenAndServe(addr, r))
}
