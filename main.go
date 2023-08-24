package main

import (
	"log"
	"net/http"

	"github.com/BaseMax/RabbitMQOrderGo/conf"
)

func main() {
	if err := conf.Init(); err != nil {
		log.Fatal("conf:", err)
	}

	r := initRoutes()
	addr := conf.GetHttpServerAddr()
	log.Fatal("http:", http.ListenAndServe(addr, r))
}
