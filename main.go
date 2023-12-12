package main

import (
	"log"

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

	if err := models.Migrate(); err != nil {
		log.Fatal("migrate:", err)
	}

	r := initRoutes()
	r.Logger.Fatal(r.Start(conf.GetHttpServerAddr()))
}
